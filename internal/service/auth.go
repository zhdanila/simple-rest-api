package service

import (
	"todo-list/internal/models"
	"todo-list/internal/repository"
)

type AuthService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func(s *AuthService) SignUp(person models.Person) int {
	return s.repo.SignUp(person)
}