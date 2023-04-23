package service

import (
	"github.com/AlibekDalgat/pos-credition"
	"github.com/AlibekDalgat/pos-credition/pkg/repository"
)

type Authorization interface {
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

type TodoAgent interface {
	Create(agent posCreditation.TodoAgent) (string, error)
	GetAll() ([]posCreditation.TodoAgent, error)
	GetById(agentId string) (posCreditation.TodoAgent, error)
	UpdateById(agentId string, input posCreditation.UpdateAgentInput) error
	DeleteById(id string) error
	NewAccessToMP(marketPlace posCreditation.AccessingToMP, id string) (int, error)
}

type Credit interface {
	Create(cr posCreditation.NewCredit, mpId, agentId string) (int, error)
}

type Service struct {
	Authorization
	TodoShop
	TodoMarketPlace
	TodoAgent
	Credit
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:   NewAuthService(repos.Authorization),
		TodoShop:        NewTodoShopService(repos.TodoShop),
		TodoMarketPlace: NewTodoMarketPlaceService(repos.TodoMarketPlace),
		TodoAgent:       NewTodoAgentService(repos.TodoAgent),
		Credit:          NewCreditService(repos.Credit),
	}
}
