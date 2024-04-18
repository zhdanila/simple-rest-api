package service

import (
	"crypto/sha1"
	"fmt"
	"todo-list/internal/models"
	"todo-list/internal/repository"
)

const salt = "asoif30FJ#(F_#IJfolf)#(FPSldjfPO85fdsf"

type AuthService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func(s *AuthService) SignUp(person models.Person) int {
	person.Password = generatePasswordHash(person.Password)
	return s.repo.SignUp(person)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}