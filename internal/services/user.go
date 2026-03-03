package services

import (
	"backend/internal/models"
	"backend/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) GetAll() []models.User {
	return s.repo.GetAll()
}