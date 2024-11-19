package inmem

import (
	"slices"
	"sort"

	"github.com/rcharre/psapi/pkg/pkmn"
	"github.com/rcharre/psapi/pkg/utils/pagination"
)

type InMemoryPokemonStore struct {
	pokemonBySymbol map[string]pkmn.Pokemon
	pokemonList     []pkmn.Pokemon
}

// NewInMemoryPokemonStore Create an in memory pokemon store
func NewInMemoryPokemonStore() *InMemoryPokemonStore {
	pokemonBySymbol := make(map[string]pkmn.Pokemon)
	pokemonList := make([]pkmn.Pokemon, 0)

	return &InMemoryPokemonStore{
		pokemonBySymbol: pokemonBySymbol,
		pokemonList:     pokemonList,
	}
}

func (s *InMemoryPokemonStore) Add(pokemon pkmn.Pokemon) {
	s.pokemonBySymbol[pokemon.DbSymbol] = pokemon
	index := sort.Search(len(s.pokemonList), func(i int) bool {
		return s.pokemonList[i].Id >= pokemon.Id
	})

	s.pokemonList = slices.Insert(s.pokemonList, index, pokemon)
}

// FindAll Find a page of pokemon corresponding to the page request
// pageRequest the page request
func (s InMemoryPokemonStore) FindAll(pageRequest pagination.PageRequest) pagination.Page[pkmn.Pokemon] {
	pokemonPage := pagination.ApplyPageRequest(pageRequest, s.pokemonList)
	pokemonPage.Content = slices.Clone(pokemonPage.Content)
	return pokemonPage
}

// FindBySymbol Find pokemon by symbol
// symbol The symbol of the pokemon to find
func (s InMemoryPokemonStore) FindBySymbol(symbol string) *pkmn.Pokemon {
	pokemon, ok := s.pokemonBySymbol[symbol]
	if ok {
		copy := pokemon
		return &copy
	} else {
		return nil
	}
}
