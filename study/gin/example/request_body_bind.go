package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type formA struct {
	Foo string `json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `json:"bar" xml:"bar" binding:"required"`
}

func main() {
	r := gin.Default()

	r.POST("/request_bind_body", func(c *gin.Context) {
		objA := formA{}
		objB := formB{}
		fmt.Println(objA, objB)
		// 读取c.Request.Body 并将结果存入上下文。
		if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
			c.String(200, "the body should be formA")
		} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
			c.String(200, "the body is should be formB")
		} else if errB2 := c.ShouldBindBodyWith(&objB, binding.XML); errB2 == nil {
			c.String(http.StatusOK, `the body should be formB XML`)
		} else {
			c.String(400, "not found")
		}
	})
	r.Run()
}
