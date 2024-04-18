package service

import (
	"todo-list/internal/models"
	"todo-list/internal/repository"
)

type Service struct {
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
	}
}

type Authorization interface {
	SignUp(person models.Person) int
}
