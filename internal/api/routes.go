package api

import (
	"gin/internal/api/handlers"
	"net/http"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
)

// 返回 main.go 所在的文件夹
func getCurrentPath() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Dir(file)
}
func SetupRouter() *gin.Engine {
	router := gin.Default()
	basePath := getCurrentPath()
	// 向上跳两级到项目根目录
	basePath = filepath.Dir(filepath.Dir(basePath))

	// 模板和静态文件设置
	handlers.SetupTemplates(router, basePath)

	// 基础路由
	router.GET("/hello", handlers.HelloHandler())
	router.GET("/posts/index", handlers.PostsIndexHandler())
	router.GET("/users/index", handlers.UsersIndexHandler())
	router.GET("/index", handlers.CustomTemplateHandler())

	// 渲染相关路由
	router.GET("/v1/index", handlers.IndexFunc())
	router.GET("/v1/home", handlers.HomeFunc())
	router.GET("/someJSON", handlers.SomeJSONHandler())
	router.GET("moreJSON", handlers.MoreJSONHandler())
	router.GET("/pb", handlers.ProtoHandler())

	// 参数获取相关路由
	router.GET("/user/search", handlers.UserSearchHandler())
	router.POST("posts/search", handlers.PostsSearchHandler())
	router.POST("/json", handlers.JSONHandler())
	router.GET("/user/get/:username/:address", handlers.UserGetHandler())

	// 表单绑定相关路由
	router.POST("/loginJSON", handlers.LoginJSONHandler())
	router.POST("/loginForm", handlers.LoginFormHandler())
	router.GET("/loginQuery", handlers.LoginQueryHandler())
	router.GET("/loginHeader", handlers.LoginHeaderHandler())
	router.GET("/loginUri/:id", handlers.LoginUriHandler())

	// 文件上传相关路由
	router.GET("/upload/page", handlers.UploadPageHandler())
	router.POST("/upload", handlers.UploadHandler())
	router.GET("/upload/multi/page", handlers.UploadMultiPageHandler())
	router.POST("/upload/multi", handlers.UploadMultiHandler())

	// 重定向相关路由
	router.GET("/http/redirect", handlers.HTTPRedirectHandler())
	router.GET("/path/redirect", handlers.PathRedirectHandler(router))

	// 其他路由
	router.Any("/test", handlers.TestHandler())
	router.NoRoute(handlers.NoRouteHandler())

	// 路由组（已迁移到 SetupRouterWithDI，此处保留注释）
	// userGroup := router.Group("/user", middleware.StatCost())
	// {
	// 	userGroup.GET("/index", handlers.GetUser())
	// }

	return router
}

// SetupRouterWithDI 设置路由（带依赖注入）
func SetupRouterWithDI(userHandler *handlers.UserHandler) *gin.Engine {
	router := gin.Default()
	basePath := getCurrentPath()
	basePath = filepath.Dir(filepath.Dir(basePath))

	// 模板和静态文件设置
	handlers.SetupTemplates(router, basePath)

	// API 路由组
	apiGroup := router.Group("/api/v1")
	{
		// 用户相关路由
		users := apiGroup.Group("/users")
		{
			users.POST("", userHandler.CreateUser())       // POST /api/v1/users
			users.GET("", userHandler.GetAllUsers())       // GET /api/v1/users
			users.GET("/:id", userHandler.GetUser())       // GET /api/v1/users/:id
			users.PUT("/:id", userHandler.UpdateUser())    // PUT /api/v1/users/:id
			users.DELETE("/:id", userHandler.DeleteUser()) // DELETE /api/v1/users/:id
		}
	}

	return router
}

// Router01 返回第一个额外服务器的路由
func Router01() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/router01", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":  http.StatusOK,
			"error": "Welcome router01",
		})
	})
	return e
}

// Router02 返回第二个额外服务器的路由
func Router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/router02", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":  http.StatusOK,
			"error": "Welcome router02",
		})
	})
	return e
}
