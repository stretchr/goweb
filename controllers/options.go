package controllers

import (
	"github.com/stretchr/goweb/http"
)

func OptionsListForResourceCollection(controller interface{}) []string {

	var methods []string

	// POST /resource  -  Create
	if _, ok := controller.(RestfulCreator); ok {
		methods = append(methods, http.MethodPost)
	}

	// GET /resource  -  ReadMany
	if _, ok := controller.(RestfulManyReader); ok {
		methods = append(methods, http.MethodGet)
	}

	// DELETE /resource  -  DeleteMany
	if _, ok := controller.(RestfulManyDeleter); ok {
		methods = append(methods, http.MethodDelete)
	}

	// PUT /resource  -  UpdateMany
	if _, ok := controller.(RestfulManyUpdater); ok {
		methods = append(methods, http.MethodPut)
	}

	// HEAD /resource/[id]  -  Head
	if _, ok := controller.(RestfulHead); ok {
		methods = append(methods, http.MethodHead)
	}

	methods = append(methods, http.MethodOptions)

	return methods

}

func OptionsListForSingleResource(controller interface{}) []string {

	var methods []string

	// GET /resource/{id}  -  ReadOne
	if _, ok := controller.(RestfulReader); ok {
		methods = append(methods, http.MethodGet)
	}

	// DELETE /resource/{id}  -  DeleteOne
	if _, ok := controller.(RestfulDeletor); ok {
		methods = append(methods, http.MethodDelete)
	}

	// PUT /resource/{id}  -  Update
	if _, ok := controller.(RestfulUpdater); ok {
		methods = append(methods, http.MethodPut)

	}

	// POST /resource/{id}  -  Replace
	if _, ok := controller.(RestfulReplacer); ok {
		methods = append(methods, http.MethodPost)

	}

	// HEAD /resource/[id]  -  Head
	if _, ok := controller.(RestfulHead); ok {
		methods = append(methods, http.MethodHead)
	}

	methods = append(methods, http.MethodOptions)
	return methods

}
