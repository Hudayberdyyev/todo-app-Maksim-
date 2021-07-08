package service

import "github.com/Hudayberdyyev/todo-app-Maksim-/pkg/repository"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
