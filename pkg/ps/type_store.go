package ps

import (
	"slices"
)

type TypeStore interface {
	FindBySymbol(symbol string) *PokemonType
	FindAll() []*PokemonType
}

type InMemoryTypeStore struct {
	pokemonTypesBySymbol map[string]*PokemonType
	types                []*PokemonType
}

func NewInMemoryTypeStore(types []*PokemonType) *InMemoryTypeStore {
	pokemonTypesBySymbol := make(map[string]*PokemonType)
	for _, pokemonType := range types {
		pokemonTypesBySymbol[pokemonType.DbSymbol] = pokemonType
	}

	return &InMemoryTypeStore{
		pokemonTypesBySymbol: pokemonTypesBySymbol,
		types:                types,
	}
}

func (s InMemoryTypeStore) FindBySymbol(symbol string) *PokemonType {
	return s.pokemonTypesBySymbol[symbol]
}

func (s InMemoryTypeStore) FindAll() []*PokemonType {
	return slices.Clone(s.types)
}
