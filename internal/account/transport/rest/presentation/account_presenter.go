package presentation

import (
	"time"

	"github.com/charmingruby/mvplease/internal/account/domain"
	"github.com/google/uuid"
)

type AccountPresenter struct {
	ID                 uint       `json:"id"`
	UUID               uuid.UUID  `json:"uuid"`
	Name               string     `json:"name"`
	Email              string     `json:"email"`
	Role               string     `json:"role"`
	AvatarURL          string     `json:"avatar_url"`
	AggregatesQuantity uint       `json:"aggregates_quantity"`
	ExamplesQuantity   uint       `json:"example_quantity"`
	DeletedBy          string     `json:"deleted_by"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          *time.Time `json:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at"`
}

func NewHTTPAccountPresenter(a *domain.Account)
