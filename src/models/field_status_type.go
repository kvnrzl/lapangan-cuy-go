package models

import "github.com/google/uuid"

// ACTIVE
// INACTIVE

type FieldStatusType struct {
	ID   uuid.UUID `gorm:"primaryKey"`
	Name string
}
