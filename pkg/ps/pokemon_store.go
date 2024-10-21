package ps

import (
	"iter"
	"psapi/pkg/utils/iter2"
	"slices"
)

var _ PokemonStore = &InMemoryPokemonStore{}

type PokemonStore interface {
	FindBySymbol(symbol string) *Pokemon
	FindAll(filters ...iter2.FilterFunc[*Pokemon]) iter.Seq[*Pokemon]
}

type InMemoryPokemonStore struct {
	pokemonBySymbol map[string]*Pokemon
	pokemonList     []*Pokemon
}

// NewInMemoryPokemonStore Create an in memory pokemon store
// pokemonList a pokemon list to store
func NewInMemoryPokemonStore(pokemonList []*Pokemon) *InMemoryPokemonStore {
	pokemonBySymbol := make(map[string]*Pokemon)
	toStore := make([]*Pokemon, len(pokemonList))

	for i, pokemon := range pokemonList {
		copy := *pokemon
		pokemonBySymbol[pokemon.DbSymbol] = &copy
		toStore[i] = &copy
	}

	slices.SortFunc(toStore, ComparePokemonId)
	return &InMemoryPokemonStore{
		pokemonBySymbol: pokemonBySymbol,
		pokemonList:     toStore,
	}
}

// FindAll Find all pokemon corresponding to a list of filter and return an iterator
// filters a list of filter to use on the store
func (s InMemoryPokemonStore) FindAll(filters ...iter2.FilterFunc[*Pokemon]) iter.Seq[*Pokemon] {
	it := slices.Values(s.pokemonList)

	for _, filter := range filters {
		it = iter2.Filter(filter, it)
	}

	return it
}

// FindBySymbol Find pokemon by symbol
// symbol The symbol of the pokemon to find
func (s InMemoryPokemonStore) FindBySymbol(symbol string) *Pokemon {
	return s.pokemonBySymbol[symbol]
}
