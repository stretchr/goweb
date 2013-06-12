package test

import (
	"net/http"
)

func TestWeb(t *testing.T) {

	goweb.Test(t, "GET people/123", func(t *testing.T, response http.Response) {

		assert.Equal(t, 200, response.StatusCode())

	})

}
