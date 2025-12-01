package handlers

import (
	"net/http"

	"gin/internal/models"

	"github.com/gin-gonic/gin"
)

// UserSearchHandler 处理用户搜索路由
func UserSearchHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.DefaultQuery("username", "")
		address := c.Query("address")
		c.JSON(http.StatusOK, gin.H{
			"message":  "hello world",
			"username": username,
			"address":  address,
		})
	}
}

// PostsSearchHandler 处理帖子搜索路由
func PostsSearchHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		postName := c.PostForm("postName")
		content := c.PostForm("content")
		c.JSON(http.StatusOK, gin.H{
			"message":  "hello world",
			"postName": postName,
			"content":  content,
		})
	}
}

// UserGetHandler 处理获取用户路由
func UserGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		c.JSON(http.StatusOK, gin.H{
			"message":  "hello world",
			"username": username,
			"address":  address,
		})
	}
}

// LoginJSONHandler 处理JSON登录路由
func LoginJSONHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var login models.Login
		if err := c.ShouldBind(&login); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{
			"username": login.Username,
			"password": login.Password,
		})
	}
}

// LoginFormHandler 处理表单登录路由
func LoginFormHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var login models.Login
		if err := c.ShouldBind(&login); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{
			"username": login.Username,
			"password": login.Password,
		})
	}
}

// LoginQueryHandler 处理查询参数登录路由
func LoginQueryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var login models.Login
		if err := c.ShouldBind(&login); err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{
			"username": login.Username,
			"password": login.Password,
		})
	}
}

// LoginHeaderHandler 处理头部登录路由
func LoginHeaderHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var login models.Login
		if err := c.ShouldBind(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":      "from header",
			"username": login.Username,
			"password": login.Password,
		})
	}
}

// LoginUriHandler 处理URI登录路由
func LoginUriHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var login models.Login
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
	}
}
