package account

import (
	"github.com/charmingruby/mvplease/internal/account/database/postgres"
	"github.com/charmingruby/mvplease/internal/account/domain"
	"github.com/charmingruby/mvplease/internal/shared/cryptography"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func NewService(db *sqlx.DB, logger *logrus.Logger) (*domain.Service, error) {
	account, err := postgres.NewAccountRepository(db)
	if err != nil {
		return nil, err
	}

	cryptographySvc := cryptography.NewCryptographyService()

	svc := domain.NewService(&account, cryptographySvc)

	return svc, nil
}

func NewHTTPHandler() {}
