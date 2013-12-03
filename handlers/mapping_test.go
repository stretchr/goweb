package handlers

import (
	"fmt"
	codecsservices "github.com/stretchr/codecs/services"
	"github.com/stretchr/goweb/context"
	controllers_test "github.com/stretchr/goweb/controllers/test"
	handlers_test "github.com/stretchr/goweb/handlers/test"
	goweb_http "github.com/stretchr/goweb/http"
	context_test "github.com/stretchr/goweb/webcontext/test"
	"github.com/stretchr/testify/assert"
	http_test "github.com/stretchr/testify/http"
	"github.com/stretchr/testify/mock"
	"log"
	"net/http"
	"testing"
)

func TestFindMatcherFuncs(t *testing.T) {

	matcher := func(ctx context.Context) (MatcherFuncDecision, error) {
		return Match, nil
	}

	castMatcher := MatcherFunc(func(ctx context.Context) (MatcherFuncDecision, error) {
		return Match, nil
	})

	matchers := []MatcherFunc{func(ctx context.Context) (MatcherFuncDecision, error) {
		return Match, nil
	}}

	allMatchers := append(matchers, castMatcher, MatcherFunc(matcher))

	assert.Equal(t, allMatchers, findMatcherFuncs(matchers, castMatcher, matcher))

}

func TestHandlerForOptions_PlainHandler(t *testing.T) {

	codecService := codecsservices.NewWebCodecService()
	httpHandler := NewHttpHandler(codecService)
	handler1 := new(handlers_test.TestHandler)

	itself, _ := httpHandler.handlerForOptions(handler1)

	assert.Equal(t, handler1, itself, "handlerForOptions with existing handler should just return the handler")

}

// https://github.com/stretchr/goweb/issues/19
func TestMappedHandlersBreakExecution(t *testing.T) {

	codecService := codecsservices.NewWebCodecService()
	handler := NewHttpHandler(codecService)

	handlerCalled := false
	catchAllCalled := false
	handler.Map("/people/{id}", func(c context.Context) error {
		handlerCalled = true
		return nil
	})
	handler.Map(func(c context.Context) error {
		catchAllCalled = true
		return nil
	})

	ctx := context_test.MakeTestContextWithPath("people/123")
	handler.Handlers.Handle(ctx)

	assert.True(t, handlerCalled)
	assert.False(t, catchAllCalled, "Catch-all should NOT get called, becuase something else specifically handled this context.")

}

func TestMap(t *testing.T) {

	codecService := codecsservices.NewWebCodecService()
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

	codecService := codecsservices.NewWebCodecService()
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
	assert.True(t, handler.HandlersPipe()[0].(*PathMatchHandler).BreakCurrentPipeline)

}

func TestMap_WithSpecificMethods(t *testing.T) {

	codecService := codecsservices.NewWebCodecService()
	handler := NewHttpHandler(codecService)

	called := false
	handler.Map([]string{"GET", "POST"}, "/people/{id}", func(c context.Context) error {
		called = true
		return nil
	})

	assert.Equal(t, 1, len(handler.HandlersPipe()))

	ctx := context_test.MakeTestContextWithPath("people/123")
	handler.Handlers.Handle(ctx)

	assert.True(t, called)
	assert.Equal(t, "GET", handler.HandlersPipe()[0].(*PathMatchHandler).HttpMethods[0])
	assert.Equal(t, "POST", handler.HandlersPipe()[0].(*PathMatchHandler).HttpMethods[1])

}

