package ps_test

import (
	"github.com/rcharre/psapi/pkg/ps"
	"testing"
)

func TestAddAndFindBySymbol(t *testing.T) {
	store := ps.NewInMemoryTypeStore()
	pokeType := &ps.PokemonType{
		DbSymbol: "test",
	}
	store.Add(pokeType)

	find := store.FindBySymbol(pokeType.DbSymbol)
	if find == nil {
		t.Error("Should find type with symbol", pokeType.DbSymbol)
	}
}

func TestFindAll(t *testing.T) {

	store := ps.NewInMemoryTypeStore()
	store.Add(&ps.PokemonType{
		DbSymbol: "1",
	})

	store.Add(&ps.PokemonType{
		DbSymbol: "2",
	})

	store.Add(&ps.PokemonType{
		DbSymbol: "3",
	})

	all := store.FindAll()
	allLen := len(all)
	if allLen != 3 {
		t.Error("Find all length should be 3, has", allLen)
	}
}
