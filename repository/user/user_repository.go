package userrepo

import "mkptest/model/domain"

type UserRepository interface {
	LoginEmail(email string) (*domain.User, error)
	SaveUser(user domain.User) (domain.User, error)
}