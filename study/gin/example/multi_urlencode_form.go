package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/form_post", func(c *gin.Context) {
		// 获取form表单传来的值
		message := c.PostForm("message")
		fmt.Println(message)
		// 获取form表单传来的值,没传就是默认值
		nick := c.DefaultPostForm("nick", "anonymous")
		fmt.Println(nick)

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	r.Run()
}
