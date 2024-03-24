package endpoints

import (
	"net/http"

	"github.com/charmingruby/mvplease/internal/account/domain"
	"github.com/charmingruby/mvplease/internal/common/core"
	"github.com/charmingruby/mvplease/internal/common/infra/rest"
	"github.com/charmingruby/mvplease/pkg/errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func NewDeleteAccountHandler(s domain.ServiceContract, logger *logrus.Logger) http.HandlerFunc {
	return makeDeleteAccountEndpoint(s, logger)
}

func makeDeleteAccountEndpoint(s domain.ServiceContract, logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		accountID := params["id"]

		if accountID == "" || !core.IsUUID(accountID) {
			logger.Errorf("Invalid account ID")
			rest.NewResponse[any](w, "Payload error: invalid id", nil, http.StatusBadRequest)
			return
		}

		payload, err := rest.RetrievePayloadFromRequest(r)
		if err != nil {
			logger.Error(err.Error())
			rest.NewResponse[any](w, err.Error(), nil, http.StatusInternalServerError)
			return
		}

		accountToDeleteUUID := uuid.MustParse(accountID)

		if err := s.DeleteAccount(accountToDeleteUUID, payload.AccountID); err != nil {
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

		logger.Infof("Account: '%s' deleted successfully", accountToDeleteUUID)
		rest.NewResponse[any](
			w,
			"Account deleted successfully.",
			nil,
			http.StatusOK,
		)
	}
}
