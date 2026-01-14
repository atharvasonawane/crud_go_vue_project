package main

import (
	"first_project/config"
	"first_project/routes"
	"fmt"
	"log"
	"net/http"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	fmt.Println("GO BACKEND")

	config.ConnectDB()
	config.InitSession()

	r := routes.RegisterRoutes()
	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", enableCORS(r)))

}
