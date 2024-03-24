package endpoints

import (
	"fmt"
	"net/http"
	"time"

	"github.com/charmingruby/mvplease/internal/account/domain"
	"github.com/charmingruby/mvplease/internal/common/infra/rest"
	"github.com/charmingruby/mvplease/internal/common/infra/rest/multipart"
	"github.com/charmingruby/mvplease/internal/common/infra/storage"
	"github.com/charmingruby/mvplease/pkg/errors"
	"github.com/sirupsen/logrus"
)

func NewUploadAvatarHandler(s domain.ServiceContract, logger *logrus.Logger) http.HandlerFunc {
	return makeUploadAvatarEndpoint(s, logger)
}

func makeUploadAvatarEndpoint(s domain.ServiceContract, logger *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cloudflare, err := storage.NewStorageService(logger)
		if err != nil {
			cfErr := fmt.Sprintf("Cloudflare error: %s", err.Error())
			logger.Errorf(cfErr)
			rest.NewResponse[any](w, cfErr, nil, http.StatusInternalServerError)
			return
		}

		// Payload
		payload, err := rest.RetrievePayloadFromRequest(r)
		if err != nil {
			logger.Error(err.Error())
			rest.NewResponse[any](w, err.Error(), nil, http.StatusInternalServerError)
			return
		}

		// Get from multipart form
		file, entity, err := multipart.HandleMultipartFormFile(
			r,
			"Avatar",
			32,
			int64(storage.MBToBytes(10)),
			[]string{"jpg", "png", "jpeg"},
		)
		if err != nil {
			logger.Error(err.Error())
			rest.NewResponse[any](w, err.Error(), nil, http.StatusBadRequest)
			return
		}

		fileURL := fmt.Sprintf("%s-%d.%s", payload.AccountID, time.Now().Unix(), entity.Mimetype)

		// Create upload
		if err = cloudflare.Upload(file, fileURL); err != nil {
			logger.Error(err)
			rest.NewResponse[any](w, err.Error(), nil, http.StatusInternalServerError)
			return
		}

		// Service
		if err := s.UploadAvatar(payload.AccountID, fileURL); err != nil {
			if err := cloudflare.DeleteByKey(fileURL); err != nil {
				logger.Error(err.Error())
				rest.NewResponse[any](w, err.Error(), nil, http.StatusInternalServerError)
				return
			}

			notFoundError, ok := err.(*errors.NotFoundError)
			if ok {
				logger.Error(notFoundError)
				rest.NewResponse[any](w, notFoundError.Error(), nil, http.StatusNotFound)
				return
			}

			logger.Error(err)
			rest.NewResponse[any](w, err.Error(), nil, http.StatusBadRequest)
			return
		}

		msg := "Avatar updated successfully."
		logger.Info(msg)
		rest.NewResponse[any](w, msg, nil, http.StatusOK)
	}
}
