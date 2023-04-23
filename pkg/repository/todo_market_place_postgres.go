package repository

import (
	"errors"
	"fmt"
	"github.com/AlibekDalgat/pos-credition"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoMarketPlacePostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db}
}

func (marketPlacePostgres *TodoItemPostgres) Create(marketPlace posCreditation.TodoMarketPlace) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (id, title, shop_id) values ($1, $2, $3) RETURNING id", marketPlacesTable)
	row := marketPlacePostgres.db.QueryRow(query, marketPlace.Id, marketPlace.Title, marketPlace.ShopId)
	if err := row.Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}

func (marketPlacePostgres *TodoItemPostgres) GetAll() ([]posCreditation.TodoMarketPlace, error) {
	var marketPlaces []posCreditation.TodoMarketPlace
	query := fmt.Sprintf("SELECT mp.id, mp.title, mp.shop_id FROM %s mp",
		marketPlacesTable)
	if err := marketPlacePostgres.db.Select(&marketPlaces, query); err != nil {
		return nil, err
	}
	return marketPlaces, nil
}

func (marketPlacePostgres *TodoItemPostgres) GetById(markePlaceId string) (posCreditation.TodoMarketPlace, error) {
	var item posCreditation.TodoMarketPlace
	query := fmt.Sprintf("SELECT mp.id, mp.title, mp.shop_id FROM %s mp WHERE mp.id = '%s'",
		marketPlacesTable, markePlaceId)
	if err := marketPlacePostgres.db.Get(&item, query); err != nil {
		return item, err
	}
	return item, nil
}

func (marketPlacePostgres *TodoItemPostgres) UpdateById(marketPlaceId string, input posCreditation.UpdateMarketPlaceInput) error {
	inputTitle := *input.Title
	query := fmt.Sprintf("UPDATE %s mp SET title='%s' WHERE id='%s'",
		marketPlacesTable, inputTitle, marketPlaceId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s	", inputTitle)
	res, err := marketPlacePostgres.db.Exec(query)
	rowsDeleted, err := res.RowsAffected()
	if rowsDeleted == 0 {
		err = errors.New("нет такой торговой точки")
	}
	return err
}

func (marketPlacePostgres *TodoItemPostgres) DeleteById(marketPlaceId string) error {
	query := fmt.Sprintf("DELETE FROM %s mp WHERE mp.id = '%s'",
		marketPlacesTable, marketPlaceId)
	res, err := marketPlacePostgres.db.Exec(query)
	rowsDeleted, err := res.RowsAffected()
	if rowsDeleted == 0 {
		err = errors.New("нет такойторговой точки")
	}
	return err
}
