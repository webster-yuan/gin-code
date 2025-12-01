package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// UploadPageHandler 处理上传页面路由
func UploadPageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload_file.html", nil)
	}
}

// UploadHandler 处理文件上传路由
func UploadHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
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
	}
}

// UploadMultiPageHandler 处理多文件上传页面路由
func UploadMultiPageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload_multi.html", nil)
	}
}

// UploadMultiHandler 处理多文件上传路由
func UploadMultiHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
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
	}
}
