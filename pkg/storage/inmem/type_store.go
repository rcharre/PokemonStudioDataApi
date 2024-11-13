package inmem

import "github.com/rcharre/psapi/pkg/pkmn"

type InMemoryTypeStore struct {
	pokemonTypesBySymbol map[string]pkmn.PokemonType
	types                []pkmn.PokemonType
}

// NewInMemoryTypeStore Create a new in memory type store
// types The list of types to store
func NewInMemoryTypeStore() *InMemoryTypeStore {
	return &InMemoryTypeStore{
		pokemonTypesBySymbol: make(map[string]pkmn.PokemonType),
		types:                make([]pkmn.PokemonType, 0),
	}
}

// Add add a pokemon type to the store
// pokemonType the type to add
func (s *InMemoryTypeStore) Add(pokemonType pkmn.PokemonType) {
	s.types = append(s.types, pokemonType)
	s.pokemonTypesBySymbol[pokemonType.DbSymbol] = pokemonType
}

// FindBySymbol Find a type by its symbol
// symbol The symbol to find
func (s InMemoryTypeStore) FindBySymbol(symbol string) *pkmn.PokemonType {
	pokemonType, ok := s.pokemonTypesBySymbol[symbol]
	if ok {
		copy := pokemonType
		return &copy
	} else {
		return nil
	}
}

// FindAll Find all types in the store
func (s InMemoryTypeStore) FindAll() []pkmn.PokemonType {
	return s.types
}
