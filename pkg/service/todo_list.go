package service

import (
	"github.com/AlibekDalgat/todo-app"
	"github.com/AlibekDalgat/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo}
}

func (listService *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return listService.repo.Create(userId, list)
}

func (listService *TodoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return listService.repo.GetAll(userId)
}

func (listService *TodoListService) GetById(userId, id int) (todo.TodoList, error) {
	return listService.repo.GetById(userId, id)
}

func (listService *TodoListService) DeleteById(userId, id int) error {
	return listService.repo.DeleteById(userId, id)
}

func (listService *TodoListService) UpdateById(userId, id int, input todo.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return listService.repo.UpdateById(userId, id, input)
}
