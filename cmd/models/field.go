package models

import (
	"github.com/google/uuid"
	"time"
)

type Field struct {
	ID             uuid.UUID `gorm:"primaryKey"`
	Name           string
	FieldAddressID uuid.UUID
	FieldAddress   FieldAddress `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OwnerId        uuid.UUID
	Owner          User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Description    string `gorm:"type:text"`
	FieldTypeId    uuid.UUID
	FieldType      FieldType `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ThumbnailPath  string
	Cost           float64
	IsEnabled      bool
	IsDeleted      bool
	CreatedAt      time.Time
	CreatedBy      uuid.UUID
	UpdatedAt      time.Time
	UpdatedBy      uuid.UUID
	Remarks        string
}
