package services

import (
	"backend/internal/models"
	"backend/internal/repository"
	"fmt"
	"errors"
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

func(s* UserService) Update(id int, payload *models.UserUpdatePayload) (*models.User, error){
	if payload.Name == "" {
        return nil, errors.New("user name cannot be empty")
    }
    if payload.Password == "" {
        return nil, errors.New("user email cannot be empty")
    }

	 userToUpdate := &models.User{
        Name:  payload.Name,
        Email: payload.Password,
        Phone: payload.Phone,
    }

    updatedUser, err := s.repo.Update(id, userToUpdate)
    if err != nil {
        return nil, fmt.Errorf("failed to update user: %w", err)
    }

    return updatedUser, nil
}