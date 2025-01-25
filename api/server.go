package api

import (
	"encoding/json"
	"net/http"

	"github.com/bongofriend/hue-api/gen"
	"github.com/gorilla/mux"
)

var _ gen.ServerInterface = (*Server)(nil)

func ConfigureApiRouter(mainRouter *mux.Router) {
	server := newServer()
	gen.HandlerFromMuxWithBaseURL(&server, mainRouter, "/api")
}

type Server struct{}

// GetLightgroups implements gen.ServerInterface.
func (s *Server) GetLightgroups(w http.ResponseWriter, r *http.Request) {
	b, _ := json.Marshal(gen.LightGroupResponse{})
	w.Write(b)
}

// GetModeLightGroupId implements gen.ServerInterface.
func (s *Server) GetModeLightGroupId(w http.ResponseWriter, r *http.Request, lightGroupId int, params gen.GetModeLightGroupIdParams) {
	panic("unimplemented")
}

// GetToggleLightGroupId implements gen.ServerInterface.
func (s *Server) GetToggleLightGroupId(w http.ResponseWriter, r *http.Request, lightGroupId int) {
	panic("unimplemented")
}

func newServer() Server {
	return Server{}
}
