package domain

import "github.com/google/uuid"

type AccountRepository interface {
	FindAccountByID(id uuid.UUID) (Account, error)
	FindAccountByEmail(email string) (Account, error)
	FetchAccounts(page uint) ([]Account, error)
	CreateAccount(a *Account) error
	SaveAccount(a *Account) error
	DeleteAccount(a *Account) error
}
