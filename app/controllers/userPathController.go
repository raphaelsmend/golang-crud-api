package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-sample-store/database"
	"go-sample-store/models"
)

func UserPath(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	db := database.ConnectDB()

	db.First(&user, id)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user not find"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := models.UserValidade(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	db.Save(user);

	c.JSON(http.StatusOK, user)
}