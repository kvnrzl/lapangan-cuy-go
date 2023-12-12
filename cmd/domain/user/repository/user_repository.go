package repository

import (
	"com.lapangan.cuy/cmd/models"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, db *gorm.DB, user *models.User) (tx *gorm.DB)
	Update(ctx context.Context, db *gorm.DB, user *models.User) (tx *gorm.DB)
	SelectById(ctx context.Context, db *gorm.DB, id uuid.UUID) (user models.User)
	SelectByUsername(ctx context.Context, db *gorm.DB, username string)
	DeleteById(ctx context.Context, db *gorm.DB, id uuid.UUID)
	DeleteByUsername(ctx context.Context, db *gorm.DB, username string)
}
