package repository

import (
	"fmt"
	"github.com/AlibekDalgat/pos-credition"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db}
}

func (itemPostgres *TodoItemPostgres) Create(listId int, item posCreditation.TodoItem) (int, error) {
	tx, err := itemPostgres.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int

	createQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", todoItemsTable)
	row := tx.QueryRow(createQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (item_id, lists_id) values ($1, $2) RETURNING id", listsItemsTable)
	_, err = tx.Exec(createListItemsQuery, itemId, listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

func (itemPostgres *TodoItemPostgres) GetAll(userId, listId int) ([]posCreditation.TodoItem, error) {
	var items []posCreditation.TodoItem
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti "+
		"INNER JOIN %s li on li.item_id = ti.id INNER JOIN %s ul on ul.list_id = li.lists_id WHERE li.lists_id = $1 AND ul.user_id = $2",
		todoItemsTable, listsItemsTable, usersListsTable)
	if err := itemPostgres.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}
	return items, nil
}

func (itemPostgres *TodoItemPostgres) GetById(userId, itemId int) (posCreditation.TodoItem, error) {
	var item posCreditation.TodoItem
	query := fmt.Sprintf("SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti "+
		"INNER JOIN %s li on li.item_id = ti.id INNER JOIN %s ul on ul.list_id = li.lists_id WHERE ti.id = $1 AND ul.user_id = $2",
		todoItemsTable, listsItemsTable, usersListsTable)
	if err := itemPostgres.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}
	return item, nil
}

func (itemPostgres *TodoItemPostgres) UpdateById(userId, itemId int, input posCreditation.UpdateMarketPlaceInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s ti SET %s FROM %s li, %s ul WHERE ti.id = li.item_id AND li.lists_id = ul.list_id AND ti.id=$%d AND ul.user_id=$%d",
		todoItemsTable, setQuery, listsItemsTable, usersListsTable, argId, argId+1)
	args = append(args, itemId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s	", args)
	_, err := itemPostgres.db.Exec(query, args...)
	return err
}

func (itemPostgres *TodoItemPostgres) DeleteById(userId, itemId int) error {
	query := fmt.Sprintf("DELETE FROM %s ti USING %s li, %s ul "+
		"WHERE ti.id = li.item_id AND li.lists_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2",
		todoItemsTable, listsItemsTable, usersListsTable)
	_, err := itemPostgres.db.Exec(query, userId, itemId)
	return err
}
