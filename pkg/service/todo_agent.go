package service

import (
	"github.com/AlibekDalgat/pos-credition"
	"github.com/AlibekDalgat/pos-credition/pkg/repository"
)

type TodoAgentService struct {
	repo repository.TodoAgent
}

func NewTodoAgentService(agentRepo repository.TodoAgent) *TodoAgentService {
	return &TodoAgentService{agentRepo}
}

func (agentService *TodoAgentService) Create(agent posCreditation.TodoAgent) (string, error) {
	agent.Password = generatePassword(agent.Password)
	return agentService.repo.Create(agent)
}

func (agentService *TodoAgentService) GetAll() ([]posCreditation.TodoAgent, error) {
	return agentService.repo.GetAll()
}

func (agentService *TodoAgentService) GetById(agentId string) ([]posCreditation.InfoMPsAgent, error) {
	return agentService.repo.GetById(agentId)
}

func (agentService *TodoAgentService) UpdateById(agentId string, input posCreditation.UpdateAgentInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return agentService.repo.UpdateById(agentId, input)
}

func (agentService *TodoAgentService) DeleteById(agentId string) error {
	return agentService.repo.DeleteById(agentId)
}

func (agentService *TodoAgentService) NewAccessToMP(marketPlace posCreditation.AccessingToMP, id string) (int, error) {
	return agentService.repo.NewAccessToMP(marketPlace, id)
}
