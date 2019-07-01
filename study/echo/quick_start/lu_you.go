package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func getUser(c echo.Context) error {
	// url path param
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {

	// query_param
	userName := c.QueryParam("username")
	age := c.QueryParam("age")
	return c.String(200, "username："+userName+"\tage："+age)
}

func main() {
	e := echo.New()
	e.Debug = true
	e.GET("/users/:id", getUser)
	e.GET("/show", show)

	e.Logger.Fatal(e.Start(":1323"))

}
