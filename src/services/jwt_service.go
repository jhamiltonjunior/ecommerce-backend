package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	ID int
	jwt.StandardClaims
}

const SECRETE_KEY = "marakd"

func GenerateToken(userId int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		ID: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(time.Hour * 3),
			Issuer:    "test",
		},
	})

	signedString, err := token.SignedString([]byte(SECRETE_KEY))

	fmt.Printf("%v%v \n\n", signedString, err)

	return signedString
}