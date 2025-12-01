package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTPRedirectHandler 处理HTTP重定向路由
func HTTPRedirectHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.sogo.com/")
	}
}

// PathRedirectHandler 处理路径重定向路由
func PathRedirectHandler(router *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.URL.Path = "/upload/page"
		router.HandleContext(c)
	}
}
