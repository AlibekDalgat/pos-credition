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
	GetById(userId, id int) (list posCreditation.TodoShop, err error)
	UpdateById(id string, input posCreditation.UpdateShopInput) error
	DeleteById(id string) error
}

type TodoItem interface {
	Create(userId, listId int, item posCreditation.TodoItem) (int, error)
	GetAll(userId, listId int) ([]posCreditation.TodoItem, error)
	GetById(userId, itemId int) (posCreditation.TodoItem, error)
	UpdateById(userId, itemId int, input posCreditation.UpdateMarketPlaceInput) error
	DeleteById(userId, id int) error
}

type Service struct {
	Authorization
	TodoShop
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoShop:      NewTodoShopService(repos.TodoShop),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoShop),
	}
}
