package service

import (
	"github.com/AlibekDalgat/pos-credition"
	"github.com/AlibekDalgat/pos-credition/pkg/repository"
)

type Authorization interface {
	CreateUser(user posCreditation.User) (int, error)
	GenerateTokenForAgent(login, password string) (string, error)
	GenerateTokenForAdmin() (string, error)
	ParseToken(token string) (string, bool, error)
}

type TodoShop interface {
	Create(shop posCreditation.TodoShop) (string, error)
	GetAll() ([]posCreditation.TodoShop, error)
	GetById(id string) (list posCreditation.TodoShop, err error)
	UpdateById(id string, input posCreditation.UpdateShopInput) error
	DeleteById(id string) error
}

type TodoMarketPlace interface {
	Create(marketPlace posCreditation.TodoMarketPlace) (string, error)
	GetAll() ([]posCreditation.TodoMarketPlace, error)
	GetById(marketPlaceId string) (posCreditation.TodoMarketPlace, error)
	UpdateById(marketPlaceId string, input posCreditation.UpdateMarketPlaceInput) error
	DeleteById(id string) error
}

type Service struct {
	Authorization
	TodoShop
	TodoMarketPlace
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:   NewAuthService(repos.Authorization),
		TodoShop:        NewTodoShopService(repos.TodoShop),
		TodoMarketPlace: NewTodoMarketPlaceService(repos.TodoMarketPlace),
	}
}
