package main

import (
	"context"
	"net/http"

	"github.com/alecthomas/kong"
	"github.com/bongofriend/hue-api/cli/cmd"
	"github.com/bongofriend/hue-api/cli/config"
	"github.com/bongofriend/hue-api/cli/gen"
)

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=openapi_config.yaml ../openapi.yaml

var cli struct {
	Toggle      cmd.ToggleCmd          `cmd:"" help:"Toggle LightGroup"`
	LightGroups cmd.ListLightGroupsCmd `cmd:"" help:"List available LightGroups"`
	LightMode   cmd.SetLightModeCmd    `cmd:"" help:"Set LightMode mode for LightGroup"`
	Brightness  cmd.SetBrightnessCmd   `cmd:"" help:"Increase or decrease brightness of LightGroup"`
}

func authRequestEditor(cfg config.ClientConfig) gen.RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		req.SetBasicAuth(cfg.Username, cfg.Password)
		return nil
	}
}

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	apiClient, err := gen.NewClientWithResponses(cfg.ApiUrl, gen.WithRequestEditorFn(authRequestEditor(cfg)))
	if err != nil {
		panic(err)
	}

	ctx := kong.Parse(&cli)
	err = ctx.Run(cmd.CliContext{Client: apiClient})
	ctx.FatalIfErrorf(err)
}
