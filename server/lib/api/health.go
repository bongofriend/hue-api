package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ConfigureHealthRouter(mux *mux.Router) {
	mux.HandleFunc("/", health).Methods("GET")
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
