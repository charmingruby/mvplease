package http

import (
	"net/http"

	"github.com/charmingruby/mvplease/internal/account/domain"
	"github.com/charmingruby/mvplease/internal/account/transport/rest/endpoints"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	Service domain.ServiceContract
	Logger  *logrus.Logger
}

func NewHTTPHandler(r *mux.Router, service domain.ServiceContract, logger *logrus.Logger) {
	logger.Info("Registering account routes...")

	h := &Handler{
		Service: service,
		Logger:  logger,
	}

	accountsRouter := h.NewAccountRouter()

	r.PathPrefix("/accounts").Handler(accountsRouter)
	logger.Info("Registered account routes.")
}

func (h *Handler) NewAccountRouter() *mux.Router {
	r := mux.NewRouter().PathPrefix("/accounts").Subrouter().StrictSlash(true)

	createAccountHandler := endpoints.NewCreateAccountHandler(h.Service, h.Logger)
	r.Handle("", createAccountHandler).Methods(http.MethodPost)

	return r
}
