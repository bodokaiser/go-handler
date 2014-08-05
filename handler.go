// The handler package provides common http.Handler for use in HTTP APIs.
package handler

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/satisfeet/go-context"
)

// Interface implemented by all handlers.
type Handler interface {
	// Handler wraps the given http.Handler.
	Handle(http.Handler) http.Handler
}

// Auth implements http.Handler and only calls another http.Handler when
// requests have valid credentials defined by http basic authorization.
type Auth struct {
	Error    error
	Username string
	Password string
	Handler  http.Handler
}

// Default HTTP Basic realm to use.
var DefaultRealm = "secure"

// Defines the http.Handler to secure.
// Returns the top level http.Handler for easy chaining.
func (a *Auth) Handle(handler http.Handler) http.Handler {
	a.Handler = handler

	return a
}

// Implementation of http.Handler interface. Contains the HTTP Basic logic.
func (a *Auth) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c := &context.Context{
		Request:  request,
		Response: writer,
	}
	h := c.Get("Authorization")

	if i := strings.IndexRune(h, ' '); i != -1 {
		b := []byte(a.Username + ":" + a.Password)

		if base64.StdEncoding.EncodeToString(b) == h[i+1:] {
			if a.Handler != nil {
				a.Handler.ServeHTTP(writer, request)
			}

			return
		}
	}

	c.Set("WWW-Authenticate", fmt.Sprintf("Basic realm=%s", DefaultRealm))
	c.Error(a.Error, http.StatusUnauthorized)
}

// Logger prints method and url of each request.
type Logger struct {
	Handler http.Handler
}

// Defines the http.Handler to log.
// Returns the top level http.Handler for easy chaining.
func (l *Logger) Handle(handler http.Handler) http.Handler {
	l.Handler = handler

	return l
}

// Implementation of the logger algorithm.
func (l *Logger) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Printf("%s %s", request.Method, request.URL.String())

	l.Handler.ServeHTTP(writer, request)
}

// NotFound will send a context conform NotFound response.
func NotFound(w http.ResponseWriter, r *http.Request) {
	c := &context.Context{
		Request:  r,
		Response: w,
	}
	c.Error(nil, http.StatusNotFound)
}
