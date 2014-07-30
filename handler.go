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

// Auth is a http.Handler which secures a http handler from  requests which do
// not have valid http basic authorization.
type Auth struct {
	Error    error
	Username string
	Password string
	Handler  http.Handler
}

// Default HTTP Basic realm to use.
var DefaultRealm = "secure"

func (a *Auth) Handle(h http.Handler) http.Handler {
	a.Handler = h

	return a
}

func (a *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := &context.Context{
		Request:  r,
		Response: w,
	}
	h := c.Get("Authorization")

	if i := strings.IndexRune(h, ' '); i != -1 {
		b := []byte(a.Username + ":" + a.Password)

		if base64.StdEncoding.EncodeToString(b) == h[i+1:] {
			if a.Handler != nil {
				a.Handler.ServeHTTP(w, r)
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

func (l *Logger) Handle(h http.Handler) http.Handler {
	l.Handler = h

	return l
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.String())

	l.Handler.ServeHTTP(w, r)
}

// NotFound will send a context conform NotFound response.
func NotFound(w http.ResponseWriter, r *http.Request) {
	c := &context.Context{
		Request:  r,
		Response: w,
	}
	c.Error(nil, http.StatusNotFound)
}
