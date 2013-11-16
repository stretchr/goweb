package handlers

import (
	"github.com/stretchr/goweb/controllers"
	"github.com/stretchr/goweb/http"
)

func optionsListForResourceCollection(handler *HttpHandler, controller interface{}) []string {

	var methods []string

	// POST /resource  -  Create
	if _, ok := controller.(controllers.RestfulCreator); ok {
		methods = append(methods, handler.HttpMethodForCreate)
	}

	// GET /resource  -  ReadMany
	if _, ok := controller.(controllers.RestfulManyReader); ok {
		methods = append(methods, handler.HttpMethodForReadMany)
	}

	// DELETE /resource  -  DeleteMany
	if _, ok := controller.(controllers.RestfulManyDeleter); ok {
		methods = append(methods, handler.HttpMethodForDeleteMany)
	}

	// PATCH /resource  -  UpdateMany
	if _, ok := controller.(controllers.RestfulManyUpdater); ok {
		methods = append(methods, handler.HttpMethodForUpdateMany)
	}

	// HEAD /resource/[id]  -  Head
	if _, ok := controller.(controllers.RestfulHead); ok {
		methods = append(methods, handler.HttpMethodForHead)
	}

	if _, ok := controller.(controllers.RestfulOptions); ok {
		methods = append(methods, handler.HttpMethodForOptions)
	} else {
		methods = append(methods, http.MethodOptions)
	}

	return methods

}

func optionsListForSingleResource(handler *HttpHandler, controller interface{}) []string {

	var methods []string

	// GET /resource/{id}  -  ReadOne
	if _, ok := controller.(controllers.RestfulReader); ok {
		methods = append(methods, handler.HttpMethodForReadOne)
	}

	// DELETE /resource/{id}  -  DeleteOne
	if _, ok := controller.(controllers.RestfulDeletor); ok {
		methods = append(methods, handler.HttpMethodForDeleteOne)
	}

	// PATCH /resource/{id}  -  Update
	if _, ok := controller.(controllers.RestfulUpdater); ok {
		methods = append(methods, handler.HttpMethodForUpdateOne)

	}

	// PUT /resource/{id}  -  Replace
	if _, ok := controller.(controllers.RestfulReplacer); ok {
		methods = append(methods, handler.HttpMethodForReplace)

	}

	// HEAD /resource/[id]  -  Head
	if _, ok := controller.(controllers.RestfulHead); ok {
		methods = append(methods, handler.HttpMethodForHead)
	}

	if _, ok := controller.(controllers.RestfulOptions); ok {
		methods = append(methods, handler.HttpMethodForOptions)
	} else {
		methods = append(methods, http.MethodOptions)
	}

	return methods

}
