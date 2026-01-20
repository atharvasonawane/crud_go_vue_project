package main

import (
	"fmt"
	"log"
	"net/http"
	"reporting-utility/internal/db"
	"reporting-utility/internal/route"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// 1. Initialize DB & Session
	dsn := "root:mysql@atharva04@tcp(localhost:3306)/student_db?parseTime=true"
	db.ConnectDB(dsn)
	db.InitSession()

	// 2. Setup Router
	r := route.InitRoutes()

	// 3. Start Server with CORS
	fmt.Println("Server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", enableCORS(r)))
}