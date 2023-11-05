package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-sample-store/database"
	"go-sample-store/models"
	"go-sample-store/validations"
	"go-sample-store/resources"
)

func UserLogin(c *gin.Context) {
	var userLogin validations.UserLogin
	
	if err := c.ShouldBindJSON(&userLogin); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	if err := validations.UserLoginValidade(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	var user models.User

	db := database.ConnectDB()
	db.First(&user, "email = ?", userLogin.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not find"})
        return
	}

	if user.Password != userLogin.Password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password incorrect"})
        return
	}

	tokenString, err := models.UserBuildTokenJwt(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
        return
	}

	c.JSON(http.StatusOK, resources.BuildUserLoginResource(user, tokenString))
}
