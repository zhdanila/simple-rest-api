package repository

import (
	"github.com/jmoiron/sqlx"
	"todo-list/internal/models"
)

const usersTable = "users"

type Repository struct {
	Authorization
	Users
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Users: NewUsersRepository(db),

	}
}

type Authorization interface {
	SignUp(person models.User) int
	GetUser(username, password string) (models.User, error)
}

type Users interface {

}