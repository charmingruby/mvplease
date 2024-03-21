package domain

import "github.com/google/uuid"

type MemberRepository interface {
	FindMemberByID(id uuid.UUID) (GroupMember, error)
	FetchMembersByAccountID(memberID uuid.UUID, page uint) ([]GroupMember, error)
	CreateMember(m *GroupMember) error
	SaveMember(m *GroupMember) error
	DeleteMember(m *GroupMember) error
}
