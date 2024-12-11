package usecases

import (
	"gm-startd/database"
	"gm-startd/entities"
)

type userUseCase struct {
	DB database.UserDB
}

func Init(userRepo database.UserDB) *userUseCase {
	return &userUseCase{
		DB: userRepo,
	}
}

func (uc *userUseCase) GetAll() (*entities.Users, error) {
	return uc.DB.GetAll()
}

func (uc *userUseCase) GetByID(id int) (*entities.User, error) {
	return uc.DB.GetByID(id)
}

func (uc *userUseCase) CreateNew(user *entities.User) error {
	return uc.DB.CreateNew(user)
}

func (uc *userUseCase) DeleteByID(id int) error {
	return uc.DB.DeleteByID(id)
}
