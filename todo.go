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
type TodoMarketPlace struct {
	Id     string `json:"id" db:"id"`
	Title  string `json:"title" db:"title" binding:"required"`
	ShopId string `json:"shop_id" db:"shop_id" binding:"required"`
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
