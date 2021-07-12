package form

import (
	"fmt"
	"net/http"

	user "github.com/leiendrulat/leisite/router/handler/db/user"

	"github.com/labstack/echo/v4"
)

func ProcessForm(c echo.Context) error {
	fname := c.FormValue("fname")
	lname := c.FormValue("lname")
	email := c.FormValue("email")
	password := c.FormValue("password")
	u := user.CreateUser(fname, lname, email, password)
	fmt.Println(u)
	return c.Render(http.StatusOK, "welcome.html", map[string]interface{}{
		"name":  u.Fname,
		"email": u.Email,
	})

}
