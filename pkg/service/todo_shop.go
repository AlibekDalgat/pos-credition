package service

import (
	"github.com/AlibekDalgat/pos-credition"
	"github.com/AlibekDalgat/pos-credition/pkg/repository"
)

type TodoShopService struct {
	repo repository.TodoShop
}

func NewTodoShopService(repo repository.TodoShop) *TodoShopService {
	return &TodoShopService{repo}
}

func (shopService *TodoShopService) Create(list posCreditation.TodoShop) (string, error) {
	return shopService.repo.Create(list)
}

func (shopService *TodoShopService) GetAll() ([]posCreditation.TodoShop, error) {
	return shopService.repo.GetAll()
}

func (shopService *TodoShopService) GetById(userId, id int) (posCreditation.TodoShop, error) {
	return shopService.repo.GetById(userId, id)
}

func (shopService *TodoShopService) DeleteById(id string) error {
	return shopService.repo.DeleteById(id)
}

func (shopService *TodoShopService) UpdateById(id string, input posCreditation.UpdateShopInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return shopService.repo.UpdateById(id, input)
}
