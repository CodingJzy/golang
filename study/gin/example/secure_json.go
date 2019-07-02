package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/secure_json", func(c *gin.Context) {
		c.SecureJSON(200, []string{"Go", "Python", "Linux", "Java"})
	})
	r.Run()
}
