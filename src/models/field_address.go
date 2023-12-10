package models

import (
	"github.com/google/uuid"
	"time"
)

type FieldAddress struct {
	ID         uuid.UUID `gorm:"primaryKey"`
	Name       string
	Country    string
	Province   string
	City       string
	District   string
	Village    string
	Street     string
	Detail     string
	PostalCode string
	Longitude  float64 `gorm:"type:decimal(10,8)"`
	Latitude   float64 `gorm:"type:decimal(11,8)"`
	CreatedAt  time.Time
	CreatedBy  uuid.UUID
	UpdatedAt  time.Time
	UpdatedBy  uuid.UUID
	Remarks    string
}
