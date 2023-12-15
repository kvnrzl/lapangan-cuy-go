package common

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"log"
	"os"
	"time"
)

type Claims struct {
	UserID uuid.UUID
	jwt.RegisteredClaims
}

func GenerateToken(userID uuid.UUID) (string, string, error) {

	// prepare the object Claims
	accessTokenClaims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "subjectToken",
			Issuer:    "com.field.cuy",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// generate the token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	log.Println("token : ", accessToken)

	// make the accessToken more secure with the jwt_secret
	accessTokenString, err := accessToken.SignedString(os.Getenv("JWT_SECRET"))
	log.Println("accessTokenString : ", accessTokenString)
	LogOnError(err, "cannot signed jwt access token string")

	// ========= refresh token =======

	refreshTokenClaims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "subjectToken",
			Issuer:    "com.field.cuy",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
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
		LogOnError(err, "cannot signed jwt refresh token string")
		return "", "", nil
	}

	return accessTokenString, refreshTokenString, err
}

func RefreshToken(refreshTokenString string) (string, error) {
	// validate refresh token
	claims, err := ValidateToken(refreshTokenString)
	LogOnError(err, "error validate refresh token")

	//create new access token
	newAccessTokenClaims := Claims{
		UserID: claims.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "subjectToken",
			Issuer:    "com.field.cuy",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// generate the token
	NewAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newAccessTokenClaims)
	log.Println("token : ", NewAccessToken)

	// make the NewAccessToken more secure with the jwt_secret
	NewAccessTokenString, err := NewAccessToken.SignedString(os.Getenv("JWT_SECRET"))
	log.Println("NewAccessTokenString : ", NewAccessTokenString)
	if err != nil {
		LogOnError(err, "cannot signed jwt access token string")
		return "", err
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

	LogOnError(err, "error parse token string")

	// parse token to Claims struct
	if claims, ok := token.Claims.(Claims); ok && token.Valid {
		return &claims, nil
	}

	return nil, err

}
