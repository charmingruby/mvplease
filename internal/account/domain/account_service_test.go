package domain

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type AccountMockRepository struct {
	mock.Mock
}

func (m *AccountMockRepository) FindAccountByID(id uuid.UUID) (Account, error) {
	return Account{}, nil
}

func (m *AccountMockRepository) FindAccountByEmail(email string) (Account, error) {
	return Account{}, nil
}

func (m *AccountMockRepository) FetchAccounts(page uint) ([]Account, error) {
	return []Account{}, nil
}

func (m *AccountMockRepository) CreateAccount(a *Account) error {
	return nil
}

func (m *AccountMockRepository) SaveAccount(a *Account) error {
	return nil
}

func (m *AccountMockRepository) DeleteAccount(a *Account) error {
	return nil
}

func TestCreateAccount(t *testing.T) {

}
