package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"todo-list/internal/models"
)

type UsersRepository struct {
	db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

func(u *UsersRepository) GetAll() ([]models.Person, error) {
	var people []models.Person

	query := fmt.Sprintf("SELECT * FROM %s", usersTable)
	err := u.db.Select(&people, query)
	if err != nil {
		return nil, err
	}

	fmt.Println(people)

	return people, nil
}