package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	pb "gin/pkg/pb/echo"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/proto"
)

// loadTemplates 模板继承
// templates文件夹下有以下模板文件，其中home.tmpl和index.tmpl继承了base.tmpl
func loadTemplates(templateDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	layouts, err := filepath.Glob(templateDir + "/layouts/*.tmpl")
	if err != nil {
		panic(err)
	}
	includes, err := filepath.Glob(templateDir + "/includes/*.tmpl")
	if err != nil {
		panic(err)
	}
	pages, err := filepath.Glob(templateDir + "/*.html") // 根目录的 html
	if err != nil {
		panic(err)
	}
	// 简单页面直接注册
	for _, p := range pages {
		r.AddFromFiles(filepath.Base(p), p)
	}
	// 为layouts 和 includes 目录生成 templates map
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
func indexFunc(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"Now":   time.Now().Format("2006-01-02 15:04:05"),
		"Items": []string{"Gin 框架", "模板继承", "multitemplate", "静态文件"},
	})
}
func homeFunc(c *gin.Context) {
	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"Now": time.Now().Format("2006-01-02 15:04:05"),
	})
}

// Login ShouldBind 会按顺序合并 Query + Form + JSON + Header + Uri，一份代码多端通用
type Login struct {
	Username string `form:"username" json:"username" binding:"required" header:"username" uri:"user"`
	Password string `form:"password" json:"password"  binding:"required" header:"password" uri:"user"`
}

