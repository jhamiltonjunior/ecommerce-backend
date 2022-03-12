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

// The User struct is responsible for getting the req.Body and inserting it into the database
// and the same is responsible for "porpulating" the JSON that returns from the database
//
// Please you from the frontend, redirect the user to the route
//  /api/v{n}/authenticate
// Here I just create the user, I don't have any JWT authenticate here
type User struct {
	ID uint `json:"user_id" gorm:"primaryKey; autoIncrement"`

	// I put Name, because if I put UserName when going to use
	// would have to call user.UserName and I don't like that
	// user.Name is already implied
	//  used for get in user account
	// Login string `json:"login" gorm:"unique; not null"`
	// this is full name of people
	// not is used for get in
	FullName string `json:"full_name"`
	// also used for get in user account
	Email    string `json:"email" gorm:"type: varchar(100); unique; not null"`
	Password string `json:"password"`

	// the timestamp
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) CreateUser() http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		json.NewDecoder(req.Body).Decode(user)

		db, err := config.OpenDB()
		if err != nil {
			json.NewEncoder(response).Encode(map[string]string{
				"message": fmt.Sprint(err),
			})

			// the return after the error the application continues that prevents
			return
		}

		result := db.Create(&user)

		fmt.Println("\n", result.Error)
		
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

// ListAllUser will list ALL users that are registered in the database
// that's right, if there are 1 thousand users I advise you to put another 8 or 16 GB's of RAM
// on your machine
//
// In a new feature, a DESC LIMIT {NUMBER} OFFSET {NUMBER} could be placed in the query
// This will prevent server crashes or slowdowns.
//
// Can you imagine having to list 1000 users for 20 people at the same time?
//
// Consider being very careful with this.
//
// Why wasn't this implemented?
//
// This function shouldn't even exist!
// I ended up creating this function by accident, now it's a feature
func (user *User) ListAllUser() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		json.NewEncoder(response).Encode("user")
	}
}

// ListUniqueUser Wil list a single user by id of url
//  /api/v{1}/user/{id:[0-9]+}
// If there is no error it will return a JSON with the referring user
// to the id of the url
func (user *User) ListUniqueUser() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)

		fmt.Println(params)

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
