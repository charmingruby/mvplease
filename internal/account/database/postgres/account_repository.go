package postgres

import (
	"github.com/charmingruby/mvplease/internal/account/domain"
	"github.com/charmingruby/mvplease/pkg/errors"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func NewAccountRepository(db *sqlx.DB) (AccountRepository, error) {
	stmts := make(map[string]*sqlx.Stmt)

	for queryName, statement := range queriesAccount() {
		stmt, err := db.Preparex(statement)
		if err != nil {
			return AccountRepository{},
				errors.NewStatementError(err, errors.StatementPreparationErrorMessage(queryName))
		}

		stmts[queryName] = stmt
	}

	return AccountRepository{
		statements: stmts,
	}, nil
}

type AccountRepository struct {
	statements map[string]*sqlx.Stmt
}

func (r *AccountRepository) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.statements[queryName]

	if !ok {
		return nil, errors.NewStatementError(nil, errors.StatementNotPreparedErrorMessage(queryName))

	}

	return stmt, nil
}

func (r *AccountRepository) FindAccountByID(id uuid.UUID) (domain.Account, error) {
	stmt, err := r.statement(getAccountByID)
	if err != nil {
		return domain.Account{}, err
	}

	var account domain.Account
	if err := stmt.Get(&account, id); err != nil {
		return domain.Account{},
			errors.NewNotFoundError(err, "Account")
	}

	return account, nil
}

func (r *AccountRepository) FindAccountByEmail(email string) (domain.Account, error) {
	stmt, err := r.statement(getAccountByEmail)
	if err != nil {
		return domain.Account{}, err
	}

	var account domain.Account
	if err := stmt.Get(&account, email); err != nil {
		return domain.Account{},
			errors.NewNotFoundError(err, "Account")
	}

	return account, nil
}

func (r *AccountRepository) FetchAccounts(page uint) ([]domain.Account, error) {
	_, err := r.statement(fetchAccounts)
	if err != nil {
		return []domain.Account{}, err
	}

	return []domain.Account{}, nil
}
func (r *AccountRepository) CreateAccount(a *domain.Account) error {
	_, err := r.statement(createAccount)
	if err != nil {
		return err
	}

	return nil
}

func (r *AccountRepository) SaveAccount(a *domain.Account) error {
	_, err := r.statement(saveAccount)
	if err != nil {
		return err
	}

	return nil
}

func (r *AccountRepository) DeleteAccount(id uuid.UUID) error {
	_, err := r.statement(deleteAccount)
	if err != nil {
		return err
	}

	return nil
}
