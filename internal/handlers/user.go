package handlers

import (
	"backend/internal/services"
	"net/http"
	"strconv"

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

func (h *UserHandler) GetByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID is required",
		})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid ID format, must be an integer",
		})
		return
	}

	user, err := h.service.GetById(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to retrieve user: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Successfully retrieved user",
		"results": user,
	})
}