package example

import (
	"time"

	entity "github.com/charmingruby/mvplease/internal/core/entitiy"
	"github.com/google/uuid"
)

func NewAggregate(
	name,
	description string,
	ownerID uuid.UUID,
) *Aggregate {
	a := Aggregate{
		ID:               entity.NewID(),
		Name:             name,
		Description:      description,
		MembersQuantity:  0,
		ExamplesQuantity: 0,
		OwnerID:          ownerID,
		DeletedBy:        nil,
		CreatedAt:        time.Now(),
		UpdatedAt:        nil,
		DeletedAt:        nil,
	}

	return &a
}

type Aggregate struct {
	ID               uuid.UUID  `db:"id" json:"id"`
	Name             string     `db:"name" json:"name"`
	Description      string     `db:"description" json:"description"`
	MembersQuantity  uint       `db:"members_quantity" json:"members_quantity"`
	ExamplesQuantity uint       `db:"examples_quantity" json:"examples_quantity"`
	OwnerID          uuid.UUID  `db:"owner_id" json:"owner_id"`
	DeletedBy        *uuid.UUID `db:"deleted_by" json:"deleted_by"`
	CreatedAt        time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt        *time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt        *time.Time `db:"deleted_at" json:"deleted_at"`
}

func (a *Aggregate) Touch() {
	now := time.Now()
	a.UpdatedAt = &now
}
