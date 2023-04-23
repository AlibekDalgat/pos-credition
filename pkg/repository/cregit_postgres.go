package repository

import (
	"fmt"
	posCreditation "github.com/AlibekDalgat/pos-credition"
	"github.com/jmoiron/sqlx"
)

type CreditPostgres struct {
	db *sqlx.DB
}

func NewCreditPostgres(db *sqlx.DB) *CreditPostgres {
	return &CreditPostgres{db}
}

func (creditPostgres *CreditPostgres) Create(cr posCreditation.NewCredit, mpId, agentId string) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, summary, timelimit, agent_id, m_place_id) values ($1, $2, $3, $4, $5) RETURNING id", creditsTable)
	row := creditPostgres.db.QueryRow(query, cr.Title, cr.Summary, cr.Timelimit, agentId, mpId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
