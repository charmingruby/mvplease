package domain

import "github.com/google/uuid"

type ServiceContract interface {
	CreateGroup(group *Group) error
	TransferGroupOwnership(groupID uuid.UUID, newAccountOwnerID uuid.UUID) error
	DeleteGroup(groupID uuid.UUID) error

	FindGroupsByAccountID(accountID uuid.UUID) ([]Group, error)

	Members(groupID uuid.UUID, page uint) ([]GroupMember, error)
	AddMember(groupID uuid.UUID, accountID uuid.UUID) error
	RemoveMember(groupID uuid.UUID, accountID uuid.UUID) error
}

type Service struct {
	groups  GroupRepository
	members MemberRepository
}

func NewService(
	groups GroupRepository,
	members MemberRepository,
) *Service {
	return &Service{
		groups:  groups,
		members: members,
	}
}
