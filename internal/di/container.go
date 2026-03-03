package di

import (
	"backend/internal/handlers"
	"backend/internal/models"
	"backend/internal/repository"
	"backend/internal/services"
)

type Container struct {
	user        *[]models.User
	userRepo    repository.UserRepository
	userService *services.UserService
	userHandler *handlers.UserHandler
}

func NewContainer() *Container {
	

	var DataUser = []models.User{ 
		{ID: 1, Name: "Alice Smith", Email: "alice@example.com"},
		{ID: 2, Name: "Bob Johnson", Email: "bob@example.com"},
		{ID: 3, Name: "Charlie Brown", Email: "charlie@example.com"},
	}

	container := Container{
		user: &DataUser,
	}

	container.initDependencies()

	return &container
}

func (c *Container) initDependencies() {
	c.userRepo = repository.NewUserRepo(*c.user)
	c.userService = services.NewUserService(c.userRepo)
	c.userHandler = handlers.NewUserHandler(c.userService)
}

func (c *Container) UserHandler() *handlers.UserHandler {
	return c.userHandler
}