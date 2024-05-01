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

func(r *AuthRepository) SignUp(person models.User) int {
	var id int

	query := fmt.Sprintf("INSERT INTO %s(name, password_hash, username) VALUES($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, person.Name, person.Password, person.Username)
	row.Scan(&id)

	return id
}

func(r *AuthRepository) GetUser(username, password string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}