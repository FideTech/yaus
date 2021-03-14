package middleware

import (
	"net/http"
)

//Logger is a middleware handler that does request logging
type Cors struct {
	handler http.Handler
}

//ServeHTTP handles the request by passing it to the real handler and logging the request details
func (c *Cors) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Content-Length, Last-Event-ID, X-Requested-With")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type, Content-Disposition, Authorization, Cache-Control, Expires, Pragma, X-Server")
	w.Header().Set("Cache-Control", "private, no-cache, no-store, must-revalidate")
	w.Header().Set("Expires", "-1")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("X-Server", "yaus")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	c.handler.ServeHTTP(w, r)
}

//NewLogger constructs a new Logger middleware handler
func NewCors(handlerToWrap http.Handler) *Cors {
	return &Cors{handlerToWrap}
}
