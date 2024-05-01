package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"todo-list/internal/models"
	"todo-list/internal/repository"
)

const (
	salt       = "asoif30FJ#(F_#IJfolf)#(FPSldjfPO85fdsf"
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

func (a *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := a.repo.Authorization.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", fmt.Errorf("error getting user: %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(user.ID),
	})

	signedToken, err := token.SignedString([]byte(signingKey))

	return signedToken, err
}

func (a *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		subStr, ok := claims["sub"].(string)
		if !ok {
			return 0, errors.New("unable to convert 'sub' to string")
		}

		id, err := strconv.Atoi(subStr)
		if err != nil {
			return 0, errors.New("unable to convert 'sub' to int")
		}

		return id, nil
	} else {
		return 0, errors.New("invalid token")
	}
}
