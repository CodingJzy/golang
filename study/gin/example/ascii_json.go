package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	mp := map[string]interface{}{
		"lang": "Go语言",
		"tag":  "<br>",
	}
	r.GET("/ascii.json", func(c *gin.Context) {
		c.AsciiJSON(http.StatusOK, mp)
	})
	r.Run()
}
