package middlewares

import (
	"os"
	"fmt"
	"github.com/gin-gonic/gin"
    "net/http"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"go-sample-store/models"
	"go-sample-store/database"
)

func ValidateToken(c *gin.Context) {
	getClaimsToken(c)
	c.Next()
}

func GetUserByToken(c *gin.Context) *models.User {
	claims := getClaimsToken(c)
	id := claims["id"].(float64)
	var user models.User
	db := database.ConnectDB()

	db.First(&user, id)

	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "token user not find"})
		c.Abort()	
		return nil
	}

	return &user
}

func getClaimsToken(c *gin.Context) (jwt.MapClaims) {
	authorizationHeader := c.GetHeader("Authorization")
	parts := strings.Split(authorizationHeader, " ")
    if len(parts) < 2 {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
		c.Abort()
        return nil
    }

	tokenString := parts[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verifique o método de assinatura
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Método de assinatura inválido")
		}
		jwtKey := os.Getenv("JWT_SECRET_KEY")
		return []byte(jwtKey), nil
	})
	
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized parse token"})
		c.Abort()
		return nil
	}
	
	if token.Valid {	
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. parse claims"})
			c.Abort()
			return nil
		}
		return claims
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	c.Abort()
	return nil
}
