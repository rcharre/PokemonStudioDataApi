package psapi

import (
	"github.com/go-chi/chi/v5"
	"psapi/pkg/ps"
	"psapi/pkg/psapi/psapigen"
)

func NewPsApiHandler(ctx *ps.AppContext) chi.Router {
	r := chi.NewRouter()

	pokemonMapper := NewPokemonMapper()
	pokemonService := NewPokemonService(ctx.PokemonStore(), pokemonMapper)
	pokemonController := psapigen.NewPokemonAPIController(pokemonService)

	r.Mount("/", psapigen.NewRouter(pokemonController))
	return r
}
