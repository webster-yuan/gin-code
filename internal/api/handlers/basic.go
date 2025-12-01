package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloHandler 处理hello路由
func HelloHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSONP(200, gin.H{
			"message": "hello world",
		})
	}
}

// PostsIndexHandler 处理posts/index路由
func PostsIndexHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "post/index",
		})
	}
}

// UsersIndexHandler 处理users/index路由
func UsersIndexHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index",
		})
	}
}

// CustomTemplateHandler 处理自定义模板路由
func CustomTemplateHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", "<a href='https://yw.com'>yw的博客</a>")
	}
}

// SomeJSONHandler 处理JSON渲染路由
func SomeJSONHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "SomeJSON",
		})
	}
}

// MoreJSONHandler 处理复杂JSON渲染路由
func MoreJSONHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string `json:"message"`
			Age     int    `json:"age"`
		}
		msg.Name = "syg"
		msg.Age = 24
		msg.Message = "hello world"
		c.JSON(http.StatusOK, msg)
	}
}

// JSONHandler 处理JSON参数路由
func JSONHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqBody, _ := c.GetRawData() // 从c.Request.Body读取请求数据
		var msg map[string]interface{}
		_ = json.Unmarshal(reqBody, &msg)
		c.JSON(http.StatusOK, gin.H{})
	}
}

// TestHandler 处理test路由
func TestHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	}
}

// NoRouteHandler 处理未找到路由
func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload_file.html", nil)
	}
}
