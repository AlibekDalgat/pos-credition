package service

import (
	"errors"
	"github.com/AlibekDalgat/pos-credition"
	"github.com/AlibekDalgat/pos-credition/pkg/repository"
)

type TodoShopService struct {
	repo repository.TodoShop
}

func NewTodoShopService(repo repository.TodoShop) *TodoShopService {
	return &TodoShopService{repo}
}

func (shopService *TodoShopService) Create(shop posCreditation.TodoShop) (string, error) {
	if shop.Id == "" {
		return "", errors.New("не введён id")
	}
	return shopService.repo.Create(shop)
}

func (shopService *TodoShopService) GetAll() ([]posCreditation.TodoShop, error) {
	return shopService.repo.GetAll()
}

func (shopService *TodoShopService) GetById(id string) (posCreditation.TodoShop, error) {
	return shopService.repo.GetById(id)
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
