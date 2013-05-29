package controllers

import (
	"github.com/stretchrcom/goweb/controllers/test"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestOptionsListForResourceCollection(t *testing.T) {

	c := new(test.TestController)
	assert.Equal(t, "POST,GET,DELETE,PUT,HEAD,OPTIONS", OptionsListForResourceCollection(c))

	c2 := new(test.TestSemiRestfulController)
	assert.Equal(t, "POST,GET,OPTIONS", OptionsListForResourceCollection(c2))

}

func TestOptionsListForSingleResource(t *testing.T) {

	c := new(test.TestController)
	assert.Equal(t, "GET,DELETE,PUT,POST,HEAD,OPTIONS", OptionsListForSingleResource(c))

	c2 := new(test.TestSemiRestfulController)
	assert.Equal(t, "GET,OPTIONS", OptionsListForSingleResource(c2))

}
