package api

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/bongofriend/hue-api/server/lib/gen"
	"github.com/bongofriend/hue-api/server/lib/services"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	middleware "github.com/oapi-codegen/nethttp-middleware"
)

const (
	authKey string = "AuthKey"
)

type authValidator interface {
	validate(authorizationHeaderValue string) (basicAuthUser, error)
}

type basicAuthUser struct {
	Username string
	Password string
}

type basicAuthValidator struct {
	user basicAuthUser
}

func newBasicAuthenticator(authConfig services.AuthConfig) authValidator {
	return basicAuthValidator{
		user: basicAuthUser{
			Username: authConfig.Username,
			Password: authConfig.Password,
		},
	}
}

func (b basicAuthValidator) validate(authHeaderValue string) (basicAuthUser, error) {
	if !strings.HasPrefix(authHeaderValue, "Basic ") {
		return basicAuthUser{}, errors.New("authorization Header has not Basic prefix")
	}

	encodedHeaderContent := strings.TrimPrefix(authHeaderValue, "Basic ")
	decoded, err := base64.StdEncoding.DecodeString(encodedHeaderContent)
	if err != nil {
		return basicAuthUser{}, fmt.Errorf("authorization header decoding: %w", err)
	}
	split := strings.Split(string(decoded), ":")
	if len(split) != 2 {
		return basicAuthUser{}, errors.New("authorization header malformed")
	}

	if b.user.Username == split[0] && b.user.Password == split[1] {
		return b.user, nil
	}

	return basicAuthUser{}, errors.New("no matching BasicAuth credentials found")
}

func authMiddleWare(validator authValidator) gen.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if len(authHeader) == 0 {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			user, err := validator.validate(authHeader)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			requestCtx := context.WithValue(r.Context(), authKey, user)
			r = r.WithContext(requestCtx)
			h.ServeHTTP(w, r)
		})
	}
}

func openApiValidator(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	if input.SecuritySchemeName != "basicAuth" {
		return fmt.Errorf("expected basicAuth security schema, received %s", input.SecuritySchemeName)
	}
	return nil
}

func openApiValidatorMiddleware(swagger *openapi3.T) gen.MiddlewareFunc {
	return middleware.OapiRequestValidatorWithOptions(swagger, &middleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: openApiValidator,
		},
		SilenceServersWarning: true,
	})
}

type responseWithStatusCode struct {
	http.ResponseWriter
	statusCode int
}

func (r *responseWithStatusCode) WriteHeader(status int) {
	r.statusCode = status
}

func loggerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rsp := &responseWithStatusCode{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}
		h.ServeHTTP(rsp, r)
		log.Println(rsp.statusCode, r.Method, r.URL.Path, time.Since(start))
	})
}

func Cors(allowedHost string) gen.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Access-Control-Allow-Origin", allowedHost)
			w.Header().Add("Access-Control-Allow-Credentials", "true")
			w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

			if r.Method == "OPTIONS" {
				http.Error(w, "No Content", http.StatusNoContent)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
