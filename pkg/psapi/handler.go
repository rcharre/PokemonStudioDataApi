package psapi

import (
	"github.com/go-chi/chi/v5"
	"github.com/rcharre/psapi/pkg/pkmn"
	"github.com/rcharre/psapi/pkg/psapi/psapigen"
)

func MakeDefaultRouter(store pkmn.Store) chi.Router {
	typeMapper := NewTypeMapper()
	typeService := NewTypeService(store.GetTypeStore(), typeMapper)
	typeController := psapigen.NewTypesAPIController(typeService)

	pokemonMapper := NewPokemonMapper(typeMapper, store.GetTypeStore())
	pokemonService := NewPokemonService(store.GetPokemonStore(), pokemonMapper)
	pokemonController := psapigen.NewPokemonAPIController(pokemonService)

	return psapigen.NewRouter(pokemonController, typeController)
}
