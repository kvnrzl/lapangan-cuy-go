package service

import (
	"com.lapangan.cuy/cmd/common"
	"com.lapangan.cuy/cmd/domain/user/repository"
	"com.lapangan.cuy/cmd/models"
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepositoryPostgresImpl
	DB             *gorm.DB
	RedisClient    *redis.Client
	Validate       *validator.Validate
}

func (u *UserServiceImpl) Register(ctx context.Context, payload []byte) (*models.User, error) {

	return nil, nil
}

func (u *UserServiceImpl) Login(ctx context.Context, payload []byte) (string, string, error) {
	// decode the input (json) to struct
	var user models.User
	if err := json.Unmarshal(payload, &user); err != nil {
		newError := common.HandleError(err, "error - unmarshal login payload")

		return "", "", newError
	}

	// validate the struct
	if err := u.Validate.StructCtx(ctx, user); err != nil {
		newError := common.HandleError(err, "error - validate struct login")

		return "", "", newError
	}

	// get the data from repository
	userFromDB, res := u.UserRepository.FindUserByEmail(ctx, u.DB, user.Email)
	if res.Error != nil {
		newError := common.HandleError(res.Error, "error - user is not found")

		return "", "", newError
	}

	// compare the password using bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(user.Password)); err != nil {
		newError := common.HandleError(err, "error - password is not match ")

		return "", "", newError
	}

	// generate the token (access and refresh token)
	accessToken, refreshToken, err := common.GenerateToken(userFromDB.ID)
	if err != nil {
		newError := common.HandleError(err, err.Error())

		return "", "", newError
	}

	// save it to redis
	expAccessTokenString := os.Getenv("EXP_ACCESS_TOKEN")
	expAccessToken, err := strconv.Atoi(expAccessTokenString)
	if err != nil {
		newError := common.HandleError(err, "error - convert string to int")
		return "", "", newError
	}

	err = u.RedisClient.Set(ctx, userFromDB.ID.String(), accessToken, time.Minute*time.Duration(expAccessToken)).Err()
	if err != nil {
		newError := common.HandleError(err, "error - set redis client")

		return "", "", newError
	}

	return accessToken, refreshToken, nil
}

func (u *UserServiceImpl) Logout(ctx context.Context, payload []byte) error {
	// decode the input (json) to struct
	var user models.User
	if err := json.Unmarshal(payload, &user); err != nil {
		newError := common.HandleError(err, "error - unmarshal logout payload")

		return newError
	}

	// remove the token from redis
	err := u.RedisClient.Del(ctx, user.ID.String()).Err()
	if err != nil {
		newError := common.HandleError(err, "error - delete redis client")

		return newError
	}

	return nil
}

func (u *UserServiceImpl) ForgotPassword(ctx context.Context) error {

	return nil
}

func (u *UserServiceImpl) ResetPassword(ctx context.Context) error {

	return nil
}
