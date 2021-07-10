package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/Hudayberdyyev/todo-app-Maksim-"
	"github.com/Hudayberdyyev/todo-app-Maksim-/pkg/repository"
)

const salt = "sdfasfsaf12313"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprint("%s", hash.Sum([]byte(salt)))
}
