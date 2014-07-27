package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/check.v1"
)

var (
	ErrTest = errors.New("a test error")
)

func TestHandler(t *testing.T) {
	check.Suite(&HandlerSuite{
		Auth: &Auth{
			Username: "foo",
			Password: "bar",
			Handler:  Hello(),
		},
		Logger: &Logger{
			Handler: Hello(),
		},
	})
	check.TestingT(t)
}

type HandlerSuite struct {
	Auth   *Auth
	Logger *Logger
}

func (s *HandlerSuite) TestAuth(c *check.C) {
	var req *http.Request
	var res *httptest.ResponseRecorder

	res = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/", nil)
	req.SetBasicAuth("foo", "bar")
	s.Auth.ServeHTTP(res, req)
	c.Check(res.Code, check.Equals, http.StatusOK)
	c.Check(res.Body.String(), check.Equals, "Hello World")

	res = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/", nil)
	req.SetBasicAuth("foo", "baz")
	s.Auth.ServeHTTP(res, req)
	c.Check(res.Code, check.Equals, http.StatusUnauthorized)

	res = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Basic ")
	s.Auth.ServeHTTP(res, req)
	c.Check(res.Code, check.Equals, http.StatusUnauthorized)
}

func (s *HandlerSuite) TestLogger(c *check.C) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/bla", nil)
	s.Logger.ServeHTTP(res, req)
	c.Check(res.Code, check.Equals, http.StatusOK)
	c.Check(res.Body.String(), check.Equals, "Hello World")
}

func (s *HandlerSuite) TestNotFound(c *check.C) {
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
