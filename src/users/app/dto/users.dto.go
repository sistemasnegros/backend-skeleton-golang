package usersDTO

type UserRes struct {
	Id        string `json:"id"`
	Email     string `json:"email" `
	FirstName string `json:"firstName" `
	LastName  string `json:"lastName" `
}

type Create struct {
	Id        string `json:"id" validate:"required,objectId" `
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

type Update struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}
