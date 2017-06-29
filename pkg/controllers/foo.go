package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

//Foo Dependencies and Configs
type Foo struct {
}

//Standard controller register function
func RegisterFoo(e *echo.Echo) Foo {
	foo := Foo{}

	e.GET("/foo/:id", foo.getFoo)

	return foo
}

func (f *Foo) getFoo(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}