func TestMap_WithMatcherFuncs(t *testing.T) {

	codecService := codecsservices.NewWebCodecService()
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

	codecService := codecsservices.NewWebCodecService()
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

func TestMapBefore_WithMatcherFuncs(t *testing.T) {

	codecService := codecsservices.NewWebCodecService()
	handler := NewHttpHandler(codecService)

	matcherFunc := MatcherFunc(func(c context.Context) (MatcherFuncDecision, error) {
		return Match, nil
	})

	handler.MapBefore("/people/{id}", func(c context.Context) error {
		return nil
	}, matcherFunc)

	assert.Equal(t, 1, len(handler.PreHandlersPipe()))
	h := handler.PreHandlersPipe()[0].(*PathMatchHandler)
	assert.Equal(t, 1, len(h.MatcherFuncs))
	assert.Equal(t, matcherFunc, h.MatcherFuncs[0], "Matcher func (first)")
	assert.False(t, handler.PreHandlersPipe()[0].(*PathMatchHandler).BreakCurrentPipeline)

}

func TestMapAfter_WithMatcherFuncs(t *testing.T) {

	codecService := codecsservices.NewWebCodecService()
	handler := NewHttpHandler(codecService)

	matcherFunc := MatcherFunc(func(c context.Context) (MatcherFuncDecision, error) {
		return Match, nil
	})

	handler.MapAfter("/people/{id}", func(c context.Context) error {
		return nil
	}, matcherFunc)

	assert.Equal(t, 1, len(handler.PostHandlersPipe()))
	h := handler.PostHandlersPipe()[0].(*PathMatchHandler)
	assert.Equal(t, 1, len(h.MatcherFuncs))
	assert.Equal(t, matcherFunc, h.MatcherFuncs[0], "Matcher func (first)")
	assert.False(t, handler.PostHandlersPipe()[0].(*PathMatchHandler).BreakCurrentPipeline)

}

func TestBeforeAndAfterHandlers(t *testing.T) {

	responseWriter := new(http_test.TestResponseWriter)
	testRequest, _ := http.NewRequest("GET", "http://stretchr.org/goweb", nil)
	codecService := codecsservices.NewWebCodecService()
	handler := NewHttpHandler(codecService)

	// setup some test handlers
	handler1 := new(handlers_test.TestHandler)
	handler2 := new(handlers_test.TestHandler)
	handler3 := new(handlers_test.TestHandler)

	handler.MapBefore(func(c context.Context) error {
		_, err := handler1.Handle(c)
		return err
	})
	handler.Map(func(c context.Context) error {
		_, err := handler2.Handle(c)
		return err
	})
	handler.MapAfter(func(c context.Context) error {
		_, err := handler3.Handle(c)
		return err
	})

	handler1.On("Handle", mock.Anything).Return(false, nil)
	handler2.On("Handle", mock.Anything).Return(false, nil)
	handler3.On("Handle", mock.Anything).Return(false, nil)

	handler.ServeHTTP(responseWriter, testRequest)

	mock.AssertExpectationsForObjects(t, handler1.Mock, handler2.Mock, handler3.Mock)

	// make sure it's always the same context
	ctx1 := handler1.Calls[0].Arguments[0].(context.Context)
	ctx2 := handler2.Calls[0].Arguments[0].(context.Context)
	ctx3 := handler3.Calls[0].Arguments[0].(context.Context)

	assert.Equal(t, ctx1, ctx2, "Contexts should be the same")
	assert.Equal(t, ctx2, ctx3, "Contexts should be the same")

}

/*
	MapController
*/

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

	codecService := codecsservices.NewWebCodecService()
	h := NewHttpHandler(codecService)
	h.MapController(semi)

	assert.Equal(t, 5, len(h.HandlersPipe()))

	// create
	assertPathMatchHandler(t, h.HandlersPipe()[0].(*PathMatchHandler), "/test-semi-restful", "POST", "create")

	// read one
	assertPathMatchHandler(t, h.HandlersPipe()[1].(*PathMatchHandler), "/test-semi-restful/123", "GET", "read one")

	// read many
	assertPathMatchHandler(t, h.HandlersPipe()[2].(*PathMatchHandler), "/test-semi-restful", "GET", "read many")

}

