package handlers

import (
	codecsservices "github.com/stretchr/codecs/services"
	"github.com/stretchr/goweb/controllers/test"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestOptionsListForResourceCollection(t *testing.T) {
	codecService := codecsservices.NewWebCodecService()
	h := NewHttpHandler(codecService)
	c := new(test.TestController)
	assert.Equal(t, "POST,GET,DELETE,PATCH,HEAD,OPTIONS", strings.Join(optionsListForResourceCollection(h, c), ","))

	c2 := new(test.TestSemiRestfulController)
	assert.Equal(t, "POST,GET,OPTIONS", strings.Join(optionsListForResourceCollection(h, c2), ","))

}

func TestOptionsListForSingleResource(t *testing.T) {
	codecService := codecsservices.NewWebCodecService()
	h := NewHttpHandler(codecService)
	c := new(test.TestController)
	assert.Equal(t, "GET,DELETE,PATCH,PUT,HEAD,OPTIONS", strings.Join(optionsListForSingleResource(h, c), ","))

	c2 := new(test.TestSemiRestfulController)
	assert.Equal(t, "GET,OPTIONS", strings.Join(optionsListForSingleResource(h, c2), ","))

}