// 返回 main.go 所在的文件夹
func getCurrentPath() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Dir(file)
}
func main() {
	basePath := getCurrentPath()

	router := gin.Default()
	// 使用记录响应体的中间件
	router.Use(ginBodyLogMiddleware)

	router.GET("/hello", func(c *gin.Context) {
		c.JSONP(200, gin.H{
			"message": "hello world",
		})
	})
	router.LoadHTMLGlob(filepath.Join(basePath, "templates", "**", "*"))
	router.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "post/index",
		})
	})
	router.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index",
		})
	})

	//自定义模板函数
	router.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	router.LoadHTMLFiles(filepath.Join(basePath, "templates", "index.tmpl"))
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", "<a href='https://yw.com'>yw的博客</a>")
	})
	// 静态文件的处理
	router.Static("/static", filepath.Join(basePath, "static"))
	// 使用模板继承
	router.HTMLRender = loadTemplates(filepath.Join(basePath, "templates"))
	router.GET("/v1/index", indexFunc)
	router.GET("/v1/home", homeFunc)
	// JSON渲染
	router.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "SomeJSON",
		})
	})
	router.GET("moreJSON", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string `json:"message"`
			Age     int    `json:"age"`
		}
		msg.Name = "syg"
		msg.Age = 24
		msg.Message = "hello world"
		c.JSON(http.StatusOK, msg)
	})
	router.GET("/pb", func(c *gin.Context) {
		resp := &pb.EchoResp{
			Label: "test",
			Nums:  []int64{1, 2},
		}
		out, _ := proto.Marshal(resp)
		c.Data(http.StatusOK, "application/x-protobuf", out)
	})
	// 获取 queryString 参数, 拿的是？后面的参数
	router.GET("/user/search", func(c *gin.Context) {
		username := c.DefaultQuery("username", "")
		address := c.Query("address")
		c.JSON(http.StatusOK, gin.H{
			"message":  "hello world",
			"username": username,
			"address":  address,
		})
	})
	// 获取 form 参数
	router.POST("posts/search", func(c *gin.Context) {
		postName := c.PostForm("postName")
		content := c.PostForm("content")
		c.JSON(http.StatusOK, gin.H{
			"message":  "hello world",
			"postName": postName,
			"content":  content,
		})
	})
	// 获取 JSON 参数
	router.POST("/json", func(c *gin.Context) {
		reqBody, _ := c.GetRawData() // 从c.Request.Body读取请求数据
		var msg map[string]interface{}
		_ = json.Unmarshal(reqBody, &msg)
		c.JSON(http.StatusOK, gin.H{})
	})
	// 路径参数
	router.GET("/user/get/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		c.JSON(http.StatusOK, gin.H{
			"message":  "hello world",
			"username": username,
			"address":  address,
		})
	})
	// 参数绑定：基于请求的Content-Type识别请求数据类型并利用反射机制自动提取请求中QueryString、form表单、JSON、XML等参数到结构体
	// 绑定JSON的示例 ({"username": "q1mi", "password": "123456"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBind(&login); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{
			"username": login.Username,
			"password": login.Password,
		})
	})
	// 绑定表单格式实例（username=webster&password=syg）
	router.POST("/loginForm", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBind(&login); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{
			"username": login.Username,
			"password": login.Password,
		})
	})
	// 绑定QueryString示例
	router.GET("/loginQuery", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBind(&login); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{
			"username": login.Username,
			"password": login.Password,
		})

	})
	// 请求头中 HTTP 标准把 X-* 留给“用户自定义扩展”，不会与未来官方头冲突，所以参数是 X-Password
	router.GET("/loginHeader", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBind(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":      "from header",
			"username": login.Username,
			"password": login.Password,
		})
	})
	router.GET("/loginUri/:id", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBind(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":      "from uri",
			"id":       "123",
			"username": login.Username,
			"password": login.Password,
		})
	})
	// 文件上传
	router.GET("/upload/page", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload_file.html", nil)
	})
	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error()})
			return
		}
		dst := fmt.Sprintf("C:/tmp/%s", file.Filename)
		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "upload success",
		})
	})
	// 多个文件上传
	// 渲染上传页
	router.GET("/upload/multi/page", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload_multi.html", nil)
	})
	router.POST("/upload/multi", func(c *gin.Context) {
		// 1.解析multiform
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// 2. 提取文件数组
		files := form.File["files"]
		if len(files) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no files uploaded"})
			return
		}
		// 3. 保存文件目录
		err = os.MkdirAll("C:/tmp/upload", 0755)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "make dir error"})
			return
		}
		var name []string
		for _, file := range files {
			dst := filepath.Join("C:/tmp/upload", file.Filename)
			if err := c.SaveUploadedFile(file, dst); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			name = append(name, file.Filename)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("upload success, %d files", len(name)),
			"name":    name,
		})
	})
	// HTTP 重定向
	router.GET("/http/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.sogo.com/")
	})
	// 路由重定向
	router.GET("/path/redirect", func(c *gin.Context) {
		c.Request.URL.Path = "/upload/page"
		router.HandleContext(c)
	})
	// 匹配所有请求方法的路由
	router.Any("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	// 为没有配置处理函数的路由添加处理程序
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload_file.html", nil)
	})
	// 路由组
	userGroup := router.Group("/user", StatCost())
	{
		userGroup.GET("/index", getUser())
	}

	// gin中间件中使用goroutine
	g := new(errgroup.Group) // 每个进程启动时拥有独立的 Group 实例
	srv8080 := &http.Server{Addr: ":8080", Handler: router}
	server01 := &http.Server{
		Addr:         ":8081",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	server02 := &http.Server{
		Addr:         ":8082",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	g.Go(func() error {
		return srv8080.ListenAndServe()
	})
	g.Go(func() error {
		return server01.ListenAndServe()
	})
	g.Go(func() error {
		return server02.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

func router02() http.Handler {
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

func router01() http.Handler {
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

func getUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello world"})
	}
}

// StatCost 记录接口耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// 通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		c.Set("name", "webster")
		// 1. 先记录要调用的函数名
		handlerName := c.HandlerName() // 例如 "main.getUser"
		path := c.Request.URL.Path     // 例如 "/user/index"

		c.Next() // 真正执行业务 handler
		cost := time.Since(start)

		fmt.Printf("[StatCost] %s | %s | %v\n", path, handlerName, cost)
	}
}

type BodyLogWriter struct {
	gin.ResponseWriter               // 嵌入gin框架ResponseWriter
	body               *bytes.Buffer // 记录用的response
}

func (w BodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func ginBodyLogMiddleware(c *gin.Context) {
	bodyLogWriter := &BodyLogWriter{
		body:           bytes.NewBuffer([]byte{}),
		ResponseWriter: c.Writer,
	}
	c.Writer = bodyLogWriter
	c.Next()
	fmt.Printf("response body is : %v\n", bodyLogWriter.body.String())
}
