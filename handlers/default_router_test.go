package handlers

import (
	"testing"
	"github.com/stretchrcom/testify/assert"
)

func TestDefaultRouter(t *testing.T) {

	assert.Implements(t, (*Router)(nil), new(DefaultRouter), "DefaultRouter should implement Router interface")
	router := new(DefaultRouter)

	assert.NotNil(t, router)

}

func TestDefaultRouter_AddHandler(t *testing.T) {

	

}
