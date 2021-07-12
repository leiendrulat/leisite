package special

import (
	"fmt"
	"net/http"
	"strconv"

	user "github.com/leiendrulat/leisite/router/handler/db/user"

	"github.com/labstack/echo/v4"
)

func Special(c echo.Context) error {
	id := c.Param("id")

	ids, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	fname, lname, id := user.GetUser(ids)
	return c.Render(http.StatusOK, "special.html", map[string]interface{}{
		"id":    id,
		"fname": fname,
		"lname": lname,
	})
}
