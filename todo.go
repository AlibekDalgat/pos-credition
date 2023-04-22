package posCreditation

import "errors"

type TodoShop struct {
	Id    string `json:"id" db:"id"`
	Title string `json:"title" db:"title" binding:"required"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}
type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateShopInput struct {
	Title *string `json:"title"`
}

func (input UpdateShopInput) Validate() error {
	if input.Title == nil {
		return errors.New("update strukturunuki mağnalar yoq")
	}
	return nil
}

type UpdateMarketPlaceInput struct {
	Title *string `json:"title"`
}

func (input UpdateMarketPlaceInput) Validate() error {
	if input.Title == nil {
		return errors.New("update strukturunuki mağnalar yoq")
	}
	return nil
}
