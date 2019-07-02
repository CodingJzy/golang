package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/post", func(c *gin.Context) {
		// url_path = /post?id=1234&page=1
		id := c.Query("id")
		page := c.DefaultQuery("page", "1") // 默认值
		name := c.PostForm("name")
		message := c.DefaultPostForm("message", "default message")
		c.JSON(200, gin.H{
			"id":      id,
			"page":    page,
			"name":    name,
			"message": message,
		})
	})
	r.Run()
}
