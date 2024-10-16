package ps

import (
	"iter"
	"psapi/pkg/utils/iter2"
	"slices"
)

type PokemonStore interface {
	FindBySymbol(symbol string) *Pokemon
	FindAll(filters ...iter2.FilterFunc[*Pokemon]) iter.Seq[*Pokemon]
}

type InMemoryPokemonStore struct {
	pokemonBySymbol map[string]*Pokemon
	pokemonList     []*Pokemon
}

func ComparePokemonId(p1, p2 *Pokemon) int {
	if p1.Id >= p2.Id {
		return 1
	} else {
		return -1
	}
}

func NewInMemoryPokemonStore(pokemonList []*Pokemon) *InMemoryPokemonStore {
	pokemonBySymbol := make(map[string]*Pokemon)
	slices.SortFunc(pokemonList, ComparePokemonId)

	for _, pokemon := range pokemonList {
		pokemonBySymbol[pokemon.DbSymbol] = pokemon
	}
	return &InMemoryPokemonStore{
		pokemonBySymbol: pokemonBySymbol,
		pokemonList:     pokemonList,
	}
}

// FindAll Find all pokemon with filters
func (s *InMemoryPokemonStore) FindAll(filters ...iter2.FilterFunc[*Pokemon]) iter.Seq[*Pokemon] {
	it := slices.Values(s.pokemonList)

	for _, filter := range filters {
		it = iter2.Filter(filter, it)
	}
	return it
}

// FindBySymbol Find pokemon by symbol
func (s *InMemoryPokemonStore) FindBySymbol(symbol string) *Pokemon {
	return s.pokemonBySymbol[symbol]
}
