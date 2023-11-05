package validations

import (
	"gopkg.in/validator.v2"
)

type UserLogin struct {
	Email string `json:"email" validate:"nonzero, regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$"`
	Password string `json:"password" validate:"min=5, regexp=^[0-9a-zA-Z]*$"`
}

func UserLoginValidade(userLogin *UserLogin) error {
	if err := validator.Validate(userLogin); err != nil {
		return err
	}
	return nil
}
