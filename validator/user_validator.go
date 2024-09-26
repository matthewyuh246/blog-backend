package validator

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/matthewyuh246/blogbackend/models"
)

type IUserValidator interface {
	SignUpUserValidate(user models.User) error
	LoginUserValidate(user models.User) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) SignUpUserValidate(user models.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			validation.RuneLength(1, 30).Error("limited max 30 char"),
			is.Email.Error("is not valid email format"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(6, 30).Error("limited min 6 max 30 char"),
		),
		validation.Field(
			&user.FirstName,
			validation.Required.Error("FirstName is required"),
			validation.RuneLength(1, 20).Error("limited min 1 max 20 char"),
		),
		validation.Field(
			&user.LastName,
			validation.Required.Error("LastName is required"),
			validation.RuneLength(1, 20).Error("limited min 1 max 20 char"),
		),
		validation.Field(
			&user.Phone,
			validation.Required.Error("PhoneNumber is required"),
		),
	)
}

func (uv *userValidator) LoginUserValidate(user models.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			validation.RuneLength(1, 30).Error("limited max 30 char"),
			is.Email.Error("is not valid email format"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(6, 30).Error("limited min 6 max 30 char"),
		),
	)
}
