package models

import (
	"github.com/google/uuid"
	"time"
)

type OperationalTime struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	FieldID   uuid.UUID
	Field     Field `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Day       string
	StartTime time.Time
	EndTime   time.Time
	CreatedAt time.Time
	CreatedBy uuid.UUID
	UpdatedAt time.Time
	UpdatedBy uuid.UUID
	Remarks   string
}
