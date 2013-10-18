package goweb

import (
	"github.com/stretchr/testify/assert"
	testhttp "github.com/stretchr/testify/http"
	"net/http"
	"testing"
)

func TestRedirectTo(t *testing.T) {

	r := new(testhttp.TestResponseWriter)
	RedirectTo(r, "http://www.stretchr.com", "test")
	assert.Equal(t, "http://www.stretchr.com/test", r.Header().Get("Location"))
	assert.Equal(t, 0, r.StatusCode)

}

func TestRedirect(t *testing.T) {

	r := new(testhttp.TestResponseWriter)
	Redirect(r, "http://www.stretchr.com", "test")
	assert.Equal(t, "http://www.stretchr.com/test", r.Header().Get("Location"))
	assert.Equal(t, http.StatusFound, r.StatusCode)

}

func TestRedirectTemp(t *testing.T) {

	r := new(testhttp.TestResponseWriter)
	RedirectTemp(r, "http://www.stretchr.com", "test")
	assert.Equal(t, "http://www.stretchr.com/test", r.Header().Get("Location"))
	assert.Equal(t, http.StatusTemporaryRedirect, r.StatusCode)

}

func TestRedirectPerm(t *testing.T) {

	r := new(testhttp.TestResponseWriter)
	RedirectPerm(r, "http://www.stretchr.com", "test")
	assert.Equal(t, "http://www.stretchr.com/test", r.Header().Get("Location"))
	assert.Equal(t, http.StatusMovedPermanently, r.StatusCode)

}
