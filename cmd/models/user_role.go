package models

import "github.com/google/uuid"

type UserRole struct {
	ID   uuid.UUID `gorm:"primaryKey"`
	Name string
}
