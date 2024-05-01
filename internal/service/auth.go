package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"todo-list/internal/models"
	"todo-list/internal/repository"
)

const (
	salt = "asoif30FJ#(F_#IJfolf)#(FPSldjfPO85fdsf"
	signingKey = "f39iOJF(#UFpo30i_#F{#09ifi)(#"
)

type AuthService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignUp(person models.User) int {
	person.Password = generatePasswordHash(person.Password)
	return s.repo.SignUp(person)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func(a *AuthService) GenerateToken(username, password string) (string, error) {
	//todo: get user, validate identity, create token

	user, err := a.repo.Authorization.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", fmt.Errorf("error getting user: %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(user.ID),
	})

	signedToken, err := token.SignedString([]byte(signingKey))

	return signedToken, err;
}