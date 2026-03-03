package repository

import "backend/internal/models"

type UserRepository interface {
	GetAll() []models.User
	// GetByID(id string) (*models.User, error)
	// Create(user *models.User) error
}

type UserRepo struct {
	db *[]models.User
}

func NewUserRepository(db *[]models.User) *UserRepo {
	return &UserRepo{
		db: db,  
	}
}

func (r *UserRepo) GetAll() []models.User {
	if r.db == nil {
		return []models.User{}
	}
	return *r.db
}