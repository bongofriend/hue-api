package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/bongofriend/hue-api/lib/api"
	"github.com/bongofriend/hue-api/lib/services"
	"github.com/gorilla/mux"
)

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=openapi_config.yaml ./openapi.yaml

func main() {
	cfgPath := parsArgs()
	configService := services.NewConfigServuce(string(cfgPath))
	cfg, err := configService.GetConfig()
	if err != nil {
		panic(err)
	}

	mainRouter := mux.NewRouter()
	if err := api.ConfigureDocRouter(mainRouter); err != nil {
		panic(err)
	}
	if err := api.ConfigureApiRouter(mainRouter, configService); err != nil {
		panic(err)
	}
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: mainRouter,
	}
	log.Printf("API server listening on port: %d", cfg.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func parsArgs() string {
	var p string

	flag.StringVar(&p, "configFilePath", "Path to config file", "")
	flag.Parse()

	return p
}
