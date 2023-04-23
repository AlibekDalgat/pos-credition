package repository

import (
	"fmt"
	posCreditation "github.com/AlibekDalgat/pos-credition"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db}
}

func (p *AuthPostgres) GetUser(username, password string) (posCreditation.TodoAgent, error) {
	var agent posCreditation.TodoAgent
	query := fmt.Sprintf("SELECT ag.login FROM %s ag WHERE login=$1 AND password_hash=$2", agentsTable)
	err := p.db.Get(&agent, query, username, password)
	return agent, err
}
