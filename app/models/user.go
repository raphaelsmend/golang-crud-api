package models

import (
	"os"
	"time"
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	gorm.Model
	Name string `json:"name" validate:"nonzero"`
	Email string `json:"email" validate:"nonzero, regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$"`
	Password string `json:"password" validate:"min=5, regexp=^[0-9a-zA-Z]*$"`
}

func UserValidade(user *User) error {
	if err := validator.Validate(user); err != nil {
		return err
	}
	return nil
}

func UserBuildTokenJwt(user *User) (string, error) {
	// Define as claims
	claims := jwt.MapClaims{
		"id": user.ID,
		"name": user.Name,
		"email": user.Email,
		//"exp": time.Now().Add(time.Hour * 24).Unix(),
		"exp": time.Now().Add(time.Minute * 3).Unix(),
	}

	// Cria o token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Assina o token com uma chave secreta
	jwtKey := os.Getenv("JWT_SECRET_KEY")
	tokenString, err := token.SignedString([]byte(jwtKey))

	return tokenString, err
}