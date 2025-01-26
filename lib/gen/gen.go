// Package gen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package gen

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
	"github.com/oapi-codegen/runtime"
)

const (
	BasicAuthScopes = "basicAuth.Scopes"
)

// Defines values for Brightness.
const (
	Dec Brightness = "dec"
	Inc Brightness = "inc"
)

// Defines values for GroupState.
const (
	AllOn GroupState = "AllOn"
	AnyOn GroupState = "AnyOn"
	None  GroupState = "None"
)

// Defines values for LightMode.
const (
	Day   LightMode = "day"
	Night LightMode = "night"
)

// Brightness defines model for Brightness.
type Brightness string

// GroupState defines model for GroupState.
type GroupState string

// LightGroup defines model for LightGroup.
type LightGroup struct {
	Id     int        `json:"id"`
	Lights []string   `json:"lights"`
	Name   string     `json:"name"`
	State  GroupState `json:"state"`
}

// LightMode defines model for LightMode.
type LightMode string

// LightGroupResponse defines model for LightGroupResponse.
type LightGroupResponse struct {
	Lightgroups *[]LightGroup `json:"lightgroups,omitempty"`
}

// GetBrightnessLightGroupIdParams defines parameters for GetBrightnessLightGroupId.
type GetBrightnessLightGroupIdParams struct {
	Level Brightness `form:"level" json:"level"`
}

// GetModeLightGroupIdParams defines parameters for GetModeLightGroupId.
type GetModeLightGroupIdParams struct {
	Mode LightMode `form:"mode" json:"mode"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /brightness/{lightGroupId})
	GetBrightnessLightGroupId(w http.ResponseWriter, r *http.Request, lightGroupId int, params GetBrightnessLightGroupIdParams)

	// (GET /lightgroups)
	GetLightgroups(w http.ResponseWriter, r *http.Request)

	// (GET /mode/{lightGroupId})
	GetModeLightGroupId(w http.ResponseWriter, r *http.Request, lightGroupId int, params GetModeLightGroupIdParams)

	// (GET /toggle/{lightGroupId})
	GetToggleLightGroupId(w http.ResponseWriter, r *http.Request, lightGroupId int)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetBrightnessLightGroupId operation middleware
func (siw *ServerInterfaceWrapper) GetBrightnessLightGroupId(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "lightGroupId" -------------
	var lightGroupId int

	err = runtime.BindStyledParameterWithOptions("simple", "lightGroupId", mux.Vars(r)["lightGroupId"], &lightGroupId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "lightGroupId", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{})

	r = r.WithContext(ctx)

	// Parameter object where we will unmarshal all parameters from the context
	var params GetBrightnessLightGroupIdParams

	// ------------- Required query parameter "level" -------------

	if paramValue := r.URL.Query().Get("level"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "level"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "level", r.URL.Query(), &params.Level)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "level", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetBrightnessLightGroupId(w, r, lightGroupId, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetLightgroups operation middleware
func (siw *ServerInterfaceWrapper) GetLightgroups(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetLightgroups(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetModeLightGroupId operation middleware
func (siw *ServerInterfaceWrapper) GetModeLightGroupId(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "lightGroupId" -------------
	var lightGroupId int

	err = runtime.BindStyledParameterWithOptions("simple", "lightGroupId", mux.Vars(r)["lightGroupId"], &lightGroupId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "lightGroupId", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{})

	r = r.WithContext(ctx)

	// Parameter object where we will unmarshal all parameters from the context
	var params GetModeLightGroupIdParams

	// ------------- Required query parameter "mode" -------------

	if paramValue := r.URL.Query().Get("mode"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "mode"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "mode", r.URL.Query(), &params.Mode)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "mode", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetModeLightGroupId(w, r, lightGroupId, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetToggleLightGroupId operation middleware
func (siw *ServerInterfaceWrapper) GetToggleLightGroupId(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "lightGroupId" -------------
	var lightGroupId int

	err = runtime.BindStyledParameterWithOptions("simple", "lightGroupId", mux.Vars(r)["lightGroupId"], &lightGroupId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "lightGroupId", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, BasicAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetToggleLightGroupId(w, r, lightGroupId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{})
}

type GorillaServerOptions struct {
	BaseURL          string
	BaseRouter       *mux.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r *mux.Router) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r *mux.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, GorillaServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options GorillaServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = mux.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.HandleFunc(options.BaseURL+"/brightness/{lightGroupId}", wrapper.GetBrightnessLightGroupId).Methods("GET")

	r.HandleFunc(options.BaseURL+"/lightgroups", wrapper.GetLightgroups).Methods("GET")

	r.HandleFunc(options.BaseURL+"/mode/{lightGroupId}", wrapper.GetModeLightGroupId).Methods("GET")

	r.HandleFunc(options.BaseURL+"/toggle/{lightGroupId}", wrapper.GetToggleLightGroupId).Methods("GET")

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8RVTW/jNhD9K8S0RyFy2px0c4DCDeB+oOkih8BY0OJYYkCTCjly4DX03xdDybYU2d7N",
	"BsieJA05j2/mPY52kLt15SxaCpDtwGOonA0YP+a6KGnmXV3914U5mjtLaIlfZVUZnUvSzqZPwVmOhbzE",
	"teS3yrsKPekWzDBYwWDxUxOu48uvHleQwS/pkUjaYoT0SACaBGhbIWQgvZdbaI4Bt3zCnKDhkMKQe10x",
	"I8hgrgMF4VZCbqSWRi4NiogpZi2RJoFPVtZUOq+/oPrDe+eZ1BBmWlOJlrpKxUpqgwoSKFEq9LGKh4eH",
	"z71tOOxERzSQ17Zgoky1K5LXbz2TshjiF9p6DdkjaJsDV5TDInkNkUCs4J66s/Y5U2P+sZDA1G7j829n",
	"8WR6r7UjqbTqkdaWsEDPOVHCoXoj3KFICVi5xpMbw576JQP0iuSWeXyutUfFlUbghLnusQ4EFyNrdPX+",
	"5dSgW0puIQHLSyeaxCQxr72m7T3zaZuzlEHnLPVBYs6JUThAlERVa0htVy7Wr8nwylTJitCLgH6jcxQr",
	"58W/pTZGV0H8WaO49VoVXMwGfWj9d301uZpwEa5CKysNGfweQwlUkspIK10ePJTuzEHcO9XwaoHxvrLI",
	"0cN3CjKYIR2NN++lRFwv10jR3Y870EyDz4K9omCGCUdpyNeYjM1/8FGTdHjPNfptDxA3aC4iXXJK7wo1",
	"zSIZDrLfJpPxrb6v8xxDWNXGbIVUT3UgVOLYxihNb1qwADeT63NEDgem44kSh5UsuJf9i7fgePpqMHZS",
	"DbnOkIQ0RsS9otucjPWc97BO9+Ay9RMT/31V7y9QdFHv6jwuWKSzPVk7hedt/ErHF0152VdKkBOy7ZW2",
	"hWCwU83iafDzbd+x+zHXH6fad5q+7RJniBcZRPyBo3qPyJx5c+6gVg7rSKxcbdV7LEGuKAy+Ybb9HxM+",
	"UOC3T522JjWeMueUG/Xzwy/nN9P4x7Zvb+0NZJDyL6tZNF8DAAD//0SZ2mdsCgAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
