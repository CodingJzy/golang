package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {

	route := gin.Default()
	route.POST("/login", func(c *gin.Context) {
		var form LoginForm
		var bind = c.ShouldBind(&form)
		// bind为nil时表示前端传来的字段正确，但不代表值正确
		fmt.Println(bind)
		if bind == nil {
			if form.User == "jiang_wei" && form.Password == "jw19961019" {
				c.JSON(200, gin.H{
					"status": "login success",
				})
			} else {
				c.JSON(401, gin.H{
					"status": "unauthorized",
				})
			}
		}
	})
	route.Run()
}
