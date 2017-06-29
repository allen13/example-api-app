package api

import (
	"github.com/allen13/example-api-app/pkg/controllers"
	"github.com/labstack/echo"
)

func StartAPI() {
	e := echo.New()

	controllers.RegisterFoo(e)

	e.Logger.Fatal(e.Start(":1323"))
}
