package usercontroller

import (
	"mkptest/model"
	"mkptest/model/web"
	userservice "mkptest/service/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserControllerImpl struct {
	UserService userservice.UserService
}

func NewUserController(userService userservice.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) PostUserController(c echo.Context) error {
	user := new(web.RegisterUserRequest)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
	}

	if err := c.Validate(user); err != nil {
		return err
	}

	saverUser, errSaveuser := controller.UserService.SaveUser(*user)
	if errSaveuser != nil {
		return c.JSON(http.StatusInternalServerError, model.ResponseToClient(http.StatusInternalServerError, false, errSaveuser.Error(), nil))
	}
	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true, "Register Sukses", saverUser))
}

func (controller *UserControllerImpl) LoginUser(c echo.Context) error {
	user := new(web.LoginUserRequest)

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest, false, err.Error(), nil))
	}

	userRes, errLogin := controller.UserService.LoginUser(user.Email, user.Password)
	if errLogin != nil {
		return c.JSON(http.StatusBadRequest, model.ResponseToClient(http.StatusBadRequest,false, errLogin.Error(), nil))
	}

	return c.JSON(http.StatusOK, model.ResponseToClient(http.StatusOK, true,  "login berhasil", userRes))
}