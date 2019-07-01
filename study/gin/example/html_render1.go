package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//使用不同目录下的模板文件
	r.LoadHTMLGlob("templates1/**/*")

	r.GET("/cat/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "cat/home.html", gin.H{
			"title": "cat",
		})
	})
	r.GET("/dog/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dog/home.html", gin.H{
			"title": "dog",
		})
	})

	r.Run()
}
