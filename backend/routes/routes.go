package routes

import (
	"first_project/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Backend is running"))
	}).Methods("GET")

	// Student routes
	r.HandleFunc("/students", handlers.CreateStudent).Methods("POST")
	r.HandleFunc("/students", handlers.GetStudents).Methods("GET")
	// r.HandleFunc("/students/{id}", handlers.GetStudentByID).Methods("GET")
	r.HandleFunc("/students/{id}", handlers.DeleteStudent).Methods("DELETE")
	r.HandleFunc("/students/{id}", handlers.UpdateStudent).Methods("PUT")

	return r
}
