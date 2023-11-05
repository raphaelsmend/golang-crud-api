package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"go-sample-store/controllers"
	"go-sample-store/middlewares"
)

func HandleRequest() {
	r := gin.Default()
	r.Use(cors.Default())

    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Welcome")
    })

    apiGroup := r.Group("/api")

	apiGroup.POST("/user", controllers.UserPost)
    apiGroup.GET("/user", controllers.UserGet)
	apiGroup.PATCH("/user/:id", controllers.UserPath)
	apiGroup.DELETE("/user/:id", controllers.UserDelete)
	apiGroup.POST("/login", controllers.UserLogin)
	apiGroup.GET("/test-protected", middlewares.ValidateToken, controllers.TestProtected)			

	r.Run()
}
