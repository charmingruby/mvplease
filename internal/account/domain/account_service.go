package domain

import "github.com/google/uuid"

func (s *Service) Account(id uuid.UUID) (Account, error) {
	account, err := s.accounts.FindAccountByID(id)
	if err != nil {
		return Account{}, err
	}

	return account, nil
}

func (s *Service) Accounts(page uint) ([]Account, error) {
	accounts, err := s.accounts.FetchAccounts(page)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (s *Service) CreateAccount(a *Account) error {
	if _, err := s.accounts.FindAccountByEmail(a.Email); err != nil {
		return err
	}

	hashedPassword, err := s.cryptographyService.GenerateHash(a.Password)
	if err != nil {
		return err
	}

	a.SetPassword(hashedPassword)

	if err := s.accounts.CreateAccount(a); err != nil {
		return err
	}

	return nil
}

func (s *Service) Login(email, password string) error {
	acc, err := s.accounts.FindAccountByEmail(email)
	if err != nil {
		return err
	}

	validCredentials := s.cryptographyService.VerifyHash(acc.Password, password)
	if !validCredentials {
		return err
	}

	return nil
}
