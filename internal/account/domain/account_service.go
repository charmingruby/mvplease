package domain

import (
	"github.com/charmingruby/mvplease/pkg/errors"
	"github.com/google/uuid"
)

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
	if _, err := s.accounts.FindAccountByEmail(a.Email); err == nil {
		conflictErr := errors.NewConflictError(err, "Account", "email")
		return conflictErr
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

func (s *Service) Login(email, password string) (*Account, error) {
	acc, err := s.accounts.FindAccountByEmail(email)
	if err != nil {
		invalidCredentialsErr := errors.NewInvalidCredentialsError()
		return nil, invalidCredentialsErr
	}

	validCredentials := s.cryptographyService.VerifyHash(acc.Password, password)
	if !validCredentials {
		invalidCredentialsErr := errors.NewInvalidCredentialsError()
		return nil, invalidCredentialsErr
	}

	return &acc, nil
}

func (s *Service) DeleteAccount(accountID, managerID uuid.UUID) error {
	a, err := s.accounts.FindAccountByID(accountID)
	if err != nil {
		return err
	}

	a.DeleteAccount(managerID)

	if err := s.accounts.DeleteAccount(&a); err != nil {
		return err
	}

	return nil
}

func (s *Service) UploadAvatar(accountID uuid.UUID, avatarURL string) error {
	account, err := s.accounts.FindAccountByID(accountID)
	if err != nil {
		return err
	}

	account.SetAvatarURL(avatarURL)

	if err := s.accounts.SaveAccount(&account); err != nil {
		return err
	}

	return nil
}
