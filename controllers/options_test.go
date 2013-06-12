package controllers

import (
	"github.com/stretchr/goweb/controllers/test"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestOptionsListForResourceCollection(t *testing.T) {

	c := new(test.TestController)
	assert.Equal(t, "POST,GET,DELETE,PUT,HEAD,OPTIONS", strings.Join(OptionsListForResourceCollection(c), ","))

	c2 := new(test.TestSemiRestfulController)
	assert.Equal(t, "POST,GET,OPTIONS", strings.Join(OptionsListForResourceCollection(c2), ","))

}

func TestOptionsListForSingleResource(t *testing.T) {

	c := new(test.TestController)
	assert.Equal(t, "GET,DELETE,PUT,POST,HEAD,OPTIONS", strings.Join(OptionsListForSingleResource(c), ","))

	c2 := new(test.TestSemiRestfulController)
	assert.Equal(t, "GET,OPTIONS", strings.Join(OptionsListForSingleResource(c2), ","))

}
