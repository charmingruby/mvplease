package http

import (
	"net/http"

	"github.com/charmingruby/mvplease/internal/account/domain"
	"github.com/charmingruby/mvplease/internal/account/transport/rest/endpoints"
	"github.com/charmingruby/mvplease/internal/services/token"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Service domain.ServiceContract
	Logger  *logrus.Logger
}

func NewHTTPHandler(r *mux.Router, jwtService *token.JWTService, service domain.ServiceContract, logger *logrus.Logger) {
	logger.Info("Registering account routes...")

	h := &Handler{
		Service: service,
		Logger:  logger,
	}

	accountsRouter := h.NewAccountRouter()
	sessionsRouter := h.NewSessionsRouter(jwtService)

	r.PathPrefix("/accounts").Handler(accountsRouter)
	r.PathPrefix("/sessions").Handler(sessionsRouter)

	logger.Info("Registered account routes.")
}

func (h *Handler) NewAccountRouter() *mux.Router {
	r := mux.NewRouter().PathPrefix("/accounts").Subrouter().StrictSlash(true)

	createAccountHandler := endpoints.NewCreateAccountHandler(h.Service, h.Logger)
	r.Handle("", createAccountHandler).Methods(http.MethodPost)

	return r
}

func (h *Handler) NewSessionsRouter(jwtService *token.JWTService) *mux.Router {
	r := mux.NewRouter().PathPrefix("/sessions").Subrouter().StrictSlash(true)

	authenticateHandler := endpoints.NewAuthenticateHandler(h.Service, jwtService, h.Logger)
	r.Handle("", authenticateHandler).Methods(http.MethodPost)

	return r
}
