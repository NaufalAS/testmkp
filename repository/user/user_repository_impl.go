package userrepo

import (
	"mkptest/model/domain"

	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{
		DB: db,
	}
}

func (repo *AuthRepositoryImpl) LoginEmail(email string) (*domain.User, error){
	user := new(domain.User)
	if err := repo.DB.Where("email = ?", email).Take(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
func (repo *AuthRepositoryImpl) SaveUser(user domain.User) (domain.User, error) {
	err := repo.DB.Create(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}