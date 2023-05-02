package repository

import (
	"errors"
	"fmt"
	"github.com/AlibekDalgat/pos-credition"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type TodoShopPostgres struct {
	db *sqlx.DB
}

func NewTodoShopPostgres(db *sqlx.DB) *TodoShopPostgres {
	return &TodoShopPostgres{db}
}

func (shopPostgres *TodoShopPostgres) Create(shop posCreditation.TodoShop) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (id, title) values ($1, $2) RETURNING id", shopTable)
	row := shopPostgres.db.QueryRow(query, shop.Id, shop.Title)
	if err := row.Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}

func (shopPostgres *TodoShopPostgres) GetAll() ([]posCreditation.TodoShop, error) {
	var lists []posCreditation.TodoShop
	query := fmt.Sprintf("SELECT shps.id, shps.title FROM %s shps", shopTable)
	err := shopPostgres.db.Select(&lists, query)
	return lists, err
}

func (shopPostgres *TodoShopPostgres) GetById(id string) (posCreditation.TodoShop, error) {
	var list posCreditation.TodoShop
	query := fmt.Sprintf("SELECT shps.id, shps.title FROM %s shps WHERE id= $1 ", shopTable)
	err := shopPostgres.db.Get(&list, query, id)
	return list, err
}

func (shopPostgres *TodoShopPostgres) DeleteById(id string) error {
	query := fmt.Sprintf("DELETE FROM %s shps WHERE shps.id = $1",
		shopTable)
	res, err := shopPostgres.db.Exec(query, id)
	rowsDeleted, err := res.RowsAffected()
	if rowsDeleted == 0 {
		err = errors.New("нет такой магазина")
	}
	return err
}

func (shopPostgres *TodoShopPostgres) UpdateById(id string, input posCreditation.UpdateShopInput) error {
	inputTitle := *input.Title
	query := fmt.Sprintf("UPDATE %s shps SET title= $1 WHERE id= $2",
		shopTable)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s	", input.Title)
	res, err := shopPostgres.db.Exec(query, inputTitle, id)
	rowsUpdated, err := res.RowsAffected()
	if rowsUpdated == 0 {
		err = errors.New("нет такой магазина")
	}
	return err
}
