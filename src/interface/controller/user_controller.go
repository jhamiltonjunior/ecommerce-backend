/**
 * Created by GoLang.
 * User: josehamiltonsantosjunior@gmail.com
 * Date: 2018-03-10
 * Time: 17:01
 *
 *
 * You are free to modify and share this project or its files.
 *
 */

package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jhamiltonjunior/blog-backend/src/config"
)

var (
// errValueNotExist = fmt.Errorf("record not found")
)

// The User struct is responsible for getting the req.Body and inserting it into the database
// and the same is responsible for "porpulating" the JSON that returns from the database
//
// Please you from the frontend, redirect the user to the route
//  /api/v{n}/authenticate
// Here I just create the user, I don't have any JWT authenticate here
type User struct {
	ID int `json:"user_id" gorm:"primaryKey; autoIncrement; not null"`

	// I put Name, because if I put UserName when going to use
	// would have to call user.UserName and I don't like that
	// user.Name is already implied
	//  used for get in user account
	// UserName string `json:"username" gorm:"unique; not null"`

	// this is full name of people
	// not is used for get in
	FullName string `json:"full_name"`
	// also used for get in user account
	Email    string `json:"email" gorm:"type: varchar(100); unique; not null"`
	Password string `json:"password" gorm:"not null"`

	// the timestamp
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) CreateUser() http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		json.NewDecoder(req.Body).Decode(user)

		db, err := config.OpenDB()
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(response).Encode(map[string]string{
				"message": fmt.Sprint(err),
			})

			// the return after the error the application continues that prevents
			return
		}

		result := db.Create(&user)

		emailIsDuplicate := IsDuplicate("users_email_key")
		message := result.Error.Error()

		if message == emailIsDuplicate {
			response.WriteHeader(http.StatusBadRequest)

			json.NewEncoder(response).Encode(map[string]string{
				"message": "This email already exist, try another!",
			})

			return
		}

		// I'm putting the "", to overwrite password,
		// and don't display it to the end user
		// please do not use this in frontend application
		user.Password = ""

		response.WriteHeader(http.StatusCreated)
		json.NewEncoder(response).Encode(&user)
	}

}

// ShowUser Wil list a single user by id of url
//  /api/v{1}/user/{id:[0-9]+}
// If there is no error it will return a JSON with the referring user
// to the id of the url
func (user User) ShowUser() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)

		db, err := config.OpenDB()
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(response).Encode(map[string]string{
				"message": fmt.Sprint(err),
			})

			return
		}

		// result := db.Last(&user, "id = ?", params["id"])
		result := db.Where("id = ?", params["id"]).Or("id = ?", params["id"]).Find(&user)

		fmt.Println(result.RowsAffected)
		rows := result.RowsAffected

		user.Password = ""

		// message := result.Error
		if rows == 0 {
			response.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(response).Encode(map[string]string{
				"message": "This user not exist!",
			})
			return
		}

		json.NewEncoder(response).Encode(user)
	}
}

// This function will update the user data
// I was using insomnia and when I updated user data 1
// it was no longer listed at the beginning of the function
//
//  The last user to be modified goes to the end of ListAll()
//  At least that is how it was for me using *Insomnia*
func (user *User) UpdateUser() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		fmt.Println(params)

		json.NewDecoder(request.Body).Decode(user)

		user.UpdatedAt = time.Now()

		json.NewEncoder(response).Encode(user)
	}
}

// Will delete a user by id
func (user *User) DeleteUser() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)
		fmt.Println(params)

		json.NewEncoder(response).Encode(map[string]string{
			"message": "User deleted with success!",
		})
	}
}
