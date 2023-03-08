package usersDTO

import usersDomain "backend-skeleton-golang/users/domain"

type UserRes struct {
	Id        string `json:"id"`
	Email     string `json:"email" `
	FirstName string `json:"firstName" `
	LastName  string `json:"lastName" `
}

type Create struct {
	usersDomain.User
}

type Update struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}
