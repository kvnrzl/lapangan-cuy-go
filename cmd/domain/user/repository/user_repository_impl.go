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

func (u *UserRepositoryPostgresImpl) CreateUser(ctx context.Context, db *gorm.DB, user *models.User) *gorm.DB {
	return db.WithContext(ctx).Create(user)
}

func (u *UserRepositoryPostgresImpl) UpdateUser(ctx context.Context, db *gorm.DB, user *models.User) *gorm.DB {
	return db.WithContext(ctx).Save(user)
}

func (u *UserRepositoryPostgresImpl) FindUserByEmail(ctx context.Context, db *gorm.DB, email string) (models.User, *gorm.DB) {
	var user models.User
	res := db.WithContext(ctx).Where("email = ?", email).First(&user)

	return user, res
}

func (u *UserRepositoryPostgresImpl) FindUser(ctx context.Context, db *gorm.DB, user *models.User) (models.User, *gorm.DB) {
	res := db.WithContext(ctx).First(&user)

	return *user, res
}

func (u *UserRepositoryPostgresImpl) Delete(ctx context.Context, db *gorm.DB, user *models.User) *gorm.DB {
	return db.WithContext(ctx).Delete(&user)
}
