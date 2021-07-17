package router

import (
	"github.com/labstack/echo/v4"
	Form "github.com/leiendrulat/leisite/router/handler/process/form"
	Home "github.com/leiendrulat/leisite/router/handler/serve"
	Recipe "github.com/leiendrulat/leisite/router/handler/serve/recipe"
	Special "github.com/leiendrulat/leisite/router/handler/serve/special"
	Test "github.com/leiendrulat/leisite/router/handler/test"
)

func Router(e *echo.Echo) {
	e.GET("/", Home.Home)
	e.GET("/special/:id", Special.Special)
	e.GET("/test/:id", Test.Test)
	e.GET("/recipe/:id", Recipe.Recipe)

	e.POST("/form", Form.ProcessForm)
}
