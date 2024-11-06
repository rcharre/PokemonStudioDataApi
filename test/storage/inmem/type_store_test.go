package inmem_test

import (
	"testing"

	"github.com/rcharre/psapi/pkg/pkmn"
	"github.com/rcharre/psapi/pkg/storage/inmem"
)

func TestAddAndFindBySymbol(t *testing.T) {
	store := inmem.NewInMemoryTypeStore()
	pokeType := &pkmn.PokemonType{
		DbSymbol: "test",
	}
	store.Add(pokeType)

	find := store.FindBySymbol(pokeType.DbSymbol)
	if find == nil {
		t.Error("Should find type with symbol", pokeType.DbSymbol)
	}
}

func TestFindAll(t *testing.T) {

	store := inmem.NewInMemoryTypeStore()
	store.Add(&pkmn.PokemonType{
		DbSymbol: "1",
	})

	store.Add(&pkmn.PokemonType{
		DbSymbol: "2",
	})

	store.Add(&pkmn.PokemonType{
		DbSymbol: "3",
	})

	all := store.FindAll()
	allLen := len(all)
	if allLen != 3 {
		t.Error("Find all length should be 3, has", allLen)
	}
}
