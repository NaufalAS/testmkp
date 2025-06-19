package usercontroller

import "github.com/labstack/echo/v4"

type UserController interface {
	PostUserController(c echo.Context) error
	LoginUser(c echo.Context) error
}