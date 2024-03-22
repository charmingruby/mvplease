package http

import (
	"net/http"

	"github.com/charmingruby/mvplease/internal/account/domain"
	"github.com/charmingruby/mvplease/internal/account/transport/rest/endpoints"
	"github.com/charmingruby/mvplease/internal/common/infra/rest/middlewares"
	"github.com/charmingruby/mvplease/internal/common/infra/security"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	s      domain.ServiceContract
	mw     *middlewares.Middleware
	logger *logrus.Logger
}

func NewHTTPHandler(r *mux.Router, mw *middlewares.Middleware, service domain.ServiceContract, logger *logrus.Logger) {
	logger.Info("Registering account routes...")

	h := &Handler{
		s:      service,
		mw:     mw,
		logger: logger,
	}

	jwtService := security.NewJWTService()

	createAccountHandler := endpoints.NewCreateAccountHandler(h.s, h.logger)
	r.Handle("/register", createAccountHandler).Methods(http.MethodPost)

	profileHandler := endpoints.NewProfileHandler(h.s, h.logger, jwtService)
	r.Handle("/me", h.mw.ProtectdRoute(profileHandler)).Methods(http.MethodGet)

	// Route Groups
	sessionsRouter := h.NewSessionsRouter(jwtService)
	r.PathPrefix("/sessions").Handler(sessionsRouter)

	accountsRouter := h.NewAccountsRouter()
	r.PathPrefix("/accounts").Handler(accountsRouter)

	logger.Info("Registered account routes.")
}

func (h *Handler) NewAccountsRouter() *mux.Router {
	r := mux.NewRouter().PathPrefix("/accounts").Subrouter().StrictSlash(true)

	deleteAccountHandler := endpoints.NewDeleteAccountHandler(h.s, h.logger)
	r.Handle("/{id}", h.mw.ProtectedRouteByRole("manager", deleteAccountHandler)).Methods(http.MethodDelete)

	return r
}

func (h *Handler) NewSessionsRouter(jwtService *security.JWTService) *mux.Router {
	r := mux.NewRouter().PathPrefix("/sessions").Subrouter().StrictSlash(true)

	authenticateHandler := endpoints.NewAuthenticateHandler(h.s, jwtService, h.logger)
	r.Handle("", authenticateHandler).Methods(http.MethodPost)

	return r
}
