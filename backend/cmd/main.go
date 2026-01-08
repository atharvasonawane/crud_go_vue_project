package main

import (
	"first_project/config"
	"first_project/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {

	fmt.Println("GO BACKEND")

	config.ConnectDB()

	r := routes.RegisterRoutes()
	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
