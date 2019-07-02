package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("templates/upload.html")
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(200, "upload.html", gin.H{})
	})

	r.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")

		//文件头
		log.Println(file.Header)
		// 文件名字
		log.Println(file.Filename)
		// 文件大小
		log.Println(file.Size)

		// 上传文件到指定目录
		filePath := `D:\workspace\go\study\gin\example\file\` + file.Filename
		err := c.SaveUploadedFile(file, filePath)
		log.Println(err)
		c.String(http.StatusOK, fmt.Sprintf("%s upload complete", file.Filename))
	})
	r.Run()

}
