package repository

import (
	"errors"
	"fmt"
	"github.com/AlibekDalgat/pos-credition"
	"github.com/jmoiron/sqlx"
	"strings"
)

type TodoMarketPlacePostgres struct {
	db *sqlx.DB
}

func NewTodoMarketPlacePostgres(db *sqlx.DB) *TodoMarketPlacePostgres {
	return &TodoMarketPlacePostgres{db}
}

func (marketPlacePostgres *TodoMarketPlacePostgres) Create(marketPlace posCreditation.TodoMarketPlace) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (id, title, shop_id) values ($1, $2, $3) RETURNING id", marketPlacesTable)
	row := marketPlacePostgres.db.QueryRow(query, marketPlace.Id, marketPlace.Title, marketPlace.ShopId)
	if err := row.Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}

func (marketPlacePostgres *TodoMarketPlacePostgres) GetAll() ([]posCreditation.TodoMarketPlace, error) {
	var marketPlaces []posCreditation.TodoMarketPlace
	query := fmt.Sprintf("SELECT mp.id, mp.title, mp.shop_id FROM %s mp",
		marketPlacesTable)
	if err := marketPlacePostgres.db.Select(&marketPlaces, query); err != nil {
		return nil, err
	}
	return marketPlaces, nil
}

func (marketPlacePostgres *TodoMarketPlacePostgres) GetById(markePlaceId string) (posCreditation.TodoMarketPlace, error) {
	var marketPlace posCreditation.TodoMarketPlace
	query := fmt.Sprintf("SELECT mp.id, mp.title, mp.shop_id FROM %s mp WHERE mp.id = $1",
		marketPlacesTable)
	if err := marketPlacePostgres.db.Get(&marketPlace, query, markePlaceId); err != nil {
		return marketPlace, err
	}
	return marketPlace, nil
}

func (marketPlacePostgres *TodoMarketPlacePostgres) UpdateById(marketPlaceId string, input posCreditation.UpdateMarketPlaceInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil && *input.Title != "" {
		setValues = append(setValues, fmt.Sprintf("title= $%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.ShopId != nil {
		setValues = append(setValues, fmt.Sprintf("shop_id= $%d", argId))
		args = append(args, *input.ShopId)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s mp SET %s WHERE mp.id= $%d",
		marketPlacesTable, setQuery, argId)
	args = append(args, marketPlaceId)
	_, err := marketPlacePostgres.db.Exec(query, args...)
	fmt.Println(err)
	return err
}

func (marketPlacePostgres *TodoMarketPlacePostgres) DeleteById(marketPlaceId string) error {
	query := fmt.Sprintf("DELETE FROM %s mp WHERE mp.id = $1",
		marketPlacesTable)
	res, err := marketPlacePostgres.db.Exec(query, marketPlaceId)
	rowsDeleted, err := res.RowsAffected()
	if rowsDeleted == 0 {
		err = errors.New("нет такойторговой точки")
	}
	return err
}
