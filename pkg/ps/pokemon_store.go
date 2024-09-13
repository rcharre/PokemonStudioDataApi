package ps

import (
	"iter"
	"maps"
	"psapi/pkg/utils"
)

type PokemonStore interface {
	Add(pokemon *Pokemon)
	FindBySymbol(symbol string) *Pokemon
	FindAll(filters ...utils.FilterFunc[*Pokemon]) iter.Seq[*Pokemon]
}

type PokemonStoreImpl struct {
	pokemonBySymbol map[string]*Pokemon
}

func NewPokemonStore() PokemonStore {
	return &PokemonStoreImpl{
		pokemonBySymbol: make(map[string]*Pokemon),
	}
}

// Add Add a pokemon
func (s *PokemonStoreImpl) Add(pokemon *Pokemon) {
	s.pokemonBySymbol[pokemon.DbSymbol] = pokemon
}

// FindAll Find all pokemons with pagination and filter
func (s *PokemonStoreImpl) FindAll(filters ...utils.FilterFunc[*Pokemon]) iter.Seq[*Pokemon] {
	it := maps.Values(s.pokemonBySymbol)

	for _, filter := range filters {
		it = utils.Filter(filter, it)
	}
	return it
}

// FindBySymbol Find pokemon by symbol
func (s *PokemonStoreImpl) FindBySymbol(symbol string) *Pokemon {
	return s.pokemonBySymbol[symbol]
}
