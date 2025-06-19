package entity

import "mkptest/model/domain"

type UserEntity struct {
	UserID int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}

func ToUserEntity(user domain.User) UserEntity {
	return UserEntity{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Role:   user.Role,
	}
}

func ToUserListEntity(users []domain.User) []UserEntity {
	var userData []UserEntity

	for _, user := range users {
		userData = append(userData, ToUserEntity(user))

	}
	return userData
}
