package usersDomain

type User struct {
	Id        string `json:"id" gorm:"primary_key;type:string" validate:"required" `
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}
