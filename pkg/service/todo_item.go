package service

import (
	"github.com/AlibekDalgat/pos-credition"
	"github.com/AlibekDalgat/pos-credition/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoShop
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoShop) *TodoItemService {
	return &TodoItemService{repo, listRepo}
}

func (itemService *TodoItemService) Create(userId, listId int, item posCreditation.TodoItem) (int, error) {
	_, err := itemService.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return itemService.repo.Create(listId, item)
}

func (itemService *TodoItemService) GetAll(userId, listId int) ([]posCreditation.TodoItem, error) {
	return itemService.repo.GetAll(userId, listId)
}

func (itemService *TodoItemService) GetById(userId, itemId int) (posCreditation.TodoItem, error) {
	return itemService.repo.GetById(userId, itemId)
}

func (itemService *TodoItemService) UpdateById(userId, itemId int, input posCreditation.UpdateMarketPlaceInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return itemService.repo.UpdateById(userId, itemId, input)
}

func (itemService *TodoItemService) DeleteById(userId, itemId int) error {
	return itemService.repo.DeleteById(userId, itemId)
}
