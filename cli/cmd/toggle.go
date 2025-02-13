package cmd

import (
	"context"
	"fmt"
)

type ToggleCmd struct {
	LightGroupId int `arg:"" name:"lightGroupId" help:"Id of LightGroup to toggle" type:"int"`
}

func (t *ToggleCmd) Run(ctx CliContext) error {
	_, err := ctx.Client.GetToggleLightGroupId(context.Background(), t.LightGroupId)
	if err != nil {
		return fmt.Errorf("Toggle LightGroup: %w", err)
	}
	return nil
}
