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

package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	// "github.com/gorilla/mux"
	"github.com/gorilla/mux"
	"github.com/jhamiltonjunior/ecommerce-backend/src/configs"
	"github.com/jhamiltonjunior/ecommerce-backend/src/entities"
	"github.com/jhamiltonjunior/ecommerce-backend/src/repositories"
	"github.com/jhamiltonjunior/ecommerce-backend/src/services"
	"golang.org/x/crypto/bcrypt"
)

var (

// userNotPassword *entities.UserWithoutPassword

// errValueNotExist = fmt.Errorf("record not found")
// emailIsDuplicate = `ERROR: duplicate key value violates unique constraint "users_email_key" (SQLSTATE 23505)`
)

// The User struct is responsible for getting the req.Body and inserting it into the database
// and the same is responsible for "porpulating" the JSON that returns from the database
//
// Please you from the frontend, redirect the user to the route
//  /api/v{n}/authenticate
// Here I just create the user, I don't have any JWT authenticate here
func CreateUser() http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		var user *entities.User

		json.NewDecoder(req.Body).Decode(&user)

		// se o bcrypt falhar o usuário não poderá ser criado de forma alguma
		// pois é provavél que a senha vai está com valor null, ainda que passe daqui,
		// já está setado no banco de dados para que a senha seja not null
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(response).Encode(map[string]string{
				"message": fmt.Sprint(err, "\n"),
				"Bcrypt":  "Hash not created",
			})

			// the return after the error the application continues that prevents
			return
		}

		repos := repositories.New(repositories.Options{
			ReaderSqlx: configs.GetReaderSqlx(),
			WriterSqlx: configs.GetWriterSqlx(),
		})

		id, err := repos.User.Create(context.Background(), entities.User{
			FullName: user.FullName,
			Email:    user.Email,
			Password: hash,
		})
		if err != nil {
			fmt.Println(err)

			response.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(response).Encode(map[string]string{
				"message": fmt.Sprint(err),
			})

			return
		}

		// if result := db.Create(&user); result.Error != nil {
		// 	emailIsDuplicate := IsDuplicate("users_email_key")
		// 	message := result.Error.Error()

		// 	fmt.Println(result.Error.Error())

		// 	fmt.Println(result.RowsAffected)

		// 	if message == emailIsDuplicate {
		// 		response.WriteHeader(http.StatusBadRequest)

		//

		// 		return
		// 	}

		// 	response.WriteHeader(http.StatusInternalServerError)

		// 	// Se a condição anterior for false esse json vai ser enviado
		// 	// se for true esse json não será enviado, por causo do return no fim do if
		// 	json.NewEncoder(response).Encode(map[string]string{
		// 		"message": result.Error.Error(),
		// 	})

		// 	return
		// }

		token := services.GenerateToken(id)

		response.WriteHeader(http.StatusCreated)
		json.NewEncoder(response).Encode(map[string]string{
			"Success": fmt.Sprintf("user: %v, created with success!", user.FullName),
			"token": token,
		})
	}

}

// ShowUser Wil list a single user by id of url
//  /api/v{1}/user/{id:[0-9]+}
// If there is no error it will return a JSON with the referring user
// to the id of the url
func ShowUser() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)

		repos := repositories.New(repositories.Options{
			ReaderSqlx: configs.GetReaderSqlx(),
			WriterSqlx: configs.GetWriterSqlx(),
		})

		id, err := strconv.Atoi(params["id"])
		if err != nil {
			response.WriteHeader(http.StatusBadRequest)

			json.NewEncoder(response).Encode(map[string]string{
				"message": "Isso reamente é um id?",
			})

			return
		}

		user, err := repos.User.GetById(context.Background(), id)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(response).Encode(map[string]string{
				"message": fmt.Sprint(err),
			})

			return
		}

		// user.Password = []byte("")
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(user)
	}
}

// This function will update the user data
// I was using insomnia and when I updated user data 1
// it was no longer listed at the beginning of the function
//
//  The last user to be modified goes to the end of ListAll()
//  At least that is how it was for me using *Insomnia*
func UpdateUser() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		var user entities.User

		json.NewDecoder(request.Body).Decode(&user)

		params := mux.Vars(request)

		repos := repositories.New(repositories.Options{
			ReaderSqlx: configs.GetReaderSqlx(),
			WriterSqlx: configs.GetWriterSqlx(),
		})

		id, err := strconv.Atoi(params["id"])
		if err != nil {
			response.WriteHeader(http.StatusBadRequest)

			json.NewEncoder(response).Encode(map[string]string{
				"message": "My bad!",
			})

			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(response).Encode(map[string]string{
				"message": fmt.Sprint(err, "\n"),
				"Bcrypt":  "Hash not created",
			})

			// the return after the error the application continues that prevents
			return
		}

		user.Password = hash
		user.UpdatedAt = time.Now()

		err = repos.User.UpdateById(context.Background(), id, user)
		if err != nil {
			fmt.Println(err)
			response.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(response).Encode(map[string]string{
				"message": fmt.Sprint(err),
			})

			return
		}

		json.NewEncoder(response).Encode(user)
	}
}


// Will delete a user by id
func DeleteUser() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		params := mux.Vars(request)

		id, err := strconv.Atoi(params["id"])
		if err != nil {
			response.WriteHeader(http.StatusBadRequest)

			json.NewEncoder(response).Encode(map[string]string{
				"message": "My bad!",
			})

			return
		}

		repos := repositories.New(repositories.Options{
			ReaderSqlx: configs.GetReaderSqlx(),
			WriterSqlx: configs.GetWriterSqlx(),
		})

		err = repos.User.DeleteById(id)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)

			json.NewEncoder(response).Encode(map[string]string{
				"message": fmt.Sprint(err),
			})
		}



		json.NewEncoder(response).Encode(map[string]string{
			"message": "User deleted with success!",
		})
	}
}

