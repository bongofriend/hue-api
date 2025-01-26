package api

import (
	"errors"
	"net/http"
)

func getUserFromRequest(r *http.Request) (basicAuthUser, error) {
	authVal := r.Context().Value(authKey)
	user, ok := authVal.(basicAuthUser)
	if !ok {
		return basicAuthUser{}, errors.New("no basicAuthUser found")
	}
	return user, nil
}
