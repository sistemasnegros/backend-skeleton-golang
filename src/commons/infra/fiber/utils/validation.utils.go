package utilsFiberInfra

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ErrorResponse struct {
	FailedField string `json:"failedField"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func ValidateStruct(body interface{}) []*ErrorResponse {
	var validate = validator.New()

	validate.RegisterValidation("objectId", ObjectId)
	var errors []*ErrorResponse
	err := validate.Struct(body)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

// ValidateValuer implements validator.CustomTypeFunc
func ObjectId(fl validator.FieldLevel) bool {
	return primitive.IsValidObjectID(fl.Field().String())
}
