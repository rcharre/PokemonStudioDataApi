package ps_test

import (
	"iter"
	"psapi/pkg/ps"
	"psapi/pkg/utils/iter2"
	"testing"
)

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

var _ ps.PokemonStore = &PokemonStoreMock{}

func TestComparePokemonId(t *testing.T) {
	p1 := &ps.Pokemon{
		Id: 1,
	}

	p2 := &ps.Pokemon{
		Id: 2,
	}

	if ps.ComparePokemonId(p1, p2) != -1 {
		t.Error("ComparePokemonId with p1:", p1.Id, "and p2:", p2.Id, "should return -1")
	}
	if ps.ComparePokemonId(p2, p1) != 1 {
		t.Error("ComparePokemonId with p2:", p2.Id, "and p1:", p1.Id, "should return -1")
	}
}
