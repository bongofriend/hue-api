package cmd

import (
	"context"
	"fmt"

	"github.com/bongofriend/hue-api/cli/gen"
)

type SetLightModeCmd struct {
	LightGroupId int    `arg:"" name:"lightGroupId" help:"Id of LightGroup to set mode" type:"int"`
	Mode         string `arg:"" name:"mode" help:"Set Light mode for LightGroup" enum:"day,night"`
}

func (s *SetLightModeCmd) Run(ctx CliContext) error {
	var mode gen.LightMode
	switch s.Mode {
	case "day":
		mode = gen.Day
		break
	case "night":
		mode = gen.Night
		break
	default:
		return fmt.Errorf("Set Light mode: unkonwn mode: %s", s.Mode)
	}

	if _, err := ctx.Client.GetModeLightGroupId(
		context.Background(),
		s.LightGroupId,
		&gen.GetModeLightGroupIdParams{Mode: mode},
	); err != nil {
		return fmt.Errorf("Set Light Mode: %w", err)
	}

	return nil
}
