package endpoints

import (
	"net/http"

	"github.com/charmingruby/mvplease/internal/account/domain"
	"github.com/charmingruby/mvplease/internal/services/token"
	"github.com/charmingruby/mvplease/internal/shared/rest"
	"github.com/sirupsen/logrus"
)

func NewAuthenticateHandler(s domain.ServiceContract, jwt *token.JWTService, logger *logrus.Logger) http.HandlerFunc {
	return makeAuthenticateEndpoint(s, jwt, logger)
}

type authenticateRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type authenticateResponse struct {
	Token string `json:"token"`
}

func makeAuthenticateEndpoint(s domain.ServiceContract, jwt *token.JWTService, logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := authenticateRequest{}
		if err := rest.ParseRequest[authenticateRequest](&req, r.Body); err != nil {
			logger.Errorf("Error parsing create account request")
			rest.NewResponse[error](w, "Payload error", &err, http.StatusBadRequest)
			return
		}

		acc, err := s.Login(req.Email, req.Password)
		if err != nil {
			logger.Error(err.Error())
			rest.NewResponse[any](w, err.Error(), nil, http.StatusUnauthorized)
			return
		}

		tokenStr, err := jwt.GenerateToken(acc.ID, acc.Role)
		if err != nil {
			logger.Error(err.Error())
			rest.NewResponse[any](w, err.Error(), nil, http.StatusUnauthorized)
			return
		}

		body := &authenticateResponse{
			Token: tokenStr,
		}

		rest.NewResponse[authenticateResponse](
			w,
			"Authenticated successfully",
			body,
			http.StatusOK,
		)
	}
}
