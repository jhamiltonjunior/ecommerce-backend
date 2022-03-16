package routes

import (
	"github.com/jhamiltonjunior/ecommerce-backend/src/controllers"
)

// This function will manage user routes
//
// Note that User is part of Server struct but there is no struct here
// this happens because User is part of package router
// and struct Server is part of package router
//
//  User() is package router and Server struct too
//
// Note que Usuário faz parte de Server struct, mas não tem nenhum struct aqui
// isso acontece pelo fato de que o User faz parte do packege router
// e a struct Serve também faz parte do package router
//
func (server *Server) User() {

	// user := entities.User{}

	server.HandleFunc("/api/v1/user", controllers.CreateUser()).Methods("POST")

	server.HandleFunc("/api/v1/user/{id:[0-9]+}", controllers.ShowUser()).Methods("GET")
	// server.HandleFunc("/api/v1/user/{id:[0-9]+}", controllers.DeleteUser()).Methods("DELETE")

	server.HandleFunc("/api/v1/user/{id:[0-9]+}/edit", controllers.UpdateUser()).Methods("PUT")
}
