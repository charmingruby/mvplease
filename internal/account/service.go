package account

import (
	"github.com/charmingruby/mvplease/config"
	"github.com/charmingruby/mvplease/internal/account/database/postgres"
	"github.com/charmingruby/mvplease/internal/account/domain"
	http "github.com/charmingruby/mvplease/internal/account/transport/rest"
	"github.com/charmingruby/mvplease/internal/services/cryptography"
	"github.com/charmingruby/mvplease/internal/services/token"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func NewService(db *sqlx.DB, logger *logrus.Logger) (*domain.Service, error) {
	account, err := postgres.NewAccountRepository(db, logger)
	if err != nil {
		return nil, err
	}

	cryptographySvc := cryptography.NewCryptographyService()

	svc := domain.NewService(&account, cryptographySvc)

	return svc, nil
}

func NewHTTPService(router *mux.Router, service domain.ServiceContract, cfg *config.Config) error {
	issuer := "mvplease"
	jwtService := token.NewJWTService(cfg.JWTSecretKey, issuer)

	http.NewHTTPHandler(router, jwtService, service, cfg.Logger)

	return nil
}
