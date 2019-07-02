package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 模拟一些私人数据
var secrets = gin.H{
	"user1": gin.H{"email": "user1@qq.com", "phone": "1"},
	"user2": gin.H{"email": "user2@qq.com", "phone": "2"},
	"user3": gin.H{"email": "user3@qq.com", "phone": "3"},
}

func main() {
	r := gin.Default()

	// 路由组使用gin.BasicAuth()中间件
	// gin.Accounts其实和gin.H{}一样都是map的一种快捷方式(类型定义)。
	// username:password
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"user1": "jiang_wei1",
		"user2": "jiang_wei2",
		"user3": "jiang_wei3",
	}))

	authorized.GET("/secrets", func(c *gin.Context) {
		// 获取用户，它是由BasicAuth中间件设置的
		user := c.MustGet(gin.AuthUserKey).(string)
		fmt.Println(user + "________________________")
		// 去模拟的数据里取值，然后返回
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{
				"user":   user,
				"secret": secret,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})
	// http://127.0.0.1:8080/admin/secrets
	// 弹出一窗，让登录
	// 输入密码返回模拟的数据和有关湖面
	r.Run()
}
