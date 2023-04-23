package service

import (
	"github.com/AlibekDalgat/pos-credition"
	"github.com/AlibekDalgat/pos-credition/pkg/repository"
)

type TodoMarketPlaceService struct {
	repo repository.TodoMarketPlace
}

func NewTodoMarketPlaceService(listRepo repository.TodoMarketPlace) *TodoMarketPlaceService {
	return &TodoMarketPlaceService{listRepo}
}

func (itemService *TodoMarketPlaceService) Create(marketPlace posCreditation.TodoMarketPlace) (string, error) {
	return itemService.repo.Create(marketPlace)
}

func (itemService *TodoMarketPlaceService) GetAll() ([]posCreditation.TodoMarketPlace, error) {
	return itemService.repo.GetAll()
}

func (itemService *TodoMarketPlaceService) GetById(marketPlaceId string) (posCreditation.TodoMarketPlace, error) {
	return itemService.repo.GetById(marketPlaceId)
}

func (itemService *TodoMarketPlaceService) UpdateById(marketPlaceId string, input posCreditation.UpdateMarketPlaceInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return itemService.repo.UpdateById(marketPlaceId, input)
}

func (itemService *TodoMarketPlaceService) DeleteById(marketPlaceId string) error {
	return itemService.repo.DeleteById(marketPlaceId)
}
