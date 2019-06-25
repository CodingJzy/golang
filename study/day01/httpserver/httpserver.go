// 声明main函数 main函数是go程序的入口
package main

// 导入net/http包
import (
	"net/http"
)

func main() {
	// 使用http.FileServer文件服务器将当前目录当作根目录的处理器
	http.Handle("/", http.FileServer(http.Dir(".")))
    // 服务器监听8080端口
	http.ListenAndServe(":8080", nil)
}
