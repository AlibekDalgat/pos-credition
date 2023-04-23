package repository

import (
	"errors"
	"fmt"
	"github.com/AlibekDalgat/pos-credition"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type TodoAgentPostgres struct {
	db *sqlx.DB
}

func NewTodoAgentPostgres(db *sqlx.DB) *TodoAgentPostgres {
	return &TodoAgentPostgres{db}
}

func (agentPostgres *TodoAgentPostgres) Create(agent posCreditation.TodoAgent) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (fio, login, password_hash) values ($1, $2, $3) RETURNING login", agentsTable)
	row := agentPostgres.db.QueryRow(query, agent.Fio, agent.Login, agent.Password)
	if err := row.Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}

func (agentPostgres *TodoAgentPostgres) GetAll() ([]posCreditation.TodoAgent, error) {
	var agents []posCreditation.TodoAgent
	query := fmt.Sprintf("SELECT ag.fio, ag.login FROM %s ag",
		agentsTable)
	if err := agentPostgres.db.Select(&agents, query); err != nil {
		return nil, err
	}
	return agents, nil
}

func (agentPostgres *TodoAgentPostgres) GetById(agentId string) (posCreditation.TodoAgent, error) {
	var item posCreditation.TodoAgent
	query := fmt.Sprintf("SELECT ag.fio, ag.login FROM %s ag WHERE ag.login = '%s'",
		agentsTable, agentId)
	if err := agentPostgres.db.Get(&item, query); err != nil {
		return item, err
	}
	return item, nil
}

func (agentPostgres *TodoAgentPostgres) UpdateById(agentId string, input posCreditation.UpdateAgentInput) error {
	inputFio := *input.Fio
	query := fmt.Sprintf("UPDATE %s ag SET fio='%s' WHERE login='%s'",
		agentsTable, inputFio, agentId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s	", inputFio)
	res, err := agentPostgres.db.Exec(query)
	rowsDeleted, err := res.RowsAffected()
	if rowsDeleted == 0 {
		err = errors.New("нет такой торговой точки")
	}
	return err
}

func (agentPostgres *TodoAgentPostgres) DeleteById(agentId string) error {
	query := fmt.Sprintf("DELETE FROM %s ag WHERE ag.login = '%s'",
		agentsTable, agentId)
	res, err := agentPostgres.db.Exec(query)
	rowsDeleted, err := res.RowsAffected()
	if rowsDeleted == 0 {
		err = errors.New("нет такойторговой агента")
	}
	return err
}

func (agentPostgres *TodoAgentPostgres) NewAccessToMP(marketPlace posCreditation.AccessingToMP, agentId string) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (agent_id, m_place_id) values ($1, $2) RETURNING id", agentsMarketsPlacesTable)
	row := agentPostgres.db.QueryRow(query, agentId, marketPlace.Id)
	if err := row.Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}