package repository

import (
	"github.com/jmoiron/sqlx"
	"todo-list/internal/models"
)

type TodoItemRepository struct {
	db *sqlx.DB
}

func NewTodoItemRepository(db *sqlx.DB) *TodoItemRepository {
	return &TodoItemRepository{db: db}
}

func(r *TodoItemRepository) Create(listId int, item models.TodoItem) (int, error) {
	return 0, nil
}

func(r *TodoItemRepository) GetAll(userId, listId int) ([]models.TodoItem, error) {
	return nil, nil
}

func(r *TodoItemRepository) GetById(userId, itemId int) (models.TodoItem, error) {
	return models.TodoItem{}, nil
}

func(r *TodoItemRepository) Delete(userId, itemId int) error {
	return nil;
}

func(r *TodoItemRepository) Update(userId, itemId int, input models.UpdateItemInput) error {
	return nil
}