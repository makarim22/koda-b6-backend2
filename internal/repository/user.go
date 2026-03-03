package repository

import (
	"backend/internal/models"
	"errors"
	"fmt"
)

type UserRepository interface {
	GetAll() []models.User
	GetByID(id int) (*models.User, error)
	Create(user *models.User) (*models.User, error) 
	Update(id int, user *models.User) (*models.User, error)
}

type UserRepo struct {
	db *[]models.User
	nextID int
}

// func NewUserRepository(db *[]models.User) *UserRepo {
// 	return &UserRepo{
// 		db: db,  
// 	}
// }

func NewUserRepo(initialUsers []models.User) *UserRepo {
    db := initialUsers
    nextID := 1
    if len(db) > 0 {
        for _, u := range db {
            if u.ID >= nextID {
                nextID = u.ID + 1
            }
        }
    }
    return &UserRepo{db: &db, nextID: nextID}
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

func (r *UserRepo) Create(userToSave *models.User) (*models.User, error) {
    if r.db == nil {
        return nil, errors.New("repository not initialized")
    }

    userToSave.ID = r.nextID 
    r.nextID++

    *r.db = append(*r.db, *userToSave) 

    return userToSave, nil 
}

// func (r* UserRepo) Update(id int, updatedUser *models.User) (*models.User, error){
// 	  if r.db == nil {
//         return nil, errors.New("repository not initialized")
//     }
	
// 	fmt.Println(updatedUser)
// 	for i, user := range *r.db{
// 		if user.ID == id {
// 			user.Name = updatedUser.Name
// 			user.Password = updatedUser.Password
// 			user.Phone = updatedUser.Phone

//             (*r.db)[i]= user
// 			return &user, nil

// 		}
// 	}
// 	return nil, errors.New("user tidak ditemukan")
// }

func (r* UserRepo) Update(id int, updatedUser *models.User) (*models.User, error){
	  if r.db == nil {
        return nil, errors.New("repository not initialized")
    }
	
	fmt.Println(updatedUser)
	for i, user := range *r.db { 
		if user.ID == id {
			(*r.db)[i].Name = updatedUser.Name
			(*r.db)[i].Password = updatedUser.Password
			(*r.db)[i].Phone = updatedUser.Phone

			return &(*r.db)[i], nil 
		}
	}
	return nil, errors.New("user tidak ditemukan")
}