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
	GetById(userId, id int) (list posCreditation.TodoShop, err error)
	UpdateById(id string, input posCreditation.UpdateShopInput) error
	DeleteById(id string) error
}

type TodoItem interface {
	Create(listId int, item posCreditation.TodoItem) (int, error)
	GetAll(userId, listId int) ([]posCreditation.TodoItem, error)
	GetById(userId, itemId int) (posCreditation.TodoItem, error)
	UpdateById(userId, itemId int, input posCreditation.UpdateMarketPlaceInput) error
	DeleteById(userId, itemId int) error
}

type Repository struct {
	Authorization
	TodoShop
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoShop: NewTodoShopPostgres(db),
		TodoItem: NewTodoItemPostgres(db),
	}
}
