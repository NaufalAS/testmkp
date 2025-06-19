package router

import (
	"mkptest/app"
	jadwalcontroller "mkptest/controller/jadwal"
	usercontroller "mkptest/controller/user"
	"mkptest/helper"
	"mkptest/model"
	jadwalrepo "mkptest/repository/jadwal"
	userrepo "mkptest/repository/user"
	jadwalservice "mkptest/service/jadwal"
	userservice "mkptest/service/user"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func Approuter(prefix string, e *echo.Echo) {
	db := app.DBConnection()
	tokenUseCase := helper.NewTokenUseCase()

	userAuthRepo := userrepo.NewUserRepository(db)
	userAuthService := userservice.NewUserService(userAuthRepo, tokenUseCase)
	userAuthController := usercontroller.NewUserController(userAuthService)

	jadwalRepo := jadwalrepo.NewJadwalRepository(db)
	jadwalService := jadwalservice.NewJadwalService(jadwalRepo)
	jadwalController := jadwalcontroller.NewJadwalController(jadwalService)


	g := e.Group(prefix)

	authRoute := g.Group("/auth")
	authRoute.POST("/register", userAuthController.PostUserController)
	authRoute.POST("/login", userAuthController.LoginUser)

	jadwalRoute := g.Group("/jadwal")
	jadwalRoute.POST("/save", jadwalController.JadwalController, JWTProtection())
	jadwalRoute.DELETE("/:id", jadwalController.DeleteJadwal, JWTProtection())
	jadwalRoute.GET("/list", jadwalController.GetListJadwalController)
	jadwalRoute.GET("/:id", jadwalController.DeletGetJadwalByIdControllereJadwal)
}

func JWTProtection() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(helper.JwtCustomClaims)
		},
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, model.ResponseToClient(http.StatusUnauthorized, false, "unauthorized", nil))
		},
	})
}