func TestMapController(t *testing.T) {

	rest := new(controllers_test.TestController)

	codecService := codecsservices.NewWebCodecService()
	h := NewHttpHandler(codecService)
	h.MapController(rest)

	assert.Equal(t, 10, len(h.HandlersPipe()))

	// create
	assertPathMatchHandler(t, h.HandlersPipe()[0].(*PathMatchHandler), "/test", goweb_http.MethodPost, "create")

	// read one
	assertPathMatchHandler(t, h.HandlersPipe()[1].(*PathMatchHandler), "/test/123", goweb_http.MethodGet, "read one")

	// read many
	assertPathMatchHandler(t, h.HandlersPipe()[2].(*PathMatchHandler), "/test", goweb_http.MethodGet, "read many")

	// delete one
	assertPathMatchHandler(t, h.HandlersPipe()[3].(*PathMatchHandler), "/test/123", goweb_http.MethodDelete, "delete one")

	// delete many
	assertPathMatchHandler(t, h.HandlersPipe()[4].(*PathMatchHandler), "/test", goweb_http.MethodDelete, "delete many")

	// update one
	assertPathMatchHandler(t, h.HandlersPipe()[5].(*PathMatchHandler), "/test/123", goweb_http.MethodPatch, "update one")

	// update many
	assertPathMatchHandler(t, h.HandlersPipe()[6].(*PathMatchHandler), "/test", goweb_http.MethodPatch, "update many")

	// replace one
	assertPathMatchHandler(t, h.HandlersPipe()[7].(*PathMatchHandler), "/test/123", goweb_http.MethodPut, "replace")

	// head
	assertPathMatchHandler(t, h.HandlersPipe()[8].(*PathMatchHandler), "/test/123", goweb_http.MethodHead, goweb_http.MethodHead)
	assertPathMatchHandler(t, h.HandlersPipe()[8].(*PathMatchHandler), "/test", goweb_http.MethodHead, goweb_http.MethodHead)

	// options
	assertPathMatchHandler(t, h.HandlersPipe()[9].(*PathMatchHandler), "/test/123", goweb_http.MethodOptions, goweb_http.MethodOptions)
	assertPathMatchHandler(t, h.HandlersPipe()[9].(*PathMatchHandler), "/test", goweb_http.MethodOptions, goweb_http.MethodOptions)

}

func TestNewHttpHandler_SetsDefaultMethod(t *testing.T) {

	h := NewHttpHandler(nil)

	assert.Equal(t, h.HttpMethodForCreate, goweb_http.MethodPost)
	assert.Equal(t, h.HttpMethodForReadOne, goweb_http.MethodGet)
	assert.Equal(t, h.HttpMethodForReadMany, goweb_http.MethodGet)
	assert.Equal(t, h.HttpMethodForDeleteOne, goweb_http.MethodDelete)
	assert.Equal(t, h.HttpMethodForDeleteMany, goweb_http.MethodDelete)
	assert.Equal(t, h.HttpMethodForUpdateOne, goweb_http.MethodPatch)
	assert.Equal(t, h.HttpMethodForUpdateMany, goweb_http.MethodPatch)
	assert.Equal(t, h.HttpMethodForReplace, goweb_http.MethodPut)
	assert.Equal(t, h.HttpMethodForHead, goweb_http.MethodHead)
	assert.Equal(t, h.HttpMethodForOptions, goweb_http.MethodOptions)

}

