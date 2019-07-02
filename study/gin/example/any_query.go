package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	r := gin.Default()

	r.Any("/any_query", func(c *gin.Context) {
		var person Person
		if c.ShouldBindQuery(&person) == nil {
			log.Println("====== Only Bind By Query String ======")
			log.Println(person.Name)
			log.Println(person.Address)
			c.String(200, "success")
		} else {
			c.String(400, "fail")
		}
	})

	r.Run()
}
