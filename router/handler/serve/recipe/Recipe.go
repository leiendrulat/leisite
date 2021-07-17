package recipe

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Recipe(c echo.Context) error {
	id := c.Param("id")

	return c.Render(http.StatusOK, "recipe.html", map[string]interface{}{
		"id": id,
	})
}
