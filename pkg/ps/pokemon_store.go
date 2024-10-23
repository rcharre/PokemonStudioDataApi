package ps

import (
	"psapi/pkg/utils/pagination"
	"slices"
	"sort"
)

var _ PokemonStore = &InMemoryPokemonStore{}

type PokemonStore interface {
	Add(pokemon *Pokemon)
	FindBySymbol(symbol string) *Pokemon
	FindAll(pageRequest pagination.PageRequest) pagination.Page[*Pokemon]
}

type InMemoryPokemonStore struct {
	pokemonBySymbol map[string]*Pokemon
	pokemonList     []*Pokemon
}

// NewInMemoryPokemonStore Create an in memory pokemon store
func NewInMemoryPokemonStore() *InMemoryPokemonStore {
	pokemonBySymbol := make(map[string]*Pokemon)
	pokemonList := make([]*Pokemon, 0)

	return &InMemoryPokemonStore{
		pokemonBySymbol: pokemonBySymbol,
		pokemonList:     pokemonList,
	}
}

func (s *InMemoryPokemonStore) Add(pokemon *Pokemon) {
	s.pokemonBySymbol[pokemon.DbSymbol] = pokemon
	index := sort.Search(len(s.pokemonList), func(i int) bool {
		return s.pokemonList[i].Id >= pokemon.Id
	})

	s.pokemonList = slices.Insert(s.pokemonList, index, pokemon)
}

// FindAll Find a page of pokemon corresponding to the page request
// pageRequest the page request
func (s InMemoryPokemonStore) FindAll(pageRequest pagination.PageRequest) pagination.Page[*Pokemon] {
	return pagination.ApplyPageRequest(pageRequest, s.pokemonList)
}

// FindBySymbol Find pokemon by symbol
// symbol The symbol of the pokemon to find
func (s InMemoryPokemonStore) FindBySymbol(symbol string) *Pokemon {
	return s.pokemonBySymbol[symbol]
}
