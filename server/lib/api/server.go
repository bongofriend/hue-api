package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bongofriend/hue-api/server/lib/gen"
	"github.com/bongofriend/hue-api/server/lib/services"
	"github.com/gorilla/mux"
)

const apiPathPrefix string = "/api"

var _ gen.ServerInterface = (*Server)(nil)

func ConfigureApiRouter(mainRouter *mux.Router, configService services.ConfigService) error {
	cfg, _ := configService.GetConfig()
	hueService, err := services.NewHueService(configService)
	if err != nil {
		return fmt.Errorf("could not start HueService: %w", err)
	}
	server := newServer(hueService)
	swagger, err := gen.GetSwagger()

	if err != nil {
		return err
	}

	gen.HandlerWithOptions(&server, gen.GorillaServerOptions{
		BaseURL:     apiPathPrefix,
		Middlewares: []gen.MiddlewareFunc{loggerMiddleware, openApiValidatorMiddleware(swagger), authMiddleWare(newBasicAuthenticator(cfg.Auth))},
		BaseRouter:  mainRouter,
	})

	return nil
}

type Server struct {
	hueService services.HueService
}

// GetBrightnessLightGroupId implements gen.ServerInterface.
func (s *Server) GetBrightnessLightGroupId(w http.ResponseWriter, r *http.Request, lightGroupId int, params gen.GetBrightnessLightGroupIdParams) {
	if err := s.hueService.AdjustLightGroupBrigtness(r.Context(), lightGroupId, params.Level); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}

// GetLightgroups implements gen.ServerInterface.
func (s *Server) GetLightgroups(w http.ResponseWriter, r *http.Request) {
	lightGroups, err := s.hueService.GetLightGroups(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	encoder := json.NewEncoder(w)
	rsp := gen.LightGroupResponse{
		Lightgroups: &lightGroups,
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	if err := encoder.Encode(rsp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Panicln(err)
		return
	}
}

// GetModeLightGroupId implements gen.ServerInterface.
func (s *Server) GetModeLightGroupId(w http.ResponseWriter, r *http.Request, lightGroupId int, params gen.GetModeLightGroupIdParams) {
	if err := s.hueService.UpdateLightMode(r.Context(), lightGroupId, params.Mode); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetToggleLightGroupId implements gen.ServerInterface.
func (s *Server) GetToggleLightGroupId(w http.ResponseWriter, r *http.Request, lightGroupId int) {
	if err := s.hueService.ToggleLightGroup(r.Context(), lightGroupId); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func newServer(hueService services.HueService) Server {
	return Server{
		hueService: hueService,
	}
}
