package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	ID interface{}
	jwt.StandardClaims
}

var SECRETE_KEY = os.Getenv("SECRETE_KEY")

func GenerateTokenFrom(userId interface{}) string {
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


func GetUserId(token string) {
	// claims := &Claims{}
	// jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {

	// })
}