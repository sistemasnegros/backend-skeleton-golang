package authDTO

import (
	usersDomain "backend-skeleton-golang/users/domain"
)

type Register struct {
	usersDomain.User
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResUser struct {
	Id        string `json:"id"`
	Email     string `json:"email" `
	FirstName string `json:"firstName" `
	LastName  string `json:"lastName" `
}

type LoginRes struct {
	Token string       `json:"token"`
	User  LoginResUser `json:"user"`
}

type ForgotPassword struct {
	Email string `json:"email" validate:"required,email"`
}

type RestorePassword struct {
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
}

type UpdateMe struct {
	Email     string `json:"email" validate:"email"`
	Password  string `json:"password"  `
	FirstName string `json:"firstName" `
	LastName  string `json:"lastName" `
}
