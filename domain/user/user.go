package user

import "ecomerce-go/model"

type UseCase interface {
	Create(m * model.User) error
	GetByEmail(email string) (model.User, error)
	GetAll() (model.Users, error)
}

type Repository interface {
	Create(m * model.User) error
	GetByEmail(email string) (model.User, error)
	GetAll() (model.Users, error)
}
