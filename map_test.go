package goweb

import (
	"fmt"
	"github.com/stretchrcom/goweb/context"
	"github.com/stretchrcom/goweb/handlers"
	context_test "github.com/stretchrcom/goweb/webcontext/test"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {

	defaultHttpHandler = handlers.NewHttpHandler()

	called := false
	Map("/people/{id}", func(c context.Context) error {

		API.Respond(c, 200, "data", nil)

		called = true
		fmt.Sprintf("context: %s", c)
		return nil
	})

	assert.Equal(t, 1, len(defaultHttpHandler.HandlersPipe()))

	ctx := context_test.MakeTestContext()
	defaultHttpHandler.Handlers.Handle(ctx)

}
