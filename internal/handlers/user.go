package handlers

import (
	"backend/internal/models"
	"backend/internal/services"
	"fmt"
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

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var payload models.UserCreatePayload 

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body: " + err.Error(),
		})
		return
	}

	createdUser, err := h.service.Create(&payload)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to create user: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "User created successfully",
		"results": createdUser,
	})
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "User ID is required",
		})
		return
	}

	fmt.Println(idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid User ID format",
		})
		return
	}

	var payload models.UserUpdatePayload 
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body: " + err.Error(),
		})
		return
	}

	updatedUser, err := h.service.Update(id, &payload)
	if err != nil {
		if err.Error() == "user not found" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "User not found",
			})
			return
		}
		if err.Error() == "user name cannot be empty" || err.Error() == "user email cannot be empty" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
		
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Failed to update user: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User updated successfully",
		"results": updatedUser,
	})
}