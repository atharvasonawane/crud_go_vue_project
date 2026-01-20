package route

import (
	"net/http"
	"reporting-utility/internal/handler"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()
	
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Reporting Utility & Backend Running"))
	})

	// Original API Routes (CRUD)
	r.HandleFunc("/students", handler.CreateStudent).Methods("POST")
	r.HandleFunc("/students", handler.GetStudents).Methods("GET")
	r.HandleFunc("/students", handler.UpdateStudent).Methods("PUT")
	r.HandleFunc("/students", handler.DeleteStudent).Methods("DELETE")
	
	// Session Routes
	r.HandleFunc("/select-student", handler.SelectStudent).Methods("POST")
	r.HandleFunc("/student-detail", handler.GetSelectedStudent).Methods("GET")

	// PDF Route (Mapped to old path /students/pdf)
	r.HandleFunc("/students/pdf", handler.DownloadPDF).Methods("GET")

	// File Server for Uploaded Images
	r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))

	return r
}