package http

import (
	"net/http"

	"github.com/charmingruby/mvplease/internal/account/domain"
	"github.com/charmingruby/mvplease/internal/account/transport/rest/endpoints"
	"github.com/charmingruby/mvplease/internal/common/infra/rest/middlewares"
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

	createAccountHandler := endpoints.NewCreateAccountHandler(h.s, h.logger)
	r.Handle("/register", createAccountHandler).Methods(http.MethodPost)

	// Sessions group
	sessionsRouter := h.NewSessionsRouter()
	r.PathPrefix("/sessions").Handler(sessionsRouter)

	// Accounts group
	accountsRouter := h.NewAccountsRouter()
	r.PathPrefix("/accounts").Handler(accountsRouter)

	// Profiles group
	profilesRouter := h.NewProfilesRouter()
	r.PathPrefix("/me").Handler(profilesRouter)

	logger.Info("Registered account routes.")
}

func (h *Handler) NewAccountsRouter() *mux.Router {
	r := mux.NewRouter().PathPrefix("/accounts").Subrouter().StrictSlash(true)

	deleteAccountHandler := endpoints.NewDeleteAccountHandler(h.s, h.logger)
	r.Handle("/{id}", h.mw.ProtectedRouteByRole("manager", deleteAccountHandler)).Methods(http.MethodDelete)

	return r
}

func (h *Handler) NewSessionsRouter() *mux.Router {
	r := mux.NewRouter().PathPrefix("/sessions").Subrouter().StrictSlash(true)

	authenticateHandler := endpoints.NewAuthenticateHandler(h.s, h.logger)
	r.Handle("", authenticateHandler).Methods(http.MethodPost)

	return r
}

func (h *Handler) NewProfilesRouter() *mux.Router {
	r := mux.NewRouter().PathPrefix("/me").Subrouter().StrictSlash(true)

	profileHandler := endpoints.NewProfileHandler(h.s, h.logger)
	r.Handle("", h.mw.ProtectdRoute(profileHandler)).Methods(http.MethodGet)

	uploadAvatarHandler := endpoints.NewUploadAvatarHandler(h.s, h.logger)
	r.Handle("/avatar", h.mw.ProtectdRoute(uploadAvatarHandler)).Methods(http.MethodPatch)

	return r
}
