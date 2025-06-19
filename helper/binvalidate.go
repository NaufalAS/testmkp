package helper

import (
	"fmt"
	"mkptest/model"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func BindAndValidate(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				report.Message = fmt.Sprintf("%s field is required", err.Field())
				report.Code = http.StatusBadRequest
			case "email":
				report.Message = fmt.Sprintf("field is not a valid email address")
				report.Code = http.StatusBadRequest
			}
		}
	}
	c.Logger().Error(report.Message)
	c.JSON(report.Code, model.ResponseToClient(report.Code, false, report.Message.(string), nil))
}