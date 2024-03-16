package middlewares

import (
	"github.com/charmingruby/mvplease/internal/services/token"
	"github.com/sirupsen/logrus"
)

func NewMiddleware(logger *logrus.Logger) *Middleware {
	return &Middleware{
		logger: logger,
		jwt:    token.NewJWTService(),
	}

}

type Middleware struct {
	jwt    *token.JWTService
	logger *logrus.Logger
}
