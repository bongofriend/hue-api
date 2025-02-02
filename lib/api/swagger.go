package api

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/bongofriend/hue-api/lib/gen"
	"github.com/bongofriend/hue-api/lib/services"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
	"github.com/swaggest/swgui"
	"github.com/swaggest/swgui/v5emb"
)

func ConfigureDocRouter(router *mux.Router, cfg services.AppConfig) error {
	if !cfg.ShowSwaggerUI {
		log.Println("Swagger UI disabled")
		return nil
	}

	url, err := url.Parse(cfg.BasePath)
	if err != nil {
		panic(err)
	}

	swagger, err := gen.GetSwagger()
	if err != nil {
		return err
	}

	corsHandler := Cors(url.Host)

	docsRouter := router.PathPrefix("/docs").Subrouter()
	docsRouter.Use(mux.MiddlewareFunc(corsHandler))
	docsRouter.Handle("/openapi", serveOpenApiSpec(swagger)).Methods("GET")
	docsRouter.PathPrefix("/swagger").Handler(
		v5emb.NewHandlerWithConfig(swgui.Config{
			Title:            "Hue-API",
			SwaggerJSON:      fmt.Sprintf("%s/docs/openapi", cfg.BasePath),
			BasePath:         fmt.Sprintf("%s/docs/swagger/", cfg.BasePath),
			InternalBasePath: "/docs/swagger/",
		}),
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
