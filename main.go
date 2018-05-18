package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/prongbang/arch/handler"
	"github.com/prongbang/arch/repository"
	"github.com/prongbang/arch/route"
)

func main() {

	handle := handler.BookHandler{repository.CreateBookRepository()}

	e := echo.New()

	// Router
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	route.BookRouter{&handle}.Initial(e)

	// Listener
	e.Logger.Fatal(e.Start(":1323"))
}
