package handlers

import (
	"fmt"
	codecservices "github.com/stretchrcom/codecs/services"
	"github.com/stretchrcom/goweb/context"
	controllers_test "github.com/stretchrcom/goweb/controllers/test"
	context_test "github.com/stretchrcom/goweb/webcontext/test"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {

	codecService := new(codecservices.WebCodecService)
	handler := NewHttpHandler(codecService)

	called := false
	handler.Map("/people/{id}", func(c context.Context) error {
		called = true
		return nil
	})

	assert.Equal(t, 1, len(handler.HandlersPipe()))

	ctx := context_test.MakeTestContextWithPath("people/123")
	handler.Handlers.Handle(ctx)

	assert.True(t, called)

}

func TestMap_WithSpecificMethod(t *testing.T) {

	codecService := new(codecservices.WebCodecService)
	handler := NewHttpHandler(codecService)

	called := false
	handler.Map("GET", "/people/{id}", func(c context.Context) error {
		called = true
		return nil
	})

	assert.Equal(t, 1, len(handler.HandlersPipe()))

	ctx := context_test.MakeTestContextWithPath("people/123")
	handler.Handlers.Handle(ctx)

	assert.True(t, called)
	assert.Equal(t, "GET", handler.HandlersPipe()[0].(*PathMatchHandler).HttpMethods[0])

}

func TestMap_WithMatcherFuncs(t *testing.T) {

	codecService := new(codecservices.WebCodecService)
	handler := NewHttpHandler(codecService)

	matcherFunc := MatcherFunc(func(c context.Context) (MatcherFuncDecision, error) {
		return Match, nil
	})

	handler.Map("/people/{id}", func(c context.Context) error {
		return nil
	}, matcherFunc)

	assert.Equal(t, 1, len(handler.HandlersPipe()))
	h := handler.HandlersPipe()[0].(*PathMatchHandler)
	assert.Equal(t, 1, len(h.MatcherFuncs))
	assert.Equal(t, matcherFunc, h.MatcherFuncs[0], "Matcher func (first)")

}

func TestMap_CatchAllAssumption(t *testing.T) {

	codecService := new(codecservices.WebCodecService)
	handler := NewHttpHandler(codecService)

	called := false
	handler.Map(func(c context.Context) error {
		called = true
		return nil
	})

	assert.Equal(t, 1, len(handler.HandlersPipe()))

	ctx := context_test.MakeTestContextWithPath("people/123")
	handler.Handlers.Handle(ctx)
	assert.True(t, called)

	called = false
	ctx = context_test.MakeTestContextWithPath("something-else")
	handler.Handlers.Handle(ctx)
	assert.True(t, called)

}

func assertPathMatchHandler(t *testing.T, handler *PathMatchHandler, path, method string, message string) bool {

	if assert.NotNil(t, handler) {

		ctx := context_test.MakeTestContextWithDetails(path, method)

		willHandle, _ := handler.WillHandle(ctx)
		if assert.True(t, willHandle, fmt.Sprintf("This handler is expected to handle it: %s", message)) {

			// make sure the method is in the list
			methodFound := false
			for _, methodInList := range handler.HttpMethods {
				if methodInList == method {
					methodFound = true
					break
				}
			}

			return assert.True(t, methodFound, "Method (%s) should be in the method list (%s)", method, handler.HttpMethods)
		}

	}

	return false

}

func TestMapRest_SemiInterface(t *testing.T) {

	semi := new(controllers_test.TestSemiRestfulController)

	codecService := new(codecservices.WebCodecService)
	h := NewHttpHandler(codecService)
	h.MapController(semi)

	fmt.Printf("%s", h)

	assert.Equal(t, 3, len(h.HandlersPipe()))

	// create
	assertPathMatchHandler(t, h.HandlersPipe()[0].(*PathMatchHandler), "/test-semi-restful", "POST", "create")

	// read one
	assertPathMatchHandler(t, h.HandlersPipe()[1].(*PathMatchHandler), "/test-semi-restful/123", "GET", "read one")

	// read many
	assertPathMatchHandler(t, h.HandlersPipe()[2].(*PathMatchHandler), "/test-semi-restful", "GET", "read many")

}

func TestMapRest(t *testing.T) {

	rest := new(controllers_test.TestController)

	codecService := new(codecservices.WebCodecService)
	h := NewHttpHandler(codecService)
	h.MapController(rest)

	assert.Equal(t, 10, len(h.HandlersPipe()))

	// create
	assertPathMatchHandler(t, h.HandlersPipe()[0].(*PathMatchHandler), "/test", "POST", "create")

	// read one
	assertPathMatchHandler(t, h.HandlersPipe()[1].(*PathMatchHandler), "/test/123", "GET", "read one")

	// read many
	assertPathMatchHandler(t, h.HandlersPipe()[2].(*PathMatchHandler), "/test", "GET", "read many")

	// delete one
	assertPathMatchHandler(t, h.HandlersPipe()[3].(*PathMatchHandler), "/test/123", "DELETE", "delete one")

	// delete many
	assertPathMatchHandler(t, h.HandlersPipe()[4].(*PathMatchHandler), "/test", "DELETE", "delete many")

	// update one
	assertPathMatchHandler(t, h.HandlersPipe()[5].(*PathMatchHandler), "/test/123", "PUT", "update one")

	// update many
	assertPathMatchHandler(t, h.HandlersPipe()[6].(*PathMatchHandler), "/test", "PUT", "update many")

	// replace one
	assertPathMatchHandler(t, h.HandlersPipe()[7].(*PathMatchHandler), "/test/123", "POST", "replace")

	// head
	assertPathMatchHandler(t, h.HandlersPipe()[8].(*PathMatchHandler), "/test/123", "HEAD", "head")
	assertPathMatchHandler(t, h.HandlersPipe()[8].(*PathMatchHandler), "/test", "HEAD", "head")

	// options
	assertPathMatchHandler(t, h.HandlersPipe()[9].(*PathMatchHandler), "/test/123", "OPTIONS", "options")
	assertPathMatchHandler(t, h.HandlersPipe()[9].(*PathMatchHandler), "/test", "OPTIONS", "options")

}

func TestMapRest_WithSpecificPath(t *testing.T) {

	rest := new(controllers_test.TestController)

	codecService := new(codecservices.WebCodecService)
	h := NewHttpHandler(codecService)
	h.MapController("something", rest)

	assert.Equal(t, 10, len(h.HandlersPipe()))

	// create
	assertPathMatchHandler(t, h.HandlersPipe()[0].(*PathMatchHandler), "/something", "POST", "create")

	// read one
	assertPathMatchHandler(t, h.HandlersPipe()[1].(*PathMatchHandler), "/something/123", "GET", "read one")

	// read many
	assertPathMatchHandler(t, h.HandlersPipe()[2].(*PathMatchHandler), "/something", "GET", "read many")

	// delete one
	assertPathMatchHandler(t, h.HandlersPipe()[3].(*PathMatchHandler), "/something/123", "DELETE", "delete one")

	// delete many
	assertPathMatchHandler(t, h.HandlersPipe()[4].(*PathMatchHandler), "/something", "DELETE", "delete many")

	// update one
	assertPathMatchHandler(t, h.HandlersPipe()[5].(*PathMatchHandler), "/something/123", "PUT", "update one")

	// update many
	assertPathMatchHandler(t, h.HandlersPipe()[6].(*PathMatchHandler), "/something", "PUT", "update many")

	// replace one
	assertPathMatchHandler(t, h.HandlersPipe()[7].(*PathMatchHandler), "/something/123", "POST", "replace")

	// head
	assertPathMatchHandler(t, h.HandlersPipe()[8].(*PathMatchHandler), "/something/123", "HEAD", "head")
	assertPathMatchHandler(t, h.HandlersPipe()[8].(*PathMatchHandler), "/something", "HEAD", "head")

	// options
	assertPathMatchHandler(t, h.HandlersPipe()[9].(*PathMatchHandler), "/something/123", "OPTIONS", "options")
	assertPathMatchHandler(t, h.HandlersPipe()[9].(*PathMatchHandler), "/something", "OPTIONS", "options")

}
