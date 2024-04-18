package service

import (
	"todo-list/internal/models"
	"todo-list/internal/repository"
)

type UsersService struct {
	repo *repository.Repository
}

func NewUsersService(repo *repository.Repository) *UsersService {
	return &UsersService{repo: repo}
}

func(u *UsersService) GetAll() ([]models.Person, error) {
	return u.repo.GetAll()
}