package inmem

import "github.com/rcharre/psapi/pkg/pkmn"

type InMemoryStore struct {
	PokemonStore *InMemoryPokemonStore
	TypeStore    *InMemoryTypeStore
}

// NewInMemoryStore create an in memory store with default in memory stores
func NewInMemoryStore() InMemoryStore {
	return InMemoryStore{
		PokemonStore: NewInMemoryPokemonStore(),
		TypeStore:    NewInMemoryTypeStore(),
	}
}

// GetPokemonStore return the in memory pokemon store as an abstract pokemonstore
func (s InMemoryStore) GetPokemonStore() pkmn.PokemonStore {
	return s.PokemonStore
}

// GetTypeStore return the in memory type store as an abstract typeStore
func (s InMemoryStore) GetTypeStore() pkmn.TypeStore {
	return s.TypeStore
}
