package repository

import (
	posCreditation "github.com/AlibekDalgat/pos-credition"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	GetUser(username, password string) (posCreditation.TodoAgent, error)
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
	DeleteById(marketPlaceId string) error
}

type TodoAgent interface {
	Create(agent posCreditation.TodoAgent) (string, error)
	GetAll() ([]posCreditation.TodoAgent, error)
	GetById(agentId string) ([]posCreditation.InfoMPsAgent, error)
	UpdateById(agentId string, input posCreditation.UpdateAgentInput) error
	DeleteById(id string) error
	NewAccessToMP(marketPlace posCreditation.AccessingToMP, id string) (int, error)
}
type Credit interface {
	Create(cr posCreditation.NewCredit, mpId, agentId string) (int, error)
}

type Repository struct {
	Authorization
	TodoShop
	TodoMarketPlace
	TodoAgent
	Credit
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:   NewAuthPostgres(db),
		TodoShop:        NewTodoShopPostgres(db),
		TodoMarketPlace: NewTodoMarketPlacePostgres(db),
		TodoAgent:       NewTodoAgentPostgres(db),
		Credit:          NewCreditPostgres(db),
	}
}
