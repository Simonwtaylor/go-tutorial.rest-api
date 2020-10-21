package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", getAll)
	e.POST("/", createOne)

	e.Logger.Fatal(e.Start(":1323"))
}
