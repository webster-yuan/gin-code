package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"runtime"
	"time"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
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

// 返回 main.go 所在的文件夹
func getCurrentPath() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Dir(file)
}
func main() {
	basePath := getCurrentPath()

	router := gin.Default()
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
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
