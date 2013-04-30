package handlers

import (
	"fmt"
	"github.com/stretchrcom/goweb/context"
	context_test "github.com/stretchrcom/goweb/webcontext/test"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {

	handler := NewHttpHandler()

	called := false
	handler.Map("/people/{id}", func(c context.Context) error {
		called = true
		fmt.Sprintf("context: %s", c)
		return nil
	})

	assert.Equal(t, 1, len(handler.HandlersPipe()))

	ctx := context_test.MakeTestContext()
	handler.Handlers.Handle(ctx)

}
