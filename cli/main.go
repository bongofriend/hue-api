package main

import (
	"fmt"

	"github.com/alecthomas/kong"
)

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=openapi_config.yaml ../openapi.yaml

type baseCmd struct {
	ConfigFilePath string `arg:"" name:"config" help:"Path to config file" type:"path"`
}

type ToggleCmd struct {
	baseCmd
	LightGroupId int `arg:"" name:"lightGroupId" help:"Id of LightGroup to toggle" type:"int"`
}

// TODO: Add openapi client
func (t *ToggleCmd) Run() error {
	fmt.Println("Toggle")
	return nil
}

var cli struct {
	Toggle ToggleCmd `cmd:"" help:"Toogle LightGroup"`
}

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
