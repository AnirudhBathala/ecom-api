package models

import "time"

type UserStore interface{
	GetUserByEmail(email string) (*User,error)
	GetUserByID(id int) (*User,error)
	CreateUser(User) error
}


type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"LastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
