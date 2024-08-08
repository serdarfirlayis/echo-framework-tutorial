package handler

import (
	"github.com/serdarfirlayis/echo-framework-tutorial.git/echo-with-frontend-basics/model"
	"github.com/labstack/echo/v4"
	"github.com/serdarfirlayis/echo-framework-tutorial.git/echo-with-frontend-basics/view/user"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User {
		Email: "admin@gmail.com",
	}
	return render(c, user.Show(u))
}
