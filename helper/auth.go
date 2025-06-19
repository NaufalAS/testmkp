package helper

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func GetAuthId(c echo.Context) (int, error) {
	claims := c.Get("user").(*jwt.Token).Claims.(*JwtCustomClaims)
	authID, err := strconv.Atoi(claims.ID)

	if err != nil {
		return -1, err
	}

	return authID, nil
}

func Login(userId int, userEmail string) (map[string]interface{}, error) {
	var tokenDuration time.Duration
	expiredTime := time.Now().Local().Add(tokenDuration)

	claims := JwtCustomClaims{
		ID:    strconv.Itoa(userId),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token, err := NewTokenUseCase().GenerateAccessToken(claims)
	if err != nil {
		return nil, errors.New("internal server error")
	}

	return map[string]interface{}{
		"token":      token,
		"expires_at": expiredTime.Format(time.RFC3339),
	}, nil
}