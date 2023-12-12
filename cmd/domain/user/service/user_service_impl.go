package service

import (
	"com.lapangan.cuy/cmd/domain/user/repository"
	"com.lapangan.cuy/cmd/models"
	"context"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	userRepository repository.UserRepositoryPostgresImpl
	db             *gorm.DB
	validate       *validator.Validate
}

func (u *UserServiceImpl) UserRegister(ctx context.Context, payload []byte) (*models.User, error) {

	return nil, nil
}

func (u *UserServiceImpl) UserLogin(ctx context.Context, payload []byte) (token string, err error) {
	// decode the input (json) to object

	// validate the object

	// get the data from repository

	// compare the password using bcrypt

	// if no error, set session

	// return the token

	return "", nil
}

func (u *UserServiceImpl) UserLogout(ctx context.Context) error {

	return nil
}
