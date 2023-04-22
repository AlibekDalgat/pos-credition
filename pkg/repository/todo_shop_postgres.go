package repository

import (
	"fmt"
	"github.com/AlibekDalgat/pos-credition"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
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

func (shopPostgres *TodoShopPostgres) GetAll(userId int) ([]posCreditation.TodoShop, error) {
	var lists []posCreditation.TodoShop
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1", shopTable, usersListsTable)
	err := shopPostgres.db.Select(&lists, query, userId)
	return lists, err
}

func (shopPostgres *TodoShopPostgres) GetById(userId, id int) (posCreditation.TodoShop, error) {
	var list posCreditation.TodoShop
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2", shopTable, usersListsTable)
	err := shopPostgres.db.Get(&list, query, userId, id)
	return list, err
}

func (shopPostgres *TodoShopPostgres) DeleteById(userId, id int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id = $1 AND ul.list_id = $2", shopTable, usersListsTable)
	_, err := shopPostgres.db.Exec(query, userId, id)
	return err
}

func (shopPostgres *TodoShopPostgres) UpdateById(userId, id int, input posCreditation.UpdateShopInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		shopTable, setQuery, usersListsTable, argId, argId+1)
	args = append(args, id, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s	", args)
	_, err := shopPostgres.db.Exec(query, args...)
	return err
}
