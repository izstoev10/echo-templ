package handler

import (
	"github.com/izstoev10/echo-templ/model"
	"github.com/izstoev10/echo-templ/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Email: "test@gmail.com",
	}
	return render(c, user.Show(u))
}
