package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/todos", list)
	e.POST("/todos", create)
	e.GET("/todos/:id", view)
	e.PUT("/todos/:id", done)
	e.DELETE("/todos/:id", remove)

	e.Logger.Fatal(e.Start(":8080"))
}

func list(c echo.Context) error {
	return nil
}

func view(c echo.Context) error {
	return nil
}

func create(c echo.Context) error {
	return nil
}

func done(c echo.Context) error {
	return nil
}

func remove(c echo.Context) error {
	return nil
}
