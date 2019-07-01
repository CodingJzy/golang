package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//加载模版文件(匹配的方式)
	r.LoadHTMLGlob("templates/*")
	//加载模版文件(字符串方式)
	r.LoadHTMLFiles("templates/index.html", "templates/login.html")

	r.GET("/render", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "welcome login ...",
		})
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{})
	})

	r.Run()
}
