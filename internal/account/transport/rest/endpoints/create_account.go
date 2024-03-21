package endpoints

import (
	"net/http"

	"github.com/charmingruby/mvplease/internal/account/domain"
	"github.com/charmingruby/mvplease/internal/common/infra/rest"
	"github.com/charmingruby/mvplease/pkg/errors"
	"github.com/sirupsen/logrus"
)

func NewCreateAccountHandler(s domain.ServiceContract, logger *logrus.Logger) http.HandlerFunc {
	return makeCreateAccountEndpoint(s, logger)
}

type createAccountRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func makeCreateAccountEndpoint(s domain.ServiceContract, logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := createAccountRequest{}
		if err := rest.ParseRequest[createAccountRequest](&req, r.Body); err != nil {
			logger.Errorf("Error parsing create account request")
			rest.NewResponse[error](w, "Payload error", &err, http.StatusBadRequest)
			return
		}

		if err := rest.IsRequestValid(req); err != nil {
			logger.Error(err)
			rest.NewResponse[rest.ValidationErrors](w, "Payload error", err, http.StatusBadRequest)
			return
		}

		newAccount := domain.NewAccount(req.Name, req.Email, req.Password)
		if err := s.CreateAccount(newAccount); err != nil {
			conflictError, ok := err.(*errors.ConflictError)
			if ok {
				logger.Error(conflictError.Error())
				rest.NewResponse[any](w, conflictError.Error(), nil, http.StatusConflict)
				return
			}

			logger.Error(err.Error())
			rest.NewResponse[any](w, err.Error(), nil, http.StatusBadRequest)
			return
		}

		logger.Infof("Account: '%s' registered successfully", newAccount.ID)
		rest.NewResponse[any](
			w,
			"Account registered successfully.",
			nil,
			http.StatusCreated,
		)
	}
}
