package repository

import (
	"github.com/jmoiron/sqlx"
	"todo-list/internal/models"
)

const usersTable = "users"

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}

type Authorization interface {
	SignUp(person models.Person) int
}