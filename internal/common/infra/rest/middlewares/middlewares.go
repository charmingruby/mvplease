package middlewares

import (
	"github.com/charmingruby/mvplease/internal/common/infra/security"
	"github.com/sirupsen/logrus"
)

func NewMiddleware(logger *logrus.Logger) *Middleware {
	return &Middleware{
		logger: logger,
		jwt:    security.NewJWTService(),
	}

}

type Middleware struct {
	jwt    *security.JWTService
	logger *logrus.Logger
}
