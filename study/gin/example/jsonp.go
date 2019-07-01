package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	r.GET("/jsonp", func(c *gin.Context) {
		//data := map[string]interface{}{
		//	"foo": "bar",
		//}

		// http://127.0.0.1:8080/jsonp?callback=hello
		// hello({"foo":"bar"})
		c.JSONP(http.StatusOK, gin.H{
			"foo": "bar",
		})
	})
	r.Run()

}
