package endpoints

import (
	"net/http"

	"github.com/charmingruby/mvplease/internal/account/domain"
	"github.com/charmingruby/mvplease/internal/common/infra/rest"
	"github.com/charmingruby/mvplease/internal/common/infra/token"
	"github.com/charmingruby/mvplease/pkg/errors"
	"github.com/sirupsen/logrus"
)

func NewProfileHandler(s domain.ServiceContract, logger *logrus.Logger, jwt *token.JWTService) http.HandlerFunc {
	return makeProfileEndpoint(s, logger, jwt)
}

type profileResponse struct {
	Account domain.Account `json:"account"`
}

func makeProfileEndpoint(s domain.ServiceContract, logger *logrus.Logger, jwt *token.JWTService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := rest.RetrieveTokenFromRequest(r)

		payload, err := jwt.RetriveTokenPayload(token)
		if err != nil {
			payloadErr := errors.NewPayloadError(err, errors.PayloadErrorMessage())

			logger.Error(payloadErr)
			rest.NewResponse[any](w, payloadErr.Error(), nil, http.StatusInternalServerError)
			return
		}

		account, err := s.Account(payload.AccountID)
		if err != nil {
			notFoundErr, ok := err.(*errors.NotFoundError)
			if ok {
				logger.Error(notFoundErr.Error())
				rest.NewResponse[any](w, notFoundErr.Error(), nil, http.StatusNotFound)
				return
			}

			logger.Error(err.Error())
			rest.NewResponse[any](w, err.Error(), nil, http.StatusInternalServerError)
			return
		}

		body := profileResponse{
			Account: account,
		}

		logger.Infof("'%s' profile found", payload.AccountID)
		rest.NewResponse[profileResponse](w, "Profile found.", &body, http.StatusOK)
	}
}
