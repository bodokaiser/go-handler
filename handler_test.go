package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/check.v1"
)

func TestHandler(t *testing.T) {
	check.Suite(&Suite{})
	check.TestingT(t)
}

type Suite struct {
	auth   http.Handler
	logger http.Handler
}

var ErrTest = errors.New("a test error")

func (s *Suite) SetUpSuite(c *check.C) {
	s.auth = Auth("foo:bar", Hello())
	s.logger = Logger(Hello())
}

func (s *Suite) TestAuth(c *check.C) {
	res1 := httptest.NewRecorder()
	res2 := httptest.NewRecorder()
	res3 := httptest.NewRecorder()

	req1, _ := http.NewRequest("GET", "/", nil)
	req1.SetBasicAuth("foo", "bar")
	req2, _ := http.NewRequest("GET", "/", nil)
	req2.SetBasicAuth("foo", "baz")
	req3, _ := http.NewRequest("GET", "/", nil)
	req3.Header.Set("Authorization", "Basic ")

	s.auth.ServeHTTP(res1, req1)
	s.auth.ServeHTTP(res2, req2)
	s.auth.ServeHTTP(res3, req3)

	c.Check(res1.Code, check.Equals, http.StatusOK)
	c.Check(res2.Code, check.Equals, http.StatusUnauthorized)
	c.Check(res3.Code, check.Equals, http.StatusUnauthorized)
	c.Check(res1.Body.String(), check.Equals, "Hello World")
}

func (s *Suite) TestLogger(c *check.C) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/bla", nil)

	s.logger.ServeHTTP(res, req)

	c.Check(res.Code, check.Equals, http.StatusOK)
	c.Check(res.Body.String(), check.Equals, "Hello World")
}

func (s *Suite) TestNotFound(c *check.C) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/bla", nil)

	NotFound(res, req)

	c.Check(res.Code, check.Equals, http.StatusNotFound)
	c.Check(res.Body.String(), check.Equals, "{\"error\":\"Not Found\"}\n")
}

func Hello() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello World"))
	})
}
