package ps_test

import (
	"psapi/pkg/ps"
	"testing"
)

var _ ps.TypeStore = &TypeStoreMock{}

type TypeStoreMock struct {
	AddFunc          func(pokemonType *ps.PokemonType)
	FindBySymbolFunc func(string) *ps.PokemonType
	FindAllFunc      func() []*ps.PokemonType
}

func (s *TypeStoreMock) Add(pokemonType *ps.PokemonType) {
	s.Add(pokemonType)
}

func (s TypeStoreMock) FindAll() []*ps.PokemonType {
	return s.FindAllFunc()
}

func (s TypeStoreMock) FindBySymbol(symbol string) *ps.PokemonType {
	return s.FindBySymbolFunc(symbol)
}

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
