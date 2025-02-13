package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/bongofriend/hue-api/cli/gen"
	"github.com/jedib0t/go-pretty/v6/table"
)

type ListLightGroupsCmd struct {
}

func (l *ListLightGroupsCmd) Run(ctx CliContext) error {
	rsp, err := ctx.Client.GetLightgroupsWithResponse(context.Background())
	if err != nil {
		return fmt.Errorf("List LightGroups: %w", err)
	}
	l.writeApiResponse(rsp.JSON200.Lightgroups)
	return nil
}

func (l *ListLightGroupsCmd) writeApiResponse(lightGroups *[]gen.LightGroup) {
	writer := table.NewWriter()
	writer.SetOutputMirror(os.Stdout)

	writer.AppendHeader(table.Row{"Id", "Lights", "Name", "State"})
	for _, g := range *lightGroups {
		lightIds := fmt.Sprintf("[%s]", strings.Join(g.Lights, ", "))
		writer.AppendRow(table.Row{g.Id, lightIds, g.Name, g.State})
		writer.AppendSeparator()
	}

	writer.Render()
}
