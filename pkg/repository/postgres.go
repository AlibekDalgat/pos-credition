package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	shopTable                = "shops"
	marketPlacesTable        = "market_places"
	agentsTable              = "agents"
	agentsMarketsPlacesTable = "agents_market_places"
	creditsTable             = "credits"
)

type Config struct {
	Host     string
	Post     string
	Username string
	Password string
	DBname   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Post, cfg.Username, cfg.DBname, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
