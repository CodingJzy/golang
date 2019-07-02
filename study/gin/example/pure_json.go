package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 提供unicode实体：{"html":"\u003cb\u003eHello Go---江子牙\u003c/b\u003e","test":"json"}
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello Go---江子牙</b>",
			"test": "json",
		})
	})

	// 提供字面字符：{"html":"<b>Hello Go---江子牙</b>","test":"pure_json"}
	r.GET("/pure_json", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello Go---江子牙</b>",
			"test": "pure_json",
		})
	})

	r.Run()
}
