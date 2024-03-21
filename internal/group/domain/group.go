package domain

import (
	"time"

	"github.com/charmingruby/mvplease/internal/common/core"
	"github.com/google/uuid"
)

func NewGroup(
	name, description string,
	ownerID uuid.UUID,
) *Group {
	return &Group{
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

type Group struct {
	ID              uuid.UUID  `db:"id" json:"id"`
	Name            string     `db:"name" json:"name"`
	Description     string     `db:"description" json:"description"`
	MembersQuantity uint       `db:"members_quantity" json:"members_quantity"`
	OwnerID         uuid.UUID  `db:"owner_id" json:"owner_id"`
	DeletedBy       *uuid.UUID `db:"deleted_by" json:"deleted_by,omitempty"`
	CreatedAt       time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt       *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	DeletedAt       *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}
