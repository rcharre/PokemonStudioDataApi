package psapi

import (
	"github.com/go-chi/chi/v5"
	"github.com/rcharre/psapi/pkg/pkmn"
	"github.com/rcharre/psapi/pkg/psapi/psapigen"
)

func NewPsApiHandler(store pkmn.Store) chi.Router {
	r := chi.NewRouter()

	typeMapper := NewTypeMapper()
	typeService := NewTypeService(store.GetTypeStore(), typeMapper)
	typeController := psapigen.NewTypesAPIController(typeService)

	pokemonMapper := NewPokemonMapper(typeMapper, store.GetTypeStore())
	pokemonService := NewPokemonService(store.GetPokemonStore(), pokemonMapper)
	pokemonController := psapigen.NewPokemonAPIController(pokemonService)

	r.Mount("/", psapigen.NewRouter(pokemonController, typeController))
	return r
}
