package home

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Test(c echo.Context) error {
	id := c.Param("id")

	return c.Render(http.StatusOK, "test.html", map[string]interface{}{
		"id": id,
	})
}
