package domain

import (
	"time"

	entity "github.com/charmingruby/mvplease/internal/core/entitiy"
	"github.com/google/uuid"
)

const (
	defaultRole = "default"
	adminRole   = "admin"
)

func NewAccount(
	name,
	email,
	password string,
) *Account {
	a := Account{
		ID:                 entity.NewID(),
		Name:               name,
		Email:              email,
		Role:               accountRoles()[defaultRole],
		AvatarURL:          nil,
		Password:           password,
		AggregatesQuantity: 0,
		ExamplesQuantity:   0,
		DeletedBy:          nil,
		CreatedAt:          time.Now(),
		UpdatedAt:          nil,
		DeletedAt:          nil,
	}

	return &a
}

type Account struct {
	ID                 uuid.UUID  `db:"id" json:"id"`
	Name               string     `db:"name" json:"name"`
	Email              string     `db:"email" json:"email"`
	Role               string     `db:"role" json:"role"`
	AvatarURL          *string    `db:"avatar_url" json:"avatar_url"`
	Password           string     `db:"password" json:"password"`
	AggregatesQuantity uint       `db:"aggregates_quantity" json:"aggregates_quantity"`
	ExamplesQuantity   uint       `db:"examples_quantity" json:"examples_quantity"`
	DeletedBy          *uuid.UUID `db:"deleted_by" json:"deleted_by"`
	CreatedAt          time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt          *time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt          *time.Time `db:"deleted_at" json:"deleted_at"`
}

func (a *Account) SetPassword(password string) {
	a.Password = password
}

func accountRoles() map[string]string {
	return map[string]string{
		defaultRole: "member",
		adminRole:   "manager",
	}
}
