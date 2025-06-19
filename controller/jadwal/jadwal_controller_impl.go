package jadwalcontroller

import "github.com/labstack/echo/v4"

type UserController interface {
	JadwalController(c echo.Context) error
	DeleteJadwal(c echo.Context) error
	GetListJadwalController(c echo.Context) error
	GetJadwalByIdController(c echo.Context) error
}