
package handlers

import (
	"net/http"
	"styleai/internal/models"
	"styleai/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct { authService *services.AuthService }

func NewAuthHandler(authService *services.AuthService) *AuthHandler { return &AuthHandler{authService: authService} }

func (h *AuthHandler) SignUp(c *gin.Context) { var user models.User if err := c.BindJSON(&user); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"}) return }

	createdUser, err := h.authService.SignUp(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered", "user_id": createdUser.ID})

}

func (h *AuthHandler) SignIn(c *gin.Context) { var user models.User if err := c.BindJSON(&user); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"}) return }

	foundUser, err := h.authService.SignIn(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Logged in", "user_id": foundUser.ID})

}
