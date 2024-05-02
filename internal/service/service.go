package service

import (
	"todo-list/internal/models"
	"todo-list/internal/repository"
)

type Service struct {
	Authorization
	TodoList
	TodoItem
}

type TodoList interface {
	Create(userId int, list models.TodoList) (int, error)
	GetAll(userId int) ([]models.TodoList, error)
	GetById(userId int, listId int) (models.TodoList, error)
	Delete(userId int, listId int) error
	Update(userId, listIntId int, input models.UpdateListInput) error
}

type Authorization interface {
	SignUp(person models.User) int
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoItem interface {
	Create(userId, listId int, item models.TodoItem) (int, error)
	GetAll(userId, listId int) ([]models.TodoItem, error)
	GetById(userId, itemId int) (models.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input models.UpdateItemInput) error
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewItemService(repos.TodoItem),
	}
}
