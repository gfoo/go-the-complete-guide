package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secret"

func GenerateToken(email string, userId int64) (string, error) {
	return jwt.NewWithClaims(
		jwt.SigningMethodHS256, jwt.MapClaims{
			"email":  email,
			"userId": userId,
			"exp":    time.Now().Add(time.Hour * 2).Unix(),
		}).SignedString([]byte(secretKey))
}
