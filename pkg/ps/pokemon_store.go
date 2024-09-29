package ps

import (
	"iter"
	"psapi/pkg/utils"
	"slices"
)

type PokemonStore interface {
	SetPokemonList(pokemon []*Pokemon, translations []translation)
	FindBySymbol(symbol string) *Pokemon
	FindAll(filters ...utils.FilterFunc[*Pokemon]) iter.Seq[*Pokemon]
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

// SetPokemonList Set the pokemon list of the store and reindex maps
func (s *PokemonStoreImpl) SetPokemonList(pokemonList []*Pokemon, translations []translation) {
	s.pokemonBySymbol = make(map[string]*Pokemon)

	slices.SortFunc(pokemonList, ComparePokemon)
	s.pokemonList = pokemonList

	for index, pkmn := range pokemonList {
		pkmn.Translations = translations[index+1] // translation container egg
		s.pokemonBySymbol[pkmn.DbSymbol] = pkmn
	}
}

// FindAll Find all pokemon with filters
func (s *PokemonStoreImpl) FindAll(filters ...utils.FilterFunc[*Pokemon]) iter.Seq[*Pokemon] {
	it := slices.Values(s.pokemonList)

	for _, filter := range filters {
		it = utils.Filter(filter, it)
	}
	return it
}

// FindBySymbol Find pokemon by symbol
func (s *PokemonStoreImpl) FindBySymbol(symbol string) *Pokemon {
	return s.pokemonBySymbol[symbol]
}
