package endpoints

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/charmingruby/mvplease/internal/account/domain"
	"github.com/charmingruby/mvplease/internal/common/infra/rest"
	"github.com/sirupsen/logrus"
)

func NewFetchAccountsHandler(s domain.ServiceContract, logger *logrus.Logger) http.HandlerFunc {
	return makeFetchAccountsEndpoint(s, logger)
}

type fetchAccountsResponse struct {
	Page            int              `json:"page"`
	AccountsFetched int              `json:"accounts_fetched"`
	Accounts        []domain.Account `json:"accounts"`
}

func makeFetchAccountsEndpoint(s domain.ServiceContract, logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()

		pageParam := queryParams.Get("page")
		page, err := strconv.Atoi(pageParam)
		if err != nil {
			page = 1
		}

		if page < 1 {
			rest.NewResponse[any](w, fmt.Sprintf("Invalid page param: %s", err.Error()), nil, http.StatusBadRequest)
			return
		}

		accounts, err := s.Accounts(uint(page - 1))
		if err != nil {
			logger.Error(err.Error())
			rest.NewResponse[any](w, err.Error(), nil, http.StatusBadRequest)
			return
		}

		res := &fetchAccountsResponse{
			Page:            page,
			AccountsFetched: len(accounts),
			Accounts:        accounts,
		}
		msg := fmt.Sprintf("Fetched %d accounts in page %d", len(accounts), page)
		logger.Info(msg)
		rest.NewResponse[fetchAccountsResponse](
			w,
			msg,
			res,
			http.StatusOK,
		)
	}
}
