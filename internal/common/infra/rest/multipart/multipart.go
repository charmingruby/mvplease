package multipart

import (
	"mime/multipart"
	"net/http"

	"github.com/charmingruby/mvplease/internal/common/infra/storage"
	"github.com/charmingruby/mvplease/pkg/errors"
)

func HandleMultipartFormFile(
	r *http.Request,
	key string,
	multiformMemory int64,
	fileMaxSizeInBytes int64,
	validMimetypes []string) (multipart.File, *storage.File, error) {
	r.ParseMultipartForm(multiformMemory << 20)
	multipartFormKey := key
	file, fileHeader, err := r.FormFile(multipartFormKey)
	if err != nil {
		return nil, nil, errors.NewEmptyFileError(err, key)
	}

	// Validate file
	filename, mimetype, err := storage.SplitFileData(fileHeader.Filename)
	if err != nil {
		return nil, nil, err
	}

	//validMimetypes := []string{"jpg", "png", "jpeg"}
	//maxSizeInBytes := 10000000 // 10 mb

	fileEntity, err := storage.NewFile(filename, mimetype, fileHeader.Size, validMimetypes, int64(fileMaxSizeInBytes))
	if err != nil {
		return nil, nil, err
	}

	return file, fileEntity, nil
}
