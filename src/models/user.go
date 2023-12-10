package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID `gorm:"primaryKey"`
	Username     string
	Password     string
	FullName     string
	Email        string
	MobileNumber string
	LastActive   time.Time
	RoleID       uuid.UUID
	Role         Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt    time.Time
	CreatedBy    uuid.UUID
	UpdatedAt    time.Time
	UpdatedBy    uuid.UUID
	IsEnabled    bool
	IsDeleted    bool
	Remarks      string
}
