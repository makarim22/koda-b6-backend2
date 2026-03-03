package services

import (
	"backend/internal/models"
	"backend/internal/repository"
	"fmt"
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

func (s *UserService) Create(payload *models.UserCreatePayload) (*models.User, error) {
    userToSave := &models.User{
        Name:  payload.Name,
        Email: payload.Email,
        Phone: payload.Phone,
    }

    createdUser, err := s.repo.Create(userToSave)
    if err != nil {
        return nil, fmt.Errorf("failed to create user in repository: %w", err)
    }

    return createdUser, nil
}