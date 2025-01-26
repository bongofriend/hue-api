package services

import (
	"context"
	"fmt"

	"github.com/amimof/huego"
	"github.com/bongofriend/hue-api/config"
	"github.com/bongofriend/hue-api/gen"
)

type HueService interface {
	GetLightGroups(ctx context.Context) ([]gen.LightGroup, error)
	ToggleLightGroup(ctx context.Context, lightGroupId int) error
}

type hueService struct {
	bridge *huego.Bridge
}

func NewHueService(hueConfig config.HueConfig) (HueService, error) {
	bridge, err := huego.Discover()
	if err != nil {
		return nil, err
	}
	bridge = bridge.Login(hueConfig.Username)
	return hueService{
		bridge: bridge,
	}, nil
}

// GetLightGroups implements HueService.
func (h hueService) GetLightGroups(ctx context.Context) ([]gen.LightGroup, error) {
	groups, err := h.bridge.GetGroupsContext(ctx)
	if err != nil {
		return nil, err
	}

	lightGroups := make([]gen.LightGroup, len(groups))
	for idx, g := range groups {
		lightGroups[idx] = mapToLightGroup(g)
	}
	return lightGroups, nil
}

// ToggleLightGroup implements HueService.
func (h hueService) ToggleLightGroup(ctx context.Context, lightGrouoId int) error {
	group, err := h.bridge.GetGroupContext(ctx, lightGrouoId)
	if err != nil {
		return fmt.Errorf("could not toggle Light Group: %w", err)
	}
	if group.GroupState.AllOn || group.GroupState.AnyOn {
		group.State.On = false
	} else {
		group.State.On = true
	}
	group.State.On = !group.GroupState.AllOn && !group.GroupState.AnyOn
	_, err = h.bridge.SetGroupState(lightGrouoId, *group.State)
	return err
}

func mapToLightGroup(g huego.Group) gen.LightGroup {
	lightGroup := gen.LightGroup{
		Lights: make([]string, len(g.Lights)),
	}
	lightGroup.Id = g.ID
	lightGroup.Name = g.Name

	if g.GroupState.AllOn {
		lightGroup.State = gen.AllOn
	} else if g.GroupState.AnyOn {
		lightGroup.State = gen.AnyOn
	} else {
		lightGroup.State = gen.None
	}

	for idx, l := range g.Lights {
		lightGroup.Lights[idx] = l
	}
	return lightGroup
}
