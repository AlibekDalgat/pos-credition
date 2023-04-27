package posCreditation

import "errors"

type TodoShop struct {
	Id    string `json:"id" db:"id"`
	Title string `json:"title" db:"title" binding:"required"`
}

type TodoMarketPlace struct {
	Id     string `json:"id" db:"id"`
	Title  string `json:"title" db:"title" binding:"required"`
	ShopId string `json:"shop_id" db:"shop_id" binding:"required"`
}

type TodoAgent struct {
	Login    string `json:"login" db:"login"`
	Fio      string `json:"fio" db:"fio" binding:"required"`
	Password string `json:"password" db:"password_hash" binding:"required"`
}
type AccessingToMP struct {
	Id string `json:"id" db:"id"`
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

type UpdateAgentInput struct {
	Fio *string `json:"fio"`
}

func (input UpdateAgentInput) Validate() error {
	if input.Fio == nil {
		return errors.New("update strukturunuki mağnalar yoq")
	}
	return nil
}

type NewCredit struct {
	Id        int    `json:"id" db:"id"`
	Title     string `json:"title" db:"title" binding:"required"`
	Summary   string `json:"summary" db:"summary" binding:"required"`
	Timelimit string `json:"timelimit" db:"timelimit" binding:"required"`
}
