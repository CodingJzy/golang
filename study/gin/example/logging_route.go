package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v \n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	r.GET("/f1", func(c *gin.Context) {
		c.JSON(http.StatusOK, "f1")
	})
	r.GET("/f2", func(c *gin.Context) {
		c.JSON(http.StatusOK, "f2")
	})
	r.GET("/f3", func(c *gin.Context) {
		c.JSON(http.StatusOK, "f3")
	})
	r.Run()
}
