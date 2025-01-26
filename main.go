package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/bongofriend/hue-api/api"
	"github.com/bongofriend/hue-api/config"
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
	return string(*c)
}

func main() {
	//cfgPath := parsArgs()
	cfg, err := config.GetConfig("config.json")
	if err != nil {
		panic(err)
	}

	mainRouter := mux.NewRouter()
	if err := api.ConfigureDocRouter(mainRouter); err != nil {
		panic(err)
	}
	if err := api.ConfigureApiRouter(mainRouter, cfg); err != nil {
		panic(err)
	}
	http.ListenAndServe(fmt.Sprintf("localhost:%d", cfg.Port), mainRouter)
}

func parsArgs() configFilPath {
	var p configFilPath

	flag.Var(&p, "configFilePath", "Path to config file")
	flag.Parse()

	return p
}
