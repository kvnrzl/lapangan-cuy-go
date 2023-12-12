package models

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	ID                 uuid.UUID `gorm:"primaryKey"`
	StartTime          time.Time
	EndTime            time.Time
	TotalHours         uint
	TotalCost          float64
	TotalPayment       float64
	UserID             uuid.UUID
	User               User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FieldID            uuid.UUID
	Field              Field `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OrderPaymentTypeID uuid.UUID
	OrderPaymentType   OrderPaymentType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsEnabled          bool
	CreatedAt          time.Time
	CreatedBy          uuid.UUID
	Remarks            string
}
