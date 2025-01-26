package api

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/bongofriend/hue-api/lib/gen"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
	"github.com/swaggest/swgui/v5emb"
)

func ConfigureDocRouter(router *mux.Router) error {
	swagger, err := gen.GetSwagger()
	if err != nil {
		return err
	}

	docsRouter := router.PathPrefix("/docs").Subrouter()
	docsRouter.Handle("/openapi", serveOpenApiSpec(swagger)).Methods("GET")
	docsRouter.PathPrefix("/swagger").Handler(v5emb.New("Hue-Adapter", "http://localhost:8080/docs/openapi", "/docs/swagger/"))

	return nil
}

func serveOpenApiSpec(swaggerSpec *openapi3.T) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		specBuf, _ := swaggerSpec.MarshalJSON()

		if _, err := io.Copy(w, bytes.NewReader(specBuf)); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return

		}
	}
}
