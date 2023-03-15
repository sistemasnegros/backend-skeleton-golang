package usersModel

import (
	commonsInfraGorm "backend-skeleton-golang/commons/infra/gorm/models"
	usersDomain "backend-skeleton-golang/users/domain"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	usersDomain.User
	commonsInfraGorm.Default
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)

	if err != nil {
		return errors.New("err in bcrypt")
	}

	u.Password = string(bytes)

	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)

	if err != nil {
		return errors.New("err in bcrypt")
	}

	u.Password = string(bytes)

	return
}


