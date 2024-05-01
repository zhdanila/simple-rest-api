package service

import (
	"todo-list/internal/models"
	"todo-list/internal/repository"
)

type Service struct {
	Authorization
	Users
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
		Users:         NewUsersService(repo),
	}
}

type Authorization interface {
	SignUp(person models.Person) int
	GenerateToken(username, password string) (string, error)
}

type Users interface {
	GetAll() ([]models.Person, error)
}
