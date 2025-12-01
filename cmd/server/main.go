package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gin/internal/api"
	apimiddleware "gin/internal/api/middleware" // API中间件包别名
	"gin/internal/config"
	"gin/internal/database"
	"gin/internal/di"
	"gin/internal/errors"
	"gin/internal/logger"
	"gin/internal/metrics"
	appmiddleware "gin/internal/middleware" // 应用中间件包别名

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func Main() {
	// 1. 加载配置
	cfg := config.LoadConfig()

	// 2. 初始化日志系统
	log := logger.InitLogger(&cfg.Logging)
	defer func(log *zap.Logger) {
		err := log.Sync()
		if err != nil {
			panic(err)
		}
	}(log)

	// 3. 初始化数据库连接（如果有DSN配置）
	if cfg.Database.DSN != "" && cfg.Database.DSN != "user:password@tcp(localhost:3306)/dbname" {
		db, err := database.InitDB(cfg.Database.DSN)
		if err != nil {
			log.Error("数据库初始化失败", zap.Error(err))
		} else {
			defer func(db database.DB) {
				err := db.Close()
				if err != nil {
					panic(err)
				}
			}(db)
			// 注册到DI容器
			di.GetContainer().Register("db", db)
		}
	}

	// 4. 设置路由
	router := api.SetupRouter()

	// 5. 添加中间件
	router.Use(apimiddleware.GinBodyLogMiddleware()) // 使用API中间件
	router.Use(errors.ErrorHandler())                // 添加错误处理中间件
	router.Use(metrics.PrometheusMiddleware())       // 添加指标中间件
	router.Use(appmiddleware.Recovery())             // 使用应用中间件的Recovery

	// 6. 添加指标路由
	router.GET("/metrics", metrics.MetricsHandler())

	// 创建errgroup
	g, ctx := errgroup.WithContext(context.Background())

	// 启动多个服务器
	servers := []struct {
		addr    string
		handler http.Handler
	}{
		{":" + cfg.Server.Port, router}, // 添加冒号前缀
		{":8081", api.Router01()},
		{":8082", api.Router02()},
	}

	// 为每个服务器启动一个goroutine
	for _, s := range servers {
		srv := &http.Server{
			Addr:         s.addr,
			Handler:      s.handler,
			ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
		}

		// 使用局部变量保存当前服务器地址，避免闭包问题
		addr := s.addr

		g.Go(func() error {
			log.Info("服务器启动", zap.String("addr", addr))
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				return fmt.Errorf("监听错误 %s: %w", addr, err)
			}
			return nil
		})

		// 优雅关闭
		g.Go(func() error {
			<-ctx.Done()
			log.Info("正在关闭服务器", zap.String("addr", addr))
			shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			return srv.Shutdown(shutdownCtx)
		})
	}

	// 监听中断信号
	g.Go(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			log.Info("收到信号", zap.Stringer("signal", sig))
			return fmt.Errorf("收到信号: %v", sig)
		case <-ctx.Done():
			return ctx.Err()
		}
	})

	if err := g.Wait(); err != nil {
		log.Error("服务器异常退出", zap.Error(err))
	}
}
