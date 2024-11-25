package studio_test

import (
	"testing"

	"github.com/rcharre/psapi/pkg/studio"
	"github.com/rcharre/psapi/pkg/utils/pagination"
)

func TestInMemoryPokemonStore_FindAll(t *testing.T) {
	pokemonList := []studio.Pokemon{{
		Id:       1,
		DbSymbol: "1",
	}, {
		Id:       2,
		DbSymbol: "2",
	}, {
		Id:       4,
		DbSymbol: "4",
	}}

	store := studio.NewPokemonStore()
	for _, pokemon := range pokemonList {
		store.Add(pokemon)
	}

	result := store.FindAll(pagination.NewPageRequest(0, 1000))

	expectLen := 3
	resultLen := len(result.Content)
	if expectLen != resultLen {
		t.Error("Expected result to have length", expectLen, ", has", resultLen)
	}
}

func TestInMemoryPokemonStore_FindBySymbol(t *testing.T) {
	pokemonList := []studio.Pokemon{
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
	store := studio.NewPokemonStore()
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
