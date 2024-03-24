package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/charmingruby/mvplease/internal/common/infra/security"
	cErrors "github.com/charmingruby/mvplease/pkg/errors"
	"github.com/go-playground/validator/v10"
)

func NewResponse[T any](
	w http.ResponseWriter,
	message string,
	data *T,
	statusCode int,
) {
	res := Response{
		Message:    message,
		StatusCode: statusCode,
		Data:       nil,
	}

	if data != nil {
		res.Data = data
	}

	writeResponse(w, &res)
}

type Response struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Data       any    `json:"data,omitempty"`
}

func (r *Response) marshal() []byte {
	jsonResponse, err := json.Marshal(r)

	if err != nil {
		fmt.Printf("Failed to marshal response: %v", err)
	}

	return jsonResponse
}

func writeResponse(w http.ResponseWriter, r *Response) {
	jsonResponse := r.marshal()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)

	_, err := w.Write(jsonResponse)
	if err != nil {
		log.Printf("Error writing response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func ParseRequest[T any](r *T, body io.ReadCloser) error {
	if err := json.NewDecoder(body).Decode(&r); err != nil {
		return err
	}

	return nil
}

type ValidationErrors struct {
	Errors map[string][]string `json:"errors,omitempty"`
}

func NewValidationErrors(err error) *ValidationErrors {
	var validationErrors validator.ValidationErrors
	errors.As(err, &validationErrors)

	fieldErrors := make(map[string][]string)
	for _, vErr := range validationErrors {
		fieldName := vErr.Field()
		fieldError := fieldName + " " + vErr.Tag()

		fieldErrors[fieldName] = append(fieldErrors[fieldName], fieldError)
	}

	return &ValidationErrors{Errors: fieldErrors}
}

func IsRequestValid(request any) *ValidationErrors {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(request)
	if err != nil {
		return NewValidationErrors(err)
	}

	return nil
}

func RetrieveTokenFromRequest(r *http.Request) string {
	fullToken := r.Header.Get("Authorization")

	splittedToken := strings.Split(fullToken, " ")

	if len(splittedToken) != 2 {
		return ""
	}

	return splittedToken[1]
}

func RetrievePayloadFromRequest(r *http.Request) (*security.Payload, error) {
	tokenStr := RetrieveTokenFromRequest(r)

	payload, err := security.NewJWTService().RetriveTokenPayload(tokenStr)

	if err != nil {
		return nil, cErrors.NewTokenRetrieveError(err)
	}

	return payload, nil
}
