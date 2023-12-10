package models

import "github.com/google/uuid"

// PENDING
// ACCEPT

type BookingStatusType struct {
	ID   uuid.UUID `gorm:"primaryKey"`
	Name string
}
