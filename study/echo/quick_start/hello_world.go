package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	// 实例化一个echo示例
	e := echo.New()
	// 执行GET方法。
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello echo")
	})
	e.Logger.Fatal(e.Start(":1323"))

}
