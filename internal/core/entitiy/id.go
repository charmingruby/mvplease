package entity

import "github.com/google/uuid"

func NewID() uuid.UUID {
	return uuid.New()
}
