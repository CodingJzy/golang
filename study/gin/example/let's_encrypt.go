package main

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	// ping
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "ping")
	})
	log.Fatal(autotls.Run(r, "example1.com", "example2.com"))
}