func TestMapController_WithExplicitHTTPMethods(t *testing.T) {

	rest := new(controllers_test.TestController)

	codecService := codecsservices.NewWebCodecService()
	h := NewHttpHandler(codecService)

	h.HttpMethodForCreate = "CREATE"
	h.HttpMethodForReadOne = "READ_ONE"
	h.HttpMethodForReadMany = "READ_MANY"
	h.HttpMethodForDeleteOne = "DELETE_ONE"
	h.HttpMethodForDeleteMany = "DELETE_MANY"
	h.HttpMethodForUpdateOne = "UPDATE_ONE"
	h.HttpMethodForUpdateMany = "UPDATE_MANY"
	h.HttpMethodForReplace = "REPLACE"
	h.HttpMethodForHead = "HEAD_CUSTOM"
	h.HttpMethodForOptions = "OPTIONS_CUSTOM"

	h.MapController(rest)

	assert.Equal(t, 10, len(h.HandlersPipe()))

	// create
	assertPathMatchHandler(t, h.HandlersPipe()[0].(*PathMatchHandler), "/test", "CREATE", "create")

	// read one
	assertPathMatchHandler(t, h.HandlersPipe()[1].(*PathMatchHandler), "/test/123", "READ_ONE", "read one")

	// read many
	assertPathMatchHandler(t, h.HandlersPipe()[2].(*PathMatchHandler), "/test", "READ_MANY", "read many")

	// delete one
	assertPathMatchHandler(t, h.HandlersPipe()[3].(*PathMatchHandler), "/test/123", "DELETE_ONE", "delete one")

	// delete many
	assertPathMatchHandler(t, h.HandlersPipe()[4].(*PathMatchHandler), "/test", "DELETE_MANY", "delete many")

	// update one
	assertPathMatchHandler(t, h.HandlersPipe()[5].(*PathMatchHandler), "/test/123", "UPDATE_ONE", "update one")

	// update many
	assertPathMatchHandler(t, h.HandlersPipe()[6].(*PathMatchHandler), "/test", "UPDATE_MANY", "update many")

	// replace one
	assertPathMatchHandler(t, h.HandlersPipe()[7].(*PathMatchHandler), "/test/123", "REPLACE", "replace")

	// head
	assertPathMatchHandler(t, h.HandlersPipe()[8].(*PathMatchHandler), "/test/123", "HEAD_CUSTOM", "head")
	assertPathMatchHandler(t, h.HandlersPipe()[8].(*PathMatchHandler), "/test", "HEAD_CUSTOM", "head")

	// options
	assertPathMatchHandler(t, h.HandlersPipe()[9].(*PathMatchHandler), "/test/123", "OPTIONS_CUSTOM", "options")
	assertPathMatchHandler(t, h.HandlersPipe()[9].(*PathMatchHandler), "/test", "OPTIONS_CUSTOM", "options")

}

func TestMapController_WithMatcherFuncs(t *testing.T) {
	rest := new(controllers_test.TestController)

	codecService := codecsservices.NewWebCodecService()
	handler := NewHttpHandler(codecService)

	matcherFunc := func(ctx context.Context) (MatcherFuncDecision, error) {
		return Match, nil
	}

	handler.MapController(rest, matcherFunc)

	assert.Equal(t, 10, len(handler.HandlersPipe()))

	castMatcherFunc := MatcherFunc(matcherFunc)
	var h *PathMatchHandler
	for i := 0; i < 10; i++ {
		h = handler.HandlersPipe()[i].(*PathMatchHandler)
		assert.Equal(t, 1, len(h.MatcherFuncs))
		assert.Equal(t, castMatcherFunc, h.MatcherFuncs[0], "Matcher func (first)")
	}

}

func TestMapController_WithSpecificPath(t *testing.T) {

	rest := new(controllers_test.TestController)

	codecService := codecsservices.NewWebCodecService()
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
	assertPathMatchHandler(t, h.HandlersPipe()[5].(*PathMatchHandler), "/something/123", "PATCH", "update one")

	// update many
	assertPathMatchHandler(t, h.HandlersPipe()[6].(*PathMatchHandler), "/something", "PATCH", "update many")

	// replace one
	assertPathMatchHandler(t, h.HandlersPipe()[7].(*PathMatchHandler), "/something/123", "PUT", "replace")

	// head
	assertPathMatchHandler(t, h.HandlersPipe()[8].(*PathMatchHandler), "/something/123", "HEAD", "head")
	assertPathMatchHandler(t, h.HandlersPipe()[8].(*PathMatchHandler), "/something", "HEAD", "head")

	// options
	assertPathMatchHandler(t, h.HandlersPipe()[9].(*PathMatchHandler), "/something/123", "OPTIONS", "options")
	assertPathMatchHandler(t, h.HandlersPipe()[9].(*PathMatchHandler), "/something", "OPTIONS", "options")

}

