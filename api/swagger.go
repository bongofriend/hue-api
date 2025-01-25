package api

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/swaggest/swgui/v5emb"
)

func ConfigureDocRouter(router *mux.Router, openApiFilePath string) {
	docsRouter := router.PathPrefix("/docs").Subrouter()
	docsRouter.Handle("/openapi", ServeOpenApiSpec(openApiFilePath)).Methods("GET")
	docsRouter.PathPrefix("/swagger").Handler(v5emb.New("Hue-Adapter", "http://localhost:8080/docs/openapi", "/docs/swagger/"))
}

func ServeOpenApiSpec(openApiFilePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open(openApiFilePath)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if _, err := io.Copy(w, file); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return

		}
	}
}
