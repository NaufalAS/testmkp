package userservice

import (
	"errors"
	"fmt"
	"mkptest/helper"
	"mkptest/model/domain"
	"mkptest/model/web"
	userrepo "mkptest/repository/user"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	userrepo userrepo.UserRepository
	tokenUseCase helper.TokenUseCase
}

func NewUserService(userrepo userrepo.UserRepository, tokenUseCase helper.TokenUseCase) *UserServiceImpl {
	return &UserServiceImpl{
		userrepo: userrepo,
		tokenUseCase: tokenUseCase,
	}
}

func (service *UserServiceImpl) SaveUser(request web.RegisterUserRequest) (map[string]interface{}, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Hashing error:", err)
		return nil, err
	}

	newUser := domain.User{
		Name:     request.Name,
		Password: string(hashedPassword),
		Email:    request.Email,
		Role:     request.Role,
	}

	saveUser, errUser := service.userrepo.SaveUser(newUser)
	if errUser != nil {
		fmt.Println("Database error:", errUser)
		return nil, errUser
	}

	data := helper.CustomResponse{
		"id":    saveUser.ID,
		"name":  saveUser.Name,
		"email": saveUser.Email,
		"role":  saveUser.Role,
	}
	return data, nil
}

func (service *UserServiceImpl) LoginUser(email string, password string) (map[string]interface{}, error) {
	user, err := service.userrepo.LoginEmail(email)
	if err != nil {
		return nil, errors.New("Email tidak ditemukan")
	}

	errPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errPass != nil {
		return nil, errors.New("Password Salah")
	}

	// Panggil helper untuk buat token dan ambil expired
	loginResponse, loginErr  := helper.Login(user.ID, user.Email)
	if loginErr != nil {
		return nil, errors.New("ada kesalahan generate token")
	}

	return helper.CustomResponse{
		"token":      loginResponse["token"],
		"expires_at": loginResponse["expires_at"],
	}, nil
}
