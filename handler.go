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

// Default HTTP Basic realm to use.
var DefaultRealm = "secure"

// Returns a http handler where each request is authenticated using HTTP Basic.
func Auth(username, password string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := &context.Context{Request: r, Response: w}
		h := c.Get("Authorization")

		if i := strings.IndexRune(h, ' '); i != -1 {
			b := []byte(username + ":" + password)

			if base64.StdEncoding.EncodeToString(b) == h[i+1:] {
				handler.ServeHTTP(w, r)

				return
			}
		}

		c.Set("WWW-Authenticate", fmt.Sprintf("Basic realm=%s", DefaultRealm))
		c.Error(nil, http.StatusUnauthorized)
	})
}

// Returns http handler where each request is logged to stdout.
func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.String())

		handler.ServeHTTP(w, r)
	})
}

// NotFound will send a context conform NotFound response.
func NotFound(w http.ResponseWriter, r *http.Request) {
	c := &context.Context{
		Request:  r,
		Response: w,
	}
	c.Error(nil, http.StatusNotFound)
}
