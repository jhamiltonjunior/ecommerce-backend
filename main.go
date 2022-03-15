/**
	A simple API  for Backend of E-commerce

	This file is part of undefined is not a function API.
	You are free to modify and share this project or its files.
	
	Author   Hamilton Júnior <https://github.com/jhamiltonjunior>

	Contact <https://www.linkedin.com/in/jhamiltonjunior>
	Contact <mailto:josehamiltonsantosjunior@gmail.com>
*/
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jhamiltonjunior/ecommerce-backend/src/routes"
	"github.com/joho/godotenv"
)

var (
	// godotenv vai fazer o arquivo .env funcionar
	// sem ele o os.Getenv não funciona
	err = godotenv.Load(".env")

	// os.Getenv irá buscar o HTTP_PORT dentro do arquivo .env
	port = os.Getenv("HTTP_PORT")
)

func main() {
	// Essa captura de erro é referente ao godotenv
	if err != nil {
		fmt.Println(err)
	}

	routes.NewServer()

	fmt.Printf("server listen in port%s \n", port)

	// Você irá encontrar o valor de port no seu arquivo .env ou .env.exemple
	http.ListenAndServe(port, nil)
}
