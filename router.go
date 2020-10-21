package main

import (
	"net/http"

	"github.com/labstack/echo"
)

var data []ExampleData

func getAll(c echo.Context) error {
	return c.JSON(http.StatusOK, data)
}

func createOne(c echo.Context) error {
	body := new(PostArg)

	err := c.Bind(body)

	if err != nil {
		return c.JSON(http.StatusNotAcceptable, "Content not acceptable")
	}

	data = append(data, ExampleData{ID: len(data) + 1, Content: body.Content})
	return c.JSON(http.StatusOK, data)
}
