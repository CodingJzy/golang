package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("templates/multi_upload.html")

	r.GET("/upload", func(c *gin.Context) {
		c.HTML(200, "multi_upload.html", gin.H{})
	})

	r.POST("/upload", func(c *gin.Context) {
		// 获取form
		form, _ := c.MultipartForm()
		// 获取form文件的map
		log.Println(form.File)
		// 获取所有文件
		files, _ := form.File["file"]
		// 遍历文件
		num := 1
		for _, file := range files {
			// 保存
			index := fmt.Sprintf("%d、", num)
			filePath := `D:\workspace\go\study\gin\example\file\` + index + file.Filename
			_ := c.SaveUploadedFile(file, filePath)
			num++
		}
		c.String(200, "upload complete...")
	})

	r.Run()
}
