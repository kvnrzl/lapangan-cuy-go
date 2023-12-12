package models

import (
	"github.com/google/uuid"
	"time"
)

type Booking struct {
	ID                  uuid.UUID `gorm:"primaryKey"`
	BookingStatusTypeID uuid.UUID
	BookingStatusType   BookingStatusType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsEnabled           bool
	IsDeleted           bool
	CreatedAt           time.Time
	CreatedBy           uuid.UUID
	UpdatedAt           time.Time
	UpdatedBy           uuid.UUID
	Remarks             string
}
