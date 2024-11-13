// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Pokemon Studio API
 *
 * API for the Pokemon Studio
 *
 * API version: 0.0.1
 */

package webgen

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// TypesAPIController binds http requests to an api service and writes the service results to the http response
type TypesAPIController struct {
	service      TypesAPIServicer
	errorHandler ErrorHandler
}

// TypesAPIOption for how the controller is set up.
type TypesAPIOption func(*TypesAPIController)

// WithTypesAPIErrorHandler inject ErrorHandler into controller
func WithTypesAPIErrorHandler(h ErrorHandler) TypesAPIOption {
	return func(c *TypesAPIController) {
		c.errorHandler = h
	}
}

// NewTypesAPIController creates a default api controller
func NewTypesAPIController(s TypesAPIServicer, opts ...TypesAPIOption) *TypesAPIController {
	controller := &TypesAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the TypesAPIController
func (c *TypesAPIController) Routes() Routes {
	return Routes{
		"GetTypeDetails": Route{
			strings.ToUpper("Get"),
			"/api/types/{symbol}",
			c.GetTypeDetails,
		},
		"GetTypes": Route{
			strings.ToUpper("Get"),
			"/api/types",
			c.GetTypes,
		},
	}
}

// GetTypeDetails - Get a type details
func (c *TypesAPIController) GetTypeDetails(w http.ResponseWriter, r *http.Request) {
	symbolParam := chi.URLParam(r, "symbol")
	if symbolParam == "" {
		c.errorHandler(w, r, &RequiredError{"symbol"}, nil)
		return
	}
	acceptLanguageParam := r.Header.Get("Accept-Language")
	result, err := c.service.GetTypeDetails(r.Context(), symbolParam, acceptLanguageParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetTypes - Get all types
func (c *TypesAPIController) GetTypes(w http.ResponseWriter, r *http.Request) {
	acceptLanguageParam := r.Header.Get("Accept-Language")
	result, err := c.service.GetTypes(r.Context(), acceptLanguageParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}
