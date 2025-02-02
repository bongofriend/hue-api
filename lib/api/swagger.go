package api

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/bongofriend/hue-api/lib/gen"
	"github.com/bongofriend/hue-api/lib/services"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
	"github.com/swaggest/swgui/v5emb"
)

func ConfigureDocRouter(router *mux.Router, cfg services.AppConfig) error {
	if !cfg.ShowSwaggerUI {
		log.Println("Swagger UI disabled")
		return nil
	}
	swagger, err := gen.GetSwagger()
	if err != nil {
		return err
	}

	docsRouter := router.PathPrefix("/docs").Subrouter()
	docsRouter.Handle("/openapi", serveOpenApiSpec(swagger)).Methods("GET")
	docsRouter.PathPrefix("/swagger").Handler(
		v5emb.New(
			"Hue-Adapter",
			fmt.Sprintf("http://%s:%d/docs/openapi", cfg.Hostname, cfg.Port),
			"/docs/swagger/",
		),
	)

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
