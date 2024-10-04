package ps

import (
	"iter"
	"log/slog"
	"psapi/pkg/utils/iter2"
	"slices"
)

type PokemonStore interface {
	Add(pokemon *Pokemon)
	FindBySymbol(symbol string) *Pokemon
	FindAll(filters ...iter2.FilterFunc[*Pokemon]) iter.Seq[*Pokemon]
}

type PokemonStoreImpl struct {
	pokemonBySymbol map[string]*Pokemon
	pokemonList     []*Pokemon
}

func NewPokemonStore() PokemonStore {
	return &PokemonStoreImpl{
		pokemonBySymbol: make(map[string]*Pokemon),
		pokemonList:     make([]*Pokemon, 0),
	}
}

// Add add the pokemon to the store
func (s *PokemonStoreImpl) Add(pokemon *Pokemon) {
	slog.Debug("Adding pokemon", "pokemon", pokemon.DbSymbol)
	index := slices.IndexFunc(s.pokemonList, func(compare *Pokemon) bool {
		return pokemon.Id < compare.Id
	})

	if index == -1 {
		index = len(s.pokemonList)
	}

	s.pokemonList = slices.Insert(s.pokemonList, index, pokemon)
	s.pokemonBySymbol[pokemon.DbSymbol] = pokemon
}

// FindAll Find all pokemon with filters
func (s *PokemonStoreImpl) FindAll(filters ...iter2.FilterFunc[*Pokemon]) iter.Seq[*Pokemon] {
	it := slices.Values(s.pokemonList)

	for _, filter := range filters {
		it = iter2.Filter(filter, it)
	}
	return it
}

// FindBySymbol Find pokemon by symbol
func (s *PokemonStoreImpl) FindBySymbol(symbol string) *Pokemon {
	return s.pokemonBySymbol[symbol]
}