func TestMapController_DefaultOptions(t *testing.T) {

	semi := new(controllers_test.TestSemiRestfulController)

	codecService := codecsservices.NewWebCodecService()
	h := NewHttpHandler(codecService)
	h.MapController(semi)

	assert.Equal(t, 5, len(h.HandlersPipe()))

	// get the last two
	handler1 := h.HandlersPipe()[len(h.HandlersPipe())-2]

	handler2 := h.HandlersPipe()[len(h.HandlersPipe())-1]
	assertPathMatchHandler(t, handler1.(*PathMatchHandler), "/test-semi-restful", "OPTIONS", "options")
	assertPathMatchHandler(t, handler2.(*PathMatchHandler), "/test-semi-restful/{id}", "OPTIONS", "options")

}

/*
	Before and After handlers
*/

func TestBeforeHandler(t *testing.T) {

	cont := new(controllers_test.TestHandlerWithBeforeAndAfters)

	codecService := codecsservices.NewWebCodecService()
	h := NewHttpHandler(codecService)

	h.MapController(cont)

	log.Printf("%s", h)

	if assert.Equal(t, 2, len(h.PreHandlersPipe()), "2 pre handler's expected") {
		assertPathMatchHandler(t, h.PreHandlersPipe()[0].(*PathMatchHandler), "/test", "POST", "before POST /test")
		assertPathMatchHandler(t, h.PreHandlersPipe()[1].(*PathMatchHandler), "/test/123", "PUT", "before PUT /test/123")
		assertPathMatchHandler(t, h.PreHandlersPipe()[0].(*PathMatchHandler), "/test", "OPTIONS", "before OPTIONS /test")
		assertPathMatchHandler(t, h.PreHandlersPipe()[1].(*PathMatchHandler), "/test/123", "OPTIONS", "before OPTIONS /test/123")
	}

	if assert.Equal(t, 2, len(h.PostHandlersPipe()), "2 post handler's expected") {
		assertPathMatchHandler(t, h.PostHandlersPipe()[0].(*PathMatchHandler), "/test", "POST", "after POST /test")
		assertPathMatchHandler(t, h.PostHandlersPipe()[1].(*PathMatchHandler), "/test/123", "PUT", "after PUT /test/123")
		assertPathMatchHandler(t, h.PostHandlersPipe()[0].(*PathMatchHandler), "/test", "OPTIONS", "after OPTIONS /test")
		assertPathMatchHandler(t, h.PostHandlersPipe()[1].(*PathMatchHandler), "/test/123", "OPTIONS", "after OPTIONS /test/123")
	}

}

