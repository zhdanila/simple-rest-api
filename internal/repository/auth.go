package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"todo-list/internal/models"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func(r *AuthRepository) SignUp(person models.Person) int {
	var id int

	query := fmt.Sprintf("INSERT INTO %s(name, password, username) VALUES($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, person.Name, person.Password, person.Username)
	row.Scan(&id)

	return id
}