package middlewares

import (
	"net/http"

	"github.com/charmingruby/mvplease/internal/common/infra/rest"
)

func (m *Middleware) ProtectdRoute(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr := rest.RetrieveTokenFromRequest(r)

		if tokenStr == "" {
			m.logger.Error("Middleware: Token not found")
			NewUnauthorizedError(w)
			return
		}

		if isTokenValid := m.jwt.ValidateToken(tokenStr); !isTokenValid {
			m.logger.Error("Middleware: Token is not valid")
			NewUnauthorizedError(w)
			return
		}

		next(w, r)
	}
}
