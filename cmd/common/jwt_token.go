package common

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"log"
	"os"
	"strconv"
	"time"
)

type Claims struct {
	UserID uuid.UUID
	jwt.RegisteredClaims
}

func GenerateToken(userID uuid.UUID) (string, string, error) {

	expAccessTokenString := os.Getenv("EXP_ACCESS_TOKEN")
	expAccessToken, err := strconv.Atoi(expAccessTokenString)
	if err != nil {
		newError := HandleError(err, "error convert string to int")
		return "", "", newError
	}

	// prepare the object Claims
	accessTokenClaims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "subjectToken",
			Issuer:    "com.lapangan.cuy",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(expAccessToken))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// generate the token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	log.Println("token : ", accessToken)

	// make the accessToken more secure with the jwt_secret
	accessTokenString, err := accessToken.SignedString(os.Getenv("JWT_SECRET"))
	log.Println("accessTokenString : ", accessTokenString)
	if err != nil {
		newError := HandleError(err, "cannot signed jwt access token string")
		return "", "", newError
	}
	// ========= refresh token =======

	expRefreshTokenString := os.Getenv("EXP_REFRESH_TOKEN")
	expRefreshToken, err := strconv.Atoi(expRefreshTokenString)
	if err != nil {
		newError := HandleError(err, "error convert string to int")
		return "", "", newError
	}

	refreshTokenClaims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "subjectToken",
			Issuer:    "com.lapangan.cuy",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * time.Duration(expRefreshToken))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// generate the token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	log.Println("refreshToken : ", refreshToken)

	// make the refreshToken more secure with the jwt_secret
	refreshTokenString, err := refreshToken.SignedString(os.Getenv("JWT_SECRET"))
	log.Println("refreshTokenString : ", refreshTokenString)
	if err != nil {
		newError := HandleError(err, "cannot signed jwt refresh token string")
		return "", "", newError
	}

	return accessTokenString, refreshTokenString, err
}

func RefreshToken(refreshTokenString string) (string, error) {
	// validate refresh token
	claims, err := ValidateToken(refreshTokenString)
	if err != nil {
		newError := HandleError(err, "error validate refresh token")
		return "", newError
	}

	expAccessTokenString := os.Getenv("EXP_ACCESS_TOKEN")
	expAccessToken, err := strconv.Atoi(expAccessTokenString)
	if err != nil {
		newError := HandleError(err, "error convert string to int")
		return "", newError
	}

	// create new access token
	newAccessTokenClaims := Claims{
		UserID: claims.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "subjectToken",
			Issuer:    "com.lapangan.cuy",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(expAccessToken))),
			//ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	// generate the token
	NewAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newAccessTokenClaims)
	log.Println("token : ", NewAccessToken)

	// make the NewAccessToken more secure with the jwt_secret
	NewAccessTokenString, err := NewAccessToken.SignedString(os.Getenv("JWT_SECRET"))
	log.Println("NewAccessTokenString : ", NewAccessTokenString)
	if err != nil {
		newError := HandleError(err, "cannot signed jwt access token string")
		return "", newError
	}

	return NewAccessTokenString, nil
}

func ValidateToken(tokenString string) (*Claims, error) {

	// parse tokenString (with jwt secret) to get the real token (without jwt secret)
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return os.Getenv("JWT_SECRET"), nil
	})

	if err != nil {
		newError := HandleError(err, "error parse token string")
		return nil, newError
	}

	// parse token to Claims struct
	if claims, ok := token.Claims.(Claims); ok && token.Valid {
		// check issuer
		issuer, err := claims.GetIssuer()
		if err != nil {
			newError := HandleError(err, "error - getting claims issuer")

			return nil, newError
		}

		if issuer != "com.lapangan.cuy" {
			return nil, errors.New("error - issuer is not valid")
		}

		// check the expiration of the token
		exp, err := claims.GetExpirationTime()
		if err != nil {
			newError := HandleError(err, "error - getting token expiration time")

			return nil, newError
		}

		if exp.Before(time.Now()) {
			return nil, errors.New("error - token has been expired")
		}

		return &claims, nil
	}

	return nil, errors.New("error - there's something error")

}
