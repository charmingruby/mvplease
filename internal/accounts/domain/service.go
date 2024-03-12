package domain

import "github.com/google/uuid"

type ServiceInterface interface {
	Account(id uuid.UUID) (Account, error)
	Accounts() ([]Account, error)
	CreateAccount(a *Account) error
}

type Service struct {
	accounts AccountRepository
}

func NewService(accounts AccountRepository) *Service {
	svc := &Service{
		accounts: accounts,
	}

	return svc
}
