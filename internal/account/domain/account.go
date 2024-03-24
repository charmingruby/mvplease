package domain

import (
	"time"

	"github.com/charmingruby/mvplease/internal/common/core"
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
		ID:             core.NewID(),
		Name:           name,
		Email:          email,
		Role:           accountRoles()[defaultRole],
		AvatarURL:      nil,
		Password:       password,
		GroupsQuantity: 0,
		DeletedBy:      nil,
		CreatedAt:      time.Now(),
		UpdatedAt:      nil,
		DeletedAt:      nil,
	}

	return &a
}

type Account struct {
	ID             uuid.UUID  `db:"id" json:"id"`
	Name           string     `db:"name" json:"name"`
	Email          string     `db:"email" json:"email"`
	Role           string     `db:"role" json:"role"`
	AvatarURL      *string    `db:"avatar_url" json:"avatar_url"`
	Password       string     `db:"password" json:"password"`
	GroupsQuantity uint       `db:"groups_quantity" json:"groups_quantity"`
	DeletedBy      *uuid.UUID `db:"deleted_by" json:"deleted_by,omitempty"`
	CreatedAt      time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt      *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	DeletedAt      *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}

func (a *Account) SetAvatarURL(avatarURL string) {
	a.AvatarURL = &avatarURL
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
func (a *Account) DeleteAccount(managerID uuid.UUID) {
	now := time.Now()

	a.DeletedAt = &now
	a.DeletedBy = &managerID
	a.Touch()
}

func (a *Account) Touch() {
	now := time.Now()
	a.UpdatedAt = &now
}
