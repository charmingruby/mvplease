package domain

import "github.com/google/uuid"

type CryptographyInterface interface {
	GenerateHash(value string) (string, error)
	VerifyHash(hash, value string) bool
}

type ServiceInterface interface {
	Account(id uuid.UUID) (Account, error)
	Accounts() ([]Account, error)
	CreateAccount(a *Account) error
	Login(email, password string) (*Account, error)
}

type Service struct {
	accounts            AccountRepository
	cryptographyService CryptographyInterface
}

func NewService(accounts AccountRepository, cryptographyService CryptographyInterface) *Service {
	svc := &Service{
		accounts:            accounts,
		cryptographyService: cryptographyService,
	}

	return svc
}
