package domain

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID                 uint       `db:"id" json:"id"`
	UUID               uuid.UUID  `db:"uuid" json:"uuid"`
	Name               string     `db:"name" json:"name"`
	Email              string     `db:"email" json:"email"`
	Role               string     `db:"role" json:"role"`
	AvatarURL          string     `db:"avatar_url" json:"avatar_url"`
	Password           string     `db:"password" json:"password"`
	AggregatesQuantity uint       `db:"aggregates_quantity" json:"aggregates_quantity"`
	ExamplesQuantity   uint       `db:"examples_quantity" json:"examples_quantity"`
	DeletedBy          string     `db:"deleted_by" json:"deleted_by"`
	CreatedAt          time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt          *time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt          *time.Time `db:"deleted_at" json:"deleted_at"`
}
