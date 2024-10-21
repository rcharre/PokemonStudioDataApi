package ps_test

import "psapi/pkg/ps"

var _ ps.TypeStore = &TypeStoreMock{}

type TypeStoreMock struct {
	FindBySymbolFunc func(string) *ps.PokemonType
	FindAllFunc      func() []*ps.PokemonType
}

func (s *TypeStoreMock) FindAll() []*ps.PokemonType {
	return s.FindAllFunc()
}

func (s *TypeStoreMock) FindBySymbol(symbol string) *ps.PokemonType {
	return s.FindBySymbolFunc(symbol)
}
