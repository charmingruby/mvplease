package domain

import "github.com/google/uuid"

func (s *Service) Account(id uuid.UUID) (Account, error) {
	a := Account{}

	return a, nil
}

func (s *Service) Accounts() ([]Account, error) {
	return []Account{}, nil
}

func (s *Service) CreateAccount(a *Account) error {
	return nil
}
