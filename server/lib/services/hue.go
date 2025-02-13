package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/amimof/huego"
	"github.com/bongofriend/hue-api/server/lib/gen"
)

const (
	brigtnessStep       uint8 = 20
	dayModeBrightness         = 255
	nightModeBrightness       = 10
	maxBrightness             = 255
	minBrightness             = 20
)

type HueService interface {
	GetLightGroups(ctx context.Context) ([]gen.LightGroup, error)
	ToggleLightGroup(ctx context.Context, lightGroupId int) error
	UpdateLightMode(ctx context.Context, lightGroupId int, mode gen.LightMode) error
	AdjustLightGroupBrigtness(ctx context.Context, lightGroupId int, brightnessLevel gen.Brightness) error
}

type hueService struct {
	bridge *huego.Bridge
}

func NewHueService(configService ConfigService) (HueService, error) {
	bridge, err := huego.Discover()
	if err != nil {
		return nil, fmt.Errorf("hue brigde discovery: %w", err)
	}
	if err := handleBridgeLogin(bridge, configService); err != nil {
		return nil, fmt.Errorf("hue bridge login: %w", err)
	}
	return hueService{
		bridge: bridge,
	}, nil
}

func handleBridgeLogin(bridge *huego.Bridge, configService ConfigService) error {
	config, _ := configService.GetConfig()
	if config.Hue.Username == nil {
		log.Printf("No username found in config. Attepting to create new user on Hue Bridge %s. Waiting for 15 seconds", bridge.ID)
		time.Sleep(15 * time.Second)
		usereName, err := bridge.CreateUser(config.Hue.PlainUsername)
		if err != nil {
			return fmt.Errorf("could not create user on Bridge %s: %s", bridge.ID, config.Hue.PlainUsername)
		}
		config.Hue.Username = &usereName
		if err := configService.UpdateConfig(config); err != nil {
			return err
		}
	} else {
		bridge = bridge.Login(*config.Hue.Username)
	}
	return nil
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
func (h hueService) ToggleLightGroup(ctx context.Context, lightGroupId int) error {
	group, err := h.bridge.GetGroupContext(ctx, lightGroupId)
	if err != nil {
		return fmt.Errorf("could not toggle Light Group: %w", err)
	}
	if group.IsOn() {
		return group.OffContext(ctx)
	}
	return group.OnContext(ctx)
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

// UpdateLightMode implements HueService.
func (h hueService) UpdateLightMode(ctx context.Context, lightGroupId int, mode gen.LightMode) error {
	g, err := h.bridge.GetGroupContext(ctx, lightGroupId)
	if err != nil {
		return fmt.Errorf("could not set LightMode %s for LightGroup %d: %w", mode, lightGroupId, err)
	}
	switch mode {
	case gen.Day:
		return g.BriContext(ctx, dayModeBrightness)
	case gen.Night:
		return g.BriContext(ctx, nightModeBrightness)
	}
	return fmt.Errorf("unsupported LightMode %s", mode)
}

// AdjustLightGroupBrigtness implements HueService.
func (h hueService) AdjustLightGroupBrigtness(ctx context.Context, lightGroupId int, brightnessLevel gen.Brightness) error {
	g, err := h.bridge.GetGroupContext(ctx, lightGroupId)
	if err != nil {
		return fmt.Errorf("could not adjust brightness for LightGroup %d: %w", lightGroupId, err)
	}
	bri := g.State.Bri
	if brightnessLevel == gen.Inc {
		if bri == maxBrightness {
			log.Printf("Max brightness level for LightGroup %d reached. Disgarding adjustment", lightGroupId)
			return nil
		}
		bri += brigtnessStep
	} else if brightnessLevel == gen.Dec {
		if bri == minBrightness {
			log.Printf("Min brightness level for LightGroup %d reached. Disgarding adjustment", lightGroupId)
			return nil
		}
		bri -= brigtnessStep
	}
	return g.BriContext(ctx, bri)
}
