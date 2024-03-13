package domain

import "github.com/google/uuid"

type AccountRepository interface {
	Account(id uuid.UUID) (Account, error)
	Accounts() ([]Account, error)
	CreateAccount(a *Account) error
	SaveAccount(a *Account) error
	DeleteAccount(id uuid.UUID) error
}
