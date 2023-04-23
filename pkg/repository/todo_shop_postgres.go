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
	query := fmt.Sprintf("SELECT shps.id, shps.title FROM %s shps WHERE id='%s' ", shopTable, id)
	err := shopPostgres.db.Get(&list, query)
	return list, err
}

func (shopPostgres *TodoShopPostgres) DeleteById(id string) error {
	query := fmt.Sprintf("DELETE FROM %s shps WHERE shps.id = '%s'",
		shopTable, id)
	res, err := shopPostgres.db.Exec(query)
	rowsDeleted, err := res.RowsAffected()
	if rowsDeleted == 0 {
		err = errors.New("нет такойторговой точки")
	}
	return err
}

func (shopPostgres *TodoShopPostgres) UpdateById(id string, input posCreditation.UpdateShopInput) error {
	inputTitle := *input.Title
	query := fmt.Sprintf("UPDATE %s shps SET title='%s' WHERE id='%s'",
		shopTable, inputTitle, id)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s	", input.Title)
	_, err := shopPostgres.db.Exec(query)
	return err
}
