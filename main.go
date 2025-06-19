package main

import (
	"fmt"
	"log"
	"mkptest/helper"
	"mkptest/router"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cV *CustomValidator) Validate(i interface{}) error {
	return cV.validator.Struct(i)
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("error loading .env file!")
	}

	r := echo.New()
	r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
	}))

	r.HTTPErrorHandler = helper.BindAndValidate

	router.Approuter("/api", r)

	r.Logger.Fatal(r.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
