package ps_test

import (
	"iter"
	"psapi/pkg/ps"
	"psapi/pkg/utils/iter2"
	"slices"
	"testing"
)

var _ ps.PokemonStore = &PokemonStoreMock{}

type PokemonStoreMock struct {
	FindBySymbolFunc func(symbol string) *ps.Pokemon
	FindAllFunc      func(filters ...iter2.FilterFunc[*ps.Pokemon]) iter.Seq[*ps.Pokemon]
}

func (s *PokemonStoreMock) FindAll(filters ...iter2.FilterFunc[*ps.Pokemon]) iter.Seq[*ps.Pokemon] {
	return s.FindAllFunc(filters...)
}

func (s *PokemonStoreMock) FindBySymbol(symbol string) *ps.Pokemon {
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
	store := ps.NewInMemoryPokemonStore(pokemonList)
	it := store.FindAll(func(p *ps.Pokemon) bool {
		return p.Id <= 2
	})

	result := slices.Collect(it)
	expectLen := 2
	resultLen := len(result)
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
	store := ps.NewInMemoryPokemonStore(pokemonList)
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
