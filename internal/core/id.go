package core

import "github.com/google/uuid"

func NewID() uuid.UUID {
	return uuid.New()
}

func IsUUID(s string) bool {
	_, err := uuid.Parse(s)
	return err == nil
}
