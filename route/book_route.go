package route

import (
	"github.com/labstack/echo"
	"github.com/prongbang/arch/handler"
)

type BookRouter struct {
	Handle *handler.BookHandler
}

func (r BookRouter) Initial(e *echo.Echo) {
	e.GET("/books", r.Handle.GetAll)
	e.GET("/book/:id", r.Handle.GetById)
	e.PUT("/book/:id", r.Handle.Update)
	e.PATCH("/book/:id", r.Handle.Update)
	e.POST("/book", r.Handle.Create)
	e.DELETE("/book/:id", r.Handle.Delete)
}
