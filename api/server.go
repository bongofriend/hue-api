package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bongofriend/hue-api/config"
	"github.com/bongofriend/hue-api/gen"
	"github.com/bongofriend/hue-api/services"
	"github.com/gorilla/mux"
)

var _ gen.ServerInterface = (*Server)(nil)

func ConfigureApiRouter(mainRouter *mux.Router, cfg config.AppConfig) error {
	hueService, err := services.NewHueService(cfg.Hue)
	if err != nil {
		return fmt.Errorf("could not start HueService: %w", err)
	}
	server := newServer(hueService)
	swagger, err := gen.GetSwagger()

	if err != nil {
		return err
	}

	gen.HandlerWithOptions(&server, gen.GorillaServerOptions{
		BaseURL:     "/api",
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
	panic("unimplemented")
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
	if err := encoder.Encode(rsp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Panicln(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetModeLightGroupId implements gen.ServerInterface.
func (s *Server) GetModeLightGroupId(w http.ResponseWriter, r *http.Request, lightGroupId int, params gen.GetModeLightGroupIdParams) {
	panic("unimplemented")
}

// TODO: Test implementation later (429 by Hue API)
// GetToggleLightGroupId implements gen.ServerInterface.
func (s *Server) GetToggleLightGroupId(w http.ResponseWriter, r *http.Request, lightGroupId int) {
	err := s.hueService.ToggleLightGroup(r.Context(), lightGroupId)
	if err != nil {
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
