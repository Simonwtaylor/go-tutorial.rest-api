package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

var data []ExampleData

func getAll(c echo.Context) error {
	return c.JSON(http.StatusOK, data)
}

func createOne(c echo.Context) error {
	body, err := convertToBody(c)

	if err != nil {
		return c.JSON(http.StatusNotAcceptable, "Content not acceptable")
	}

	data = append(data, ExampleData{ID: len(data) + 1, Content: body.Content})
	return c.JSON(http.StatusOK, data)
}

func getOne(c echo.Context) error {
	id := c.Param("id")

	for _, val := range data {
		if fmt.Sprintf("%v", val.ID) == id {
			return c.JSON(http.StatusOK, val)
		}
	}

	return c.JSON(http.StatusNotFound, nil)
}

func updateOne(c echo.Context) error {
	id := c.Param("id")

	for i, val := range data {
		if fmt.Sprintf("%v", val.ID) == id {

			body, err := convertToBody(c)

			if err != nil {
				return c.JSON(http.StatusNotAcceptable, "Content not acceptable")
			}

			data[i].Content = body.Content

			return c.JSON(http.StatusAccepted, data[i])
		}
	}

	return c.JSON(http.StatusNotFound, nil)
}

func deleteOne(c echo.Context) error {
	id := c.Param("id")

	for i, val := range data {
		if fmt.Sprintf("%v", val.ID) == id {
			data = RemoveExampleDataIndex(data, i)

			return c.JSON(http.StatusOK, nil)
		}
	}

	return c.JSON(http.StatusNotFound, nil)
}

func convertToBody(c echo.Context) (*PostArg, error) {
	body := new(PostArg)

	err := c.Bind(body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
