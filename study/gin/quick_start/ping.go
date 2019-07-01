package main

import "github.com/gin-gonic/gin"

func main() {
	// 实例化
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"code":    200,
		})
	})
	// 默认监听在0.0.0.0:8080
	r.Run()
}
