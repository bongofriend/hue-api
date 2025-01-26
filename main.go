package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bongofriend/hue-api/api"
	"github.com/bongofriend/hue-api/services"
	"github.com/gorilla/mux"
)

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=openapi_config.yaml ./openapi.yaml

type configFilPath string

// Set implements flag.Value.
func (c *configFilPath) Set(value string) error {
	stat, err := os.Stat(value)
	if err != nil {
		return err
	}
	if !stat.Mode().IsRegular() {
		return errors.New("path does not point to file")
	}

	*c = configFilPath(value)
	return nil
}

// String implements flag.Value.
func (c *configFilPath) String() string {
	if len(*c) == 0 {
		log.Println("No path to config file found. Defaulting to ./config.json")
		*c = "./config.json"
	}
	return string(*c)
}

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
	log.Printf("API server listening on port: %d", cfg.Port)
	http.ListenAndServe(fmt.Sprintf("localhost:%d", cfg.Port), mainRouter)
}

func parsArgs() configFilPath {
	var p configFilPath

	flag.Var(&p, "configFilePath", "Path to config file")
	flag.Parse()

	return p
}
