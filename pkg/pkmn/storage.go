package pkmn

import "github.com/rcharre/psapi/pkg/utils/pagination"

type PokemonStore interface {
	Add(pokemon Pokemon)
	FindBySymbol(symbol string) *Pokemon
	FindAll(pageRequest pagination.PageRequest) pagination.Page[Pokemon]
}

type TypeStore interface {
	Add(pokemonType PokemonType)
	FindBySymbol(symbol string) *PokemonType
	FindAll() []PokemonType
}

type Store interface {
	GetPokemonStore() PokemonStore
	GetTypeStore() TypeStore
}
