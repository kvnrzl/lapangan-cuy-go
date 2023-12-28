package controller

import (
	"com.lapangan.cuy/cmd/domain/user/service"
	"github.com/labstack/echo/v4"
)

type UserControllerImpl struct {
	UserService service.UserServiceImpl
}

func (u *UserControllerImpl) Login(c echo.Context) error {

	//u.UserService.Login(ctx)

	return nil
}
