package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

func main() {
	r := gin.Default()

	//自定义模版渲染器
	html := template.Must(template.ParseFiles("templates2/zidingyi.html"))
	r.SetHTMLTemplate(html)
	r.GET("/zidingyi", func(c *gin.Context) {
		c.HTML(200, "zidingyi.html", gin.H{
			"title": "自定义",
		})
	})
	r.Run()
}
