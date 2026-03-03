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

func (s *UserService) GetById(id int) (*models.User, error) {
	user, err := s.repo.GetByID(id)

	if err != nil {
		return nil, err
	}

   return user, nil
}