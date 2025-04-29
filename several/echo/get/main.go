package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/v1/user/:id", func(c echo.Context) error {

		myuser := struct {
			Name string `json:"name"`
		}{
			Name: c.Param("id"),
		}
		c.Set("Content-Type", "application/json")
		return c.JSON(http.StatusOK, myuser)
	})

	e.Start(":8080")
}
