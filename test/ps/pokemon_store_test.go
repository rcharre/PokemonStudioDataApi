package ps_test

import (
	"psapi/pkg/ps"
	"psapi/pkg/utils/pagination"
	"testing"
)

var _ ps.PokemonStore = &PokemonStoreMock{}

type PokemonStoreMock struct {
	AddFunc          func(pokemon *ps.Pokemon)
	FindBySymbolFunc func(symbol string) *ps.Pokemon
	FindAllFunc      func(pageRequest pagination.PageRequest) pagination.Page[*ps.Pokemon]
}

func (s *PokemonStoreMock) Add(pokemon *ps.Pokemon) {
	return
}

func (s PokemonStoreMock) FindAll(pageRequest pagination.PageRequest) pagination.Page[*ps.Pokemon] {
	return s.FindAllFunc(pageRequest)
}

func (s PokemonStoreMock) FindBySymbol(symbol string) *ps.Pokemon {
	return s.FindBySymbolFunc(symbol)
}

func TestInMemoryPokemonStore_FindAll(t *testing.T) {
	pokemonList := []*ps.Pokemon{{
		Id:       1,
		DbSymbol: "1",
	}, {
		Id:       2,
		DbSymbol: "2",
	}, {
		Id:       4,
		DbSymbol: "4",
	}}

	store := ps.NewInMemoryPokemonStore()
	for _, pokemon := range pokemonList {
		store.Add(pokemon)
	}

	result := store.FindAll(pagination.All)

	expectLen := 3
	resultLen := len(result.Content)
	if expectLen != resultLen {
		t.Error("Expected result to have length", expectLen, ", has", resultLen)
	}
}

func TestInMemoryPokemonStore_FindBySymbol(t *testing.T) {
	pokemonList := []*ps.Pokemon{
		{
			Id:       1,
			DbSymbol: "1",
		}, {
			Id:       2,
			DbSymbol: "2",
		}, {
			Id:       4,
			DbSymbol: "4",
		},
	}
	store := ps.NewInMemoryPokemonStore()
	for _, pokemon := range pokemonList {
		store.Add(pokemon)
	}
	notFound := store.FindBySymbol("3")
	if notFound != nil {
		t.Error("Expect result to be null")
	}

	found := store.FindBySymbol("4")
	if found == nil {
		t.Error("Expect result not to be null")
	}

	if found.Id != 4 {
		t.Error("Expect result ID to be 4, is", found.Id)
	}
}
