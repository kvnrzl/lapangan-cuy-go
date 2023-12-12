package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           uuid.UUID `gorm:"primaryKey"`
	Username     string
	Password     string
	FullName     string
	Email        string `gorm:"unique"`
	MobileNumber string `gorm:"unique"`
	LastActive   time.Time
	RoleID       uuid.UUID
	Role         UserRole `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt    time.Time
	CreatedBy    uuid.UUID
	UpdatedAt    time.Time
	UpdatedBy    uuid.UUID
	IsEnabled    bool
	IsDeleted    bool
	Remarks      string
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()

	return
}
