package service

import (
	"github.com/Hudayberdyyev/todo-app-Maksim-"
	"github.com/Hudayberdyyev/todo-app-Maksim-/pkg/repository"
)

type TodoItemService struct {
	repos     repository.TodoItem
	listRepos repository.TodoList
}

func NewTodoItemService(repos repository.TodoItem, listRepos repository.TodoList) *TodoItemService {
	return &TodoItemService{
		repos:     repos,
		listRepos: listRepos,
	}
}

func (s *TodoItemService) Create(userId, listId int, input todo.TodoItem) (int, error) {
	_, err := s.listRepos.GetById(userId, listId)
	if err != nil {
		// list does not exists or does not belongs to user
		return 0, err
	}
	return s.repos.Create(listId, input)
}
