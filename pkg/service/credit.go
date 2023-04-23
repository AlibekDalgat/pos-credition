package service

import (
	posCreditation "github.com/AlibekDalgat/pos-credition"
	"github.com/AlibekDalgat/pos-credition/pkg/repository"
)

type CreditService struct {
	repo repository.Credit
}

func NewCreditService(creditRepo repository.Credit) *CreditService {
	return &CreditService{creditRepo}
}

func (creditService *CreditService) Create(cr posCreditation.NewCredit, mpId, agentId string) (int, error) {
	return creditService.repo.Create(cr, mpId, agentId)
}
