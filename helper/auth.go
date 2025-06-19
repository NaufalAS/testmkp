package helper

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func Login(userId int, userEmail string) (map[string]interface{}, error) {
	tokenDuration := time.Hour * 1 
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