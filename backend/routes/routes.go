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
	r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))
	r.HandleFunc("/select-student", handlers.SelectStudent).Methods("POST")
	r.HandleFunc("/student-detail", handlers.GetSelectedStudent).Methods("GET")
	r.HandleFunc("/students", handlers.UpdateStudent).Methods("PUT")
	r.HandleFunc("/students", handlers.DeleteStudent).Methods("DELETE")
	r.HandleFunc("/students/pdf", handlers.DownloadStudentsPDF).Methods("GET")

	return r
}
