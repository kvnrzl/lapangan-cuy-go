package models

import (
	"github.com/google/uuid"
	"time"
)

type Booking struct {
	ID                  uuid.UUID `gorm:"primaryKey"`
	StartTime           time.Time
	EndTime             time.Time
	TotalHours          uint
	TotalCost           float64
	DownPayment         float64
	UserID              uuid.UUID
	User                User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FieldID             uuid.UUID
	Field               Field `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BookingStatusTypeID uuid.UUID
	BookingStatusType   BookingStatusType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PaymentTypeID       uuid.UUID
	PaymentType         PaymentType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsEnabled           bool
	IsDeleted           bool
	CreatedAt           time.Time
	CreatedBy           uuid.UUID
	UpdatedAt           time.Time
	UpdatedBy           uuid.UUID
	Remarks             string
}
