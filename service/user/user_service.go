package userservice

import "mkptest/model/web"

type UserService interface {
	LoginUser(email string, password string) (map[string]interface{}, error)
	SaveUser(request web.RegisterUserRequest) (map[string]interface{}, error)
}