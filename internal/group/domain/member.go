package domain

import (
	"time"

	"github.com/charmingruby/mvplease/internal/common/core"
	"github.com/google/uuid"
)

func NewGroupMember(
	groupID uuid.UUID,
	accountID uuid.UUID,
) *GroupMember {
	return &GroupMember{
		ID:        core.NewID(),
		GroupID:   groupID,
		AccountID: accountID,
		DeletedBy: nil,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
		DeletedAt: nil,
	}
}

type GroupMember struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	GroupID   uuid.UUID  `db:"example_id" json:"example_id"`
	AccountID uuid.UUID  `db:"account_id" json:"account_id"`
	DeletedBy *uuid.UUID `db:"deleted_by" json:"deleted_by,omitempty"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at,omitempty"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at,omitempty"`
}
