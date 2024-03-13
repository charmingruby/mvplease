package domain

import (
	"github.com/charmingruby/mvplease/internal/core/contracts"
	"github.com/google/uuid"
)

type ServiceContract interface {
	Account(id uuid.UUID) (Account, error)
	Accounts() ([]Account, error)
	CreateAccount(a *Account) error
	Login(email, password string) (*Account, error)
}

type Service struct {
	accounts            AccountRepository
	cryptographyService contracts.CryptographyContract
}

func NewService(accounts AccountRepository, cryptographyService contracts.CryptographyContract) *Service {
	svc := &Service{
		accounts:            accounts,
		cryptographyService: cryptographyService,
	}

	return svc
}
