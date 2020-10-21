package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", getAll)
	e.GET("/:id", getOne)
	e.POST("/", createOne)
	e.PUT("/:id", updateOne)
	e.DELETE("/:id", deleteOne)

	e.Logger.Fatal(e.Start(":1323"))
}
