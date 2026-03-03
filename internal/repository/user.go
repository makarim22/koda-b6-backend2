package repository

import (
	     "backend/internal/models"
         "errors"
       )

type UserRepository interface {
	GetAll() []models.User
	GetByID(id int) (*models.User, error)
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

func (r *UserRepo) GetByID(id int) (*models.User, error) {
    if r.db == nil {
        return nil, errors.New("repository belum dibuat")
    }
    for _, user := range *r.db {
        if user.ID == id {
            return &user, nil
        }
    }
    return nil, errors.New("user tidak ditemukan")
}