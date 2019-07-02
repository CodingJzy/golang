package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 返回xml格式的数据
	r.GET("/xml", func(c *gin.Context) {
		c.XML(200, gin.H{
			"message": "xml",
			"status":  http.StatusOK,
		})
	})

	// 返回yaml格式的数据
	r.GET("/yaml", func(c *gin.Context) {
		c.YAML(200, gin.H{
			"message": "yaml",
			"status":  http.StatusOK,
		})
	})
	r.Run()
}
