package psapi

import (
	"github.com/go-chi/chi/v5"
	"psapi/pkg/ps"
	"psapi/pkg/psapi/psapigen"
)

func NewPsApiHandler(studio *ps.Studio) chi.Router {
	r := chi.NewRouter()

	typeMapper := NewTypeMapper()
	typeService := NewTypeService(studio.TypeStore(), typeMapper)
	typeController := psapigen.NewTypesAPIController(typeService)

	pokemonMapper := NewPokemonMapper(typeMapper, studio.TypeStore())
	pokemonService := NewPokemonService(studio.PokemonStore(), pokemonMapper)
	pokemonController := psapigen.NewPokemonAPIController(pokemonService)

	r.Mount("/", psapigen.NewRouter(pokemonController, typeController))
	return r
}
