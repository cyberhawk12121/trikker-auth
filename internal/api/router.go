package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"module github.com/cyberhawk12121/trikker-auth/internal/service"
)

func SetupRouter(authService *service.AuthService) *gin.Engine {
	router := gin.Default()

	// Public routes
	router.POST("/register", registerHandler(authService))
	router.POST("/login", loginHandler(authService))

	return router
}

func registerHandler(authService *service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Registering user", c.Request.Body)
		user, err := authService.RegisterUser(c)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, user)
	}
}

func loginHandler(authService *service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := authService.LoginUser(c)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, user)
	}
}
