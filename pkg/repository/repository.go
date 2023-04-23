package repository

import (
	posCreditation "github.com/AlibekDalgat/pos-credition"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user posCreditation.User) (int, error)
	GetUser(username, password string) (posCreditation.User, error)
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

type Repository struct {
	Authorization
	TodoShop
	TodoMarketPlace
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoShop:        NewTodoShopPostgres(db),
		TodoMarketPlace: NewTodoMarketPlacePostgres(db),
	}
}
