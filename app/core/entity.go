package core

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type ORMAwareEntity interface {
	GetID() string
}

type Entity struct {
	ID        uuid.UUID `json:"id" gorm:"type:char(36);primary_key;"`
	CreatedAt time.Time // Automatically managed by GORM for creation time
	UpdatedAt time.Time // Automatically managed by GORM for update time
}

func NewEntity() Entity {
	return Entity{}
}

// BeforeCreate will set a UUID rather than numeric ID.
func (e *Entity) BeforeCreate(_ *gorm.DB) error {
	e.ID = uuid.NewV1()
	return nil
}

func (c Entity) GetID() string {
	return c.ID.String()
}
