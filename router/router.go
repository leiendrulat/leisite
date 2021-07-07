package router

import (
	"github.com/labstack/echo/v4"
	Home "github.com/leiendrulat/leisite/router/handler/serve"
)

func Router(e *echo.Echo) {
	e.GET("/hi", Home.Home)

}
