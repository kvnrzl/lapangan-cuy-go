package models

import "github.com/google/uuid"

type OrderPaymentType struct {
	ID   uuid.UUID `gorm:"primaryKey"`
	Name string
}
