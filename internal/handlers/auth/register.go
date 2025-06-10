package auth

import (
	"net/http"
	authModel "ticket-backend/internal/models/auth"
	authService "ticket-backend/internal/services/auth"

	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	registerService *authService.RegisterService
}

func NewRegisterHandler() *RegisterHandler {
	return &RegisterHandler{
		registerService: authService.NewRegisterService(),
	}
}

func (h *RegisterHandler) Register(c *gin.Context) {
	var req authModel.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.registerService.Register(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "registration successful"})
}