func TestMapStatic(t *testing.T) {

	codecService := codecsservices.NewWebCodecService()
	h := NewHttpHandler(codecService)

	h.MapStatic("/static", "/location/of/static")

	assert.Equal(t, 1, len(h.HandlersPipe()))

	staticHandler := h.HandlersPipe()[0].(*PathMatchHandler)

	if assert.Equal(t, 1, len(staticHandler.HttpMethods)) {
		assert.Equal(t, goweb_http.MethodGet, staticHandler.HttpMethods[0])
	}

	var ctx context.Context
	var willHandle bool

	ctx = context_test.MakeTestContextWithPath("/static/some/deep/file.dat")
	willHandle, _ = staticHandler.WillHandle(ctx)
	assert.True(t, willHandle, "Static handler should handle")

	ctx = context_test.MakeTestContextWithPath("/static/../static/some/deep/file.dat")
	willHandle, _ = staticHandler.WillHandle(ctx)
	assert.True(t, willHandle, "Static handler should handle")

	ctx = context_test.MakeTestContextWithPath("/static/some/../file.dat")
	willHandle, _ = staticHandler.WillHandle(ctx)
	assert.True(t, willHandle, "Static handler should handle")

	ctx = context_test.MakeTestContextWithPath("/static/../file.dat")
	willHandle, _ = staticHandler.WillHandle(ctx)
	assert.False(t, willHandle, "Static handler should not handle")

	ctx = context_test.MakeTestContextWithPath("/static")
	willHandle, _ = staticHandler.WillHandle(ctx)
	assert.True(t, willHandle, "Static handler should handle")

	ctx = context_test.MakeTestContextWithPath("/static/")
	willHandle, _ = staticHandler.WillHandle(ctx)
	assert.True(t, willHandle, "Static handler should handle")

	ctx = context_test.MakeTestContextWithPath("/static/doc.go")
	willHandle, _ = staticHandler.WillHandle(ctx)
	_, staticHandleErr := staticHandler.Handle(ctx)

	if assert.NoError(t, staticHandleErr) {

	}

}

func TestMapStatic_WithMatcherFuncs(t *testing.T) {

	codecService := codecsservices.NewWebCodecService()
	h := NewHttpHandler(codecService)

	matcherFunc := MatcherFunc(func(c context.Context) (MatcherFuncDecision, error) {
		return Match, nil
	})

	h.MapStatic("/static", "/location/of/static", matcherFunc)

	assert.Equal(t, 1, len(h.HandlersPipe()))
	staticHandler := h.HandlersPipe()[0].(*PathMatchHandler)
	assert.Equal(t, 1, len(staticHandler.MatcherFuncs))
	assert.Equal(t, matcherFunc, staticHandler.MatcherFuncs[0], "Matcher func (first)")
}

func TestMapStaticFile(t *testing.T) {

	codecService := codecsservices.NewWebCodecService()
	h := NewHttpHandler(codecService)

	h.MapStaticFile("/static-file", "/location/of/static-file")

	assert.Equal(t, 1, len(h.HandlersPipe()))

	staticHandler := h.HandlersPipe()[0].(*PathMatchHandler)

	if assert.Equal(t, 1, len(staticHandler.HttpMethods)) {
		assert.Equal(t, goweb_http.MethodGet, staticHandler.HttpMethods[0])
	}

	var ctx context.Context
	var willHandle bool

	ctx = context_test.MakeTestContextWithPath("/static-file")
	willHandle, _ = staticHandler.WillHandle(ctx)
	assert.True(t, willHandle, "Static handler should handle")

	ctx = context_test.MakeTestContextWithPath("static-file")
	willHandle, _ = staticHandler.WillHandle(ctx)
	assert.True(t, willHandle, "Static handler should handle")

	ctx = context_test.MakeTestContextWithPath("static-file/")
	willHandle, _ = staticHandler.WillHandle(ctx)
	assert.True(t, willHandle, "Static handler should handle")

	ctx = context_test.MakeTestContextWithPath("static-file/something-else")
	willHandle, _ = staticHandler.WillHandle(ctx)
	assert.False(t, willHandle, "Static handler NOT should handle")

}

func TestMapStaticFile_WithMatcherFuncs(t *testing.T) {

	codecService := codecsservices.NewWebCodecService()
	h := NewHttpHandler(codecService)

	matcherFunc := MatcherFunc(func(c context.Context) (MatcherFuncDecision, error) {
		return Match, nil
	})

	h.MapStaticFile("/static-file", "/location/of/static-file", matcherFunc)

	assert.Equal(t, 1, len(h.HandlersPipe()))
	staticHandler := h.HandlersPipe()[0].(*PathMatchHandler)
	assert.Equal(t, 1, len(staticHandler.MatcherFuncs))
	assert.Equal(t, matcherFunc, staticHandler.MatcherFuncs[0], "Matcher func (first)")
}
