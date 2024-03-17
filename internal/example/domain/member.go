package domain

import (
	"time"

	"github.com/charmingruby/mvplease/internal/core"
	"github.com/google/uuid"
)

func New(
	name, description string,
	ownerID uuid.UUID,
) *Example {
	return &Example{
		ID:              core.NewID(),
		Name:            name,
		Description:     description,
		MembersQuantity: 0,
		OwnerID:         ownerID,
		DeletedBy:       nil,
		CreatedAt:       time.Now(),
		UpdatedAt:       nil,
		DeletedAt:       nil,
	}
}

type ExampleMember struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	ExampleID string     `db:"example_id" json:"example_id"`
	AccountID uuid.UUID  `db:"account_id" json:"account_id"`
	DeletedBy *uuid.UUID `db:"deleted_by" json:"deleted_by,omitempty"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}
