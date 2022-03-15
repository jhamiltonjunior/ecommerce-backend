package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Auth struct {
	UserName string `json:"username" db:"username"`
	Password string `json:"passwd" db:"passwd"`
}

// A fake function that returns the entered data
// in the req.Body and a token that is also fake
func (auth *Auth) Authenticate() http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		json.NewDecoder(req.Body).Decode(auth)

		// 	// the return after the error the application continues that prevents
		// 	return
		//

		// I'm putting the "", to overwrite password,
		// and don't display it to the end user
		// please do not use this in frontend application
		auth.Password = ""

		token, err := getJWT()
		if err != nil {
			json.NewEncoder(response).Encode(err)

			return
		}

		json.NewEncoder(response).Encode(auth)
		json.NewEncoder(response).Encode(map[string]string{
			"token": fmt.Sprint(token),
		})
	}

}

// Just like the previous function, this one is no different,
// it will return the entered data
// in the req.Body and a token that is also fake
func (auth *Auth) AuthenticateSSO() http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		json.NewDecoder(req.Body).Decode(auth)


		// I'm putting the "", to overwrite password,
		// and don't display it to the end user
		// please do not use this in frontend application
		auth.Password = ""

		token, err := getJWT()
		if err != nil {
			json.NewEncoder(response).Encode(err)

			return
		}

		json.NewEncoder(response).Encode(auth)
		json.NewEncoder(response).Encode(map[string]string{
			"token": fmt.Sprint(token),
		})
	}

}

func getJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "Krissanawat"
	claims["aud"] = "billing.jwtgo.io"
	claims["iss"] = "jwtgo.io"
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
