package models

import "github.com/google/uuid"

type Role struct {
	ID   uuid.UUID `gorm:"primaryKey"`
	Name string
}
