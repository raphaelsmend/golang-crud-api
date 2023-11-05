package resources

import (
	"go-sample-store/models"
)

type userLoginResource struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func BuildUserLoginResource(user models.User, tokenString string) userLoginResource {
	userLoginResource := userLoginResource {
		ID: int(user.ID),
		Name: user.Name,
		Email: user.Email,
		Token: tokenString,
	}

	return userLoginResource
}
