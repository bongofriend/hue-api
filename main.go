package main

import (
	"net/http"

	"github.com/bongofriend/hue-api/api"
	"github.com/gorilla/mux"
)

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=openapi_config.yaml ./openapi.yaml

func main() {
	mainRouter := mux.NewRouter()
	api.ConfigureDocRouter(mainRouter, "./openapi.yaml")
	api.ConfigureApiRouter(mainRouter)
	http.ListenAndServe("localhost:8080", mainRouter)
}
