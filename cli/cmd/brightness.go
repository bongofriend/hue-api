package cmd

import (
	"context"
	"fmt"

	"github.com/bongofriend/hue-api/cli/gen"
)

type SetBrightnessCmd struct {
	LightGroupId    int    `arg:"" name:"lightGroupId" help:"Id of LightGroup to increase or decrease brightness" type:"int"`
	BrightnessLevel string `arg:"" name:"brightnessLevel" help:"Increase or decrease brightness of LightGroup" enum:"inc,dec" type:"string"`
}

func (s *SetBrightnessCmd) Run(ctx CliContext) error {
	var level gen.Brightness
	switch s.BrightnessLevel {
	case "inc":
		level = gen.Inc
		break
	case "dec":
		level = gen.Dec
		break
	default:
		return fmt.Errorf("Adjust brightness level LightGroup: unknown brightness adjustment: %s", s.BrightnessLevel)
	}

	if _, err := ctx.Client.GetBrightnessLightGroupId(
		context.Background(),
		s.LightGroupId,
		&gen.GetBrightnessLightGroupIdParams{Level: level},
	); err != nil {
		return fmt.Errorf("Adjust brightness level LightGroup: %w", err)
	}
	return nil
}
