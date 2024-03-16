package middlewares

import (
	"net/http"

	"github.com/charmingruby/mvplease/internal/shared/rest"
)

func NewUnauthorizedError(w http.ResponseWriter) {
	rest.NewResponse[any](w, "Unauthorized", nil, http.StatusUnauthorized)
}
