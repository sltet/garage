package core

import uuid "github.com/satori/go.uuid"

func GetTimeBasedUUID() uuid.UUID {
	return uuid.NewV1()
}
