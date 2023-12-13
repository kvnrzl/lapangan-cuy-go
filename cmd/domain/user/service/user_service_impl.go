package service

import (
	"com.lapangan.cuy/cmd/common"
	"com.lapangan.cuy/cmd/domain/user/repository"
	"com.lapangan.cuy/cmd/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepositoryPostgresImpl
	DB             *gorm.DB
	Validate       *validator.Validate
}

func (u *UserServiceImpl) UserRegister(ctx context.Context, payload []byte) (*models.User, error) {

	return nil, nil
}

func (u *UserServiceImpl) UserLogin(ctx context.Context, payload []byte) (token string, err error) {
	// decode the input (json) to object
	var user models.User
	if err := json.Unmarshal(payload, &user); err != nil {
		common.LogOnError(err, fmt.Sprintf("cannot unmarshal user (%v)", string(payload)))

		return "", err
	}

	// validate the object
	if err := u.Validate.StructCtx(ctx, user); err != nil {
		common.LogOnError(err, fmt.Sprintf("error validate struct (%v)", user))

		return "", err
	}

	// get the data from repository
	responseUserFromDB, res := u.UserRepository.FindUserByEmail(ctx, u.DB, user.Email)
	if res.Error != nil {
		common.LogOnError(err, fmt.Sprintf("error get user by email (%v)", responseUserFromDB))

		return "", err
	}

	// compare the password using bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(responseUserFromDB.Password), []byte(user.Password)); err != nil {
		common.LogOnError(err, fmt.Sprintf("password is not match (%v)", responseUserFromDB))

		return "", err
	}

	// if no error, set session

	// return the token

	return "", nil
}

func (u *UserServiceImpl) UserLogout(ctx context.Context) error {

	return nil
}
