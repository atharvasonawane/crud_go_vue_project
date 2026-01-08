package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Backend is running"))
	}).Methods("GET")
	return r
}
