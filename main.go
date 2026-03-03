package main

import (
	"backend/internal/di"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	container := di.NewContainer()

	userHandler := container.UserHandler()

	users := r.Group("/users")
	{
		users.GET("", userHandler.GetAll)
		users.GET("/:id", userHandler.GetByID)
		users.POST("", userHandler.CreateUser)
		users.PUT("/:id", userHandler.UpdateUser)

	}

	r.Run("localhost:8888")
}
