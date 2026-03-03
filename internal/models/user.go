package models

type User struct {
	ID int
	Name string
	Email string
	Password string
	Phone string
}

type UserCreatePayload struct {
    Name string
	Email string
	Password string
	Phone string
}

type UserUpdatePayload struct {
	Name string
	Password string
	Phone string
}