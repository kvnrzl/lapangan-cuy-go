package repository

import (
	"com.lapangan.cuy/cmd/models"
	"context"
	"gorm.io/gorm"
)

type UserRepositoryPostgresImpl struct{}

//func NewUserRepositoryPostgresImpl() UserRepository {
//	return &UserRepositoryPostgresImpl{}
//}

func (u *UserRepositoryPostgresImpl) Create(ctx context.Context, db *gorm.DB, user *models.User) *gorm.DB {
	return db.Create(user)
}

func (u *UserRepositoryPostgresImpl) Update(ctx context.Context, db *gorm.DB, user *models.User) *gorm.DB {
	return db.Save(user)
}

func (u *UserRepositoryPostgresImpl) Get(ctx context.Context, db *gorm.DB, user *models.User) (models.User, *gorm.DB) {
	res := db.First(&user)

	return *user, res
}

func (u *UserRepositoryPostgresImpl) Delete(ctx context.Context, db *gorm.DB, user *models.User) *gorm.DB {
	return db.Delete(&user)
}
