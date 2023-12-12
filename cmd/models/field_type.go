package models

import "github.com/google/uuid"

// FUTSAL
// MINISOCCER
// BOLA

type FieldType struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Name     string
	IconPath string
}
