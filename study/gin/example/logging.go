package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	// 禁言控制台颜色，将日志写入文件时不需要控制台颜色
	gin.DisableConsoleColor()

	// 记录到文件
	f, _ := os.Create("log/gin.log1")
	fmt.Println(f)
	gin.DefaultWriter = io.MultiWriter(f)

	// 同时将日志写入文件和控制台
	//gin.DefaultWriter = io.MultiWriter(f,os.Stdin)

	r := gin.Default()

	r.GET("/log", func(c *gin.Context) {
		c.String(200, "log")
	})
	r.Run()
}
