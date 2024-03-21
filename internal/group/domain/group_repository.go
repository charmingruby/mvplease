package domain

import "github.com/google/uuid"

type GroupRepository interface {
	FindGroupByID(id uuid.UUID) (Group, error)
	FetchGroupsByAccountID(groupID uuid.UUID, page uint) ([]Group, error)
	CreateGroup(g *Group) error
	SaveGroup(g *Group) error
	DeleteGroup(g *Group) error
}
