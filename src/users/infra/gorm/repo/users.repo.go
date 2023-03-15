package usersRepo

import (
	genericRepo "backend-skeleton-golang/commons/infra/gorm/repo"
	usersModel "backend-skeleton-golang/users/infra/models"

	"gorm.io/gorm"
)

type Users struct {
	genericRepo.Generic[usersModel.User]
}

func New(DB *gorm.DB) *Users {
	repo := &Users{genericRepo.Generic[usersModel.User]{DB: DB}}
	return repo
}


