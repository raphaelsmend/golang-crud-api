package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-sample-store/database"
	"go-sample-store/models"
)

func UserGet(c *gin.Context) {
	var users []models.User

	db := database.ConnectDB()
	db.Find(&users)
	c.JSON(http.StatusOK, users)
}