package ps

import (
	"slices"
)

var _ TypeStore = &InMemoryTypeStore{}

type TypeStore interface {
	FindBySymbol(symbol string) *PokemonType
	FindAll() []*PokemonType
}

type InMemoryTypeStore struct {
	pokemonTypesBySymbol map[string]*PokemonType
	types                []*PokemonType
}

// NewInMemoryTypeStore Create a new in memory type store
// types The list of types to store
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

// FindBySymbol Find a type by its symbol
// symbol The symbol to find
func (s InMemoryTypeStore) FindBySymbol(symbol string) *PokemonType {
	return s.pokemonTypesBySymbol[symbol]
}

// FindAll Find all types in the store
func (s InMemoryTypeStore) FindAll() []*PokemonType {
	return slices.Clone(s.types)
}
