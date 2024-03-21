package middlewares

import (
	"github.com/charmingruby/mvplease/internal/common/infra/token"
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
