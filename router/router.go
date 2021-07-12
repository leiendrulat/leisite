package router

import (
	"github.com/labstack/echo/v4"
	Form "github.com/leiendrulat/leisite/router/handler/process/form"
	Home "github.com/leiendrulat/leisite/router/handler/serve"
	Special "github.com/leiendrulat/leisite/router/handler/serve/special"
)

func Router(e *echo.Echo) {
	e.GET("/", Home.Home)
	e.GET("/special/:id", Special.Special)

	e.POST("/form", Form.ProcessForm)
}
