package models

import (
	"github.com/google/uuid"
	"time"
)

type UserFavoriteField struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	UserID    uuid.UUID
	User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	FieldID   uuid.UUID
	Field     Field `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsEnabled bool
	IsDeleted bool
	CreatedAt time.Time
	CreatedBy uuid.UUID
	UpdatedAt time.Time
	UpdatedBy uuid.UUID
}
