package controllers

import (
	"net/http"
	"go-sample-store/middlewares"
	"github.com/gin-gonic/gin"
)

func TestProtected(c *gin.Context) {
	user := middlewares.GetUserByToken(c)
	c.JSON(http.StatusOK, gin.H{
		"user_logged": user})
	c.Abort()	
}