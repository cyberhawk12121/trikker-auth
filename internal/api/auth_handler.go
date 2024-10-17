package api

import (
	"github.com/gin-gonic/gin"
	"github.com/trikker/auth/internal/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	user, err := h.authService.RegisterUser(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}
