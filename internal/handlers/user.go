package handlers

import (
	"backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(svc *services.UserService) *UserHandler {
	return &UserHandler{
		service: svc,
	}
}

func (h *UserHandler) GetAll(ctx *gin.Context) {
	users := h.service.GetAll()

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Successfully retrieved all users",
		"results": users,
	})
}