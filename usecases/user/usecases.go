package usecases

import "gm-startd/entities"

type UserUseCase interface {
	GetAll() (*entities.Users, error)
	GetByID(id int) (*entities.User, error)
	CreateNew(user *entities.User) error
	DeleteByID(id int) error
}
