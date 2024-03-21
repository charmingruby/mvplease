package domain

import (
	"github.com/charmingruby/mvplease/internal/common/core/contract"
	"github.com/google/uuid"
)

type ServiceContract interface {
	Account(id uuid.UUID) (Account, error)
	Accounts(page uint) ([]Account, error)
	CreateAccount(a *Account) error
	Login(email, password string) (*Account, error)
	DeleteAccount(accountID, managerID uuid.UUID) error
}

type Service struct {
	accounts            AccountRepository
	cryptographyService contract.CryptographyContract
}

func NewService(accounts AccountRepository, cryptographyService contract.CryptographyContract) *Service {
	svc := &Service{
		accounts:            accounts,
		cryptographyService: cryptographyService,
	}

	return svc
}
