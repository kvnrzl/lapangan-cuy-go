package models

import "github.com/google/uuid"

type PaymentType struct {
	ID   uuid.UUID `gorm:"primaryKey"`
	Name string
}
