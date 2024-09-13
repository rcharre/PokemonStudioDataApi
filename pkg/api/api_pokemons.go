// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Pokemon Studio API
 *
 * API for the Pokemon Studio
 *
 * API version: 0.0.1
 */

package api

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// PokemonsAPIController binds http requests to an api service and writes the service results to the http response
type PokemonsAPIController struct {
	service PokemonsAPIServicer
	errorHandler ErrorHandler
}

// PokemonsAPIOption for how the controller is set up.
type PokemonsAPIOption func(*PokemonsAPIController)

// WithPokemonsAPIErrorHandler inject ErrorHandler into controller
func WithPokemonsAPIErrorHandler(h ErrorHandler) PokemonsAPIOption {
	return func(c *PokemonsAPIController) {
		c.errorHandler = h
	}
}

// NewPokemonsAPIController creates a default api controller
func NewPokemonsAPIController(s PokemonsAPIServicer, opts ...PokemonsAPIOption) *PokemonsAPIController {
	controller := &PokemonsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the PokemonsAPIController
func (c *PokemonsAPIController) Routes() Routes {
	return Routes{
		"GetPokemon": Route{
			strings.ToUpper("Get"),
			"/api/pokemons/{symbol}",
			c.GetPokemon,
		},
		"GetPokemonForm": Route{
			strings.ToUpper("Get"),
			"/api/pokemons/{symbol}/{form}",
			c.GetPokemonForm,
		},
		"GetPokemons": Route{
			strings.ToUpper("Get"),
			"/api/pokemons",
			c.GetPokemons,
		},
	}
}

// GetPokemon - Get a pokemon details
func (c *PokemonsAPIController) GetPokemon(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	symbolParam := chi.URLParam(r, "symbol")
	if symbolParam == "" {
		c.errorHandler(w, r, &RequiredError{"symbol"}, nil)
		return
	}
	var langParam string
	if query.Has("lang") {
		param := query.Get("lang")

		langParam = param
	} else {
		param := "en"
		langParam = param
	}
	result, err := c.service.GetPokemon(r.Context(), symbolParam, langParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetPokemonForm - Get a pokemon form details
func (c *PokemonsAPIController) GetPokemonForm(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	symbolParam := chi.URLParam(r, "symbol")
	if symbolParam == "" {
		c.errorHandler(w, r, &RequiredError{"symbol"}, nil)
		return
	}
	formParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "form"),
		WithDefaultOrParse[int32](0, parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Param: "form", Err: err}, nil)
		return
	}
	var langParam string
	if query.Has("lang") {
		param := query.Get("lang")

		langParam = param
	} else {
		param := "en"
		langParam = param
	}
	result, err := c.service.GetPokemonForm(r.Context(), symbolParam, formParam, langParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetPokemons - Get a page of pokemons
func (c *PokemonsAPIController) GetPokemons(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	var pageParam int32
	if query.Has("page") {
		param, err := parseNumericParameter[int32](
			query.Get("page"),
			WithParse[int32](parseInt32),
			WithMinimum[int32](0),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Param: "page", Err: err}, nil)
			return
		}

		pageParam = param
	} else {
		var param int32 = 0
		pageParam = param
	}
	var sizeParam int32
	if query.Has("size") {
		param, err := parseNumericParameter[int32](
			query.Get("size"),
			WithParse[int32](parseInt32),
			WithMinimum[int32](1),
			WithMaximum[int32](50),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Param: "size", Err: err}, nil)
			return
		}

		sizeParam = param
	} else {
		var param int32 = 20
		sizeParam = param
	}
	var langParam string
	if query.Has("lang") {
		param := query.Get("lang")

		langParam = param
	} else {
		param := "en"
		langParam = param
	}
	result, err := c.service.GetPokemons(r.Context(), pageParam, sizeParam, langParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}
