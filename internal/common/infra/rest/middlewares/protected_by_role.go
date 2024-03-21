package middlewares

import (
	"net/http"

	"github.com/charmingruby/mvplease/internal/common/infra/rest"
)

func (m *Middleware) ProtectedRouteByRole(role string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr := rest.RetrieveTokenFromRequest(r)

		payload, err := m.jwt.RetriveTokenPayload(tokenStr)
		if err != nil {
			NewUnauthorizedError(w)
			return
		}

		if payload.Role != role {
			NewUnauthorizedError(w)
			return
		}

		next(w, r)
	}
}
