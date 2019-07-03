package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// /query_param?fastname=子牙&lastname=江
	r.GET("/query_param", func(c *gin.Context) {
		// 没有就默认
		firstName := c.DefaultQuery("firstname", "子牙")
		lastName := c.Query("lastname")
		c.String(200, "hello %s %s", firstName, lastName)
	})
	r.Run()
}
