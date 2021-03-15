package middleware

import (
	"net/http"
	"strings"

	"github.com/FideTech/yaus/config"
	"github.com/FideTech/yaus/utils"
)

type APIAuthHandlerFunc func(http.ResponseWriter, *http.Request)

//Logger is a middleware handler that does request logging
type APIAuth struct {
	handler APIAuthHandlerFunc
}

//ServeHTTP handles the request by passing it to the real handler and logging the request details
func (api *APIAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("authorization")
	if authHeader == "" {
		http.Error(w, "authorization required (missing header)", http.StatusUnauthorized)
		return
	}

	ahSplit := strings.Split(authHeader, " ")
	if len(ahSplit) != 2 {
		http.Error(w, "authorization required (invalid header)", http.StatusUnauthorized)
		return
	}

	if !utils.StringContains(config.Config.System.API.Keys, ahSplit[1]) {
		http.Error(w, "authorization required (invalid key)", http.StatusUnauthorized)
		return
	}

	api.handler(w, r)
}

//NewLogger constructs a new Logger middleware handler
func NewAPIAuth(handlerToWrap APIAuthHandlerFunc) *APIAuth {
	return &APIAuth{handlerToWrap}
}
