package models

import (
	"github.com/google/uuid"
	"time"
)

type Image struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Path        string
	Name        string
	Description string `gorm:"type:text"`
	FieldID     uuid.UUID
	Field       Field `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsEnabled   bool
	IsDeleted   bool
	CreatedAt   time.Time
	CreatedBy   uuid.UUID
	UpdatedAt   time.Time
	UpdatedBy   uuid.UUID
}
