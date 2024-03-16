package http

import (
	"net/http"

	"github.com/charmingruby/mvplease/internal/account/domain"
	"github.com/charmingruby/mvplease/internal/account/transport/rest/endpoints"
	"github.com/charmingruby/mvplease/internal/services/token"
	"github.com/charmingruby/mvplease/internal/shared/rest/middlewares"
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

	jwtService := token.NewJWTService()

	createAccountHandler := endpoints.NewCreateAccountHandler(h.s, h.logger)
	r.Handle("/register", createAccountHandler).Methods(http.MethodPost)

	profileHandler := endpoints.NewProfileHandler(h.s, h.logger, jwtService)
	r.Handle("/me", h.mw.ProtectdRoute(profileHandler)).Methods(http.MethodGet)

	sessionsRouter := h.NewSessionsRouter(jwtService)
	r.PathPrefix("/sessions").Handler(sessionsRouter)

	logger.Info("Registered account routes.")
}

func (h *Handler) NewAccountRouter() *mux.Router {
	r := mux.NewRouter().PathPrefix("/accounts").Subrouter().StrictSlash(true)

	return r
}

func (h *Handler) NewSessionsRouter(jwtService *token.JWTService) *mux.Router {
	r := mux.NewRouter().PathPrefix("/sessions").Subrouter().StrictSlash(true)

	authenticateHandler := endpoints.NewAuthenticateHandler(h.s, jwtService, h.logger)
	r.Handle("", authenticateHandler).Methods(http.MethodPost)

	return r
}
