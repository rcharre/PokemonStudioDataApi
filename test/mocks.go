package test

import (
	"github.com/rcharre/psapi/pkg/pkmn"
	"github.com/rcharre/psapi/pkg/psapi"
	"github.com/rcharre/psapi/pkg/psapi/psapigen"
	"github.com/rcharre/psapi/pkg/storage/inmem"
	"github.com/rcharre/psapi/pkg/utils/pagination"
)

// STORE
var _ pkmn.Store = &inmem.InMemoryStore{}
var _ pkmn.Store = &StoreMock{}

type StoreMock struct {
	PokemonStore *PokemonStoreMock
	TypeStore    *TypeStoreMock
}

func NewStoreMock() *StoreMock {
	return &StoreMock{
		&PokemonStoreMock{},
		&TypeStoreMock{},
	}
}

func (s *StoreMock) GetPokemonStore() pkmn.PokemonStore {
	return s.PokemonStore
}

func (s *StoreMock) GetTypeStore() pkmn.TypeStore {
	return s.TypeStore
}

// TYPE STORE
var _ pkmn.TypeStore = &inmem.InMemoryTypeStore{}
var _ pkmn.TypeStore = &TypeStoreMock{}

type TypeStoreMock struct {
	AddFunc          func(pokemonType *pkmn.PokemonType)
	FindBySymbolFunc func(string) *pkmn.PokemonType
	FindAllFunc      func() []*pkmn.PokemonType
}

func (s *TypeStoreMock) Add(pokemonType *pkmn.PokemonType) {
	if s.AddFunc == nil {
		panic("TypeStoreMock func 'Add' not mocked")
	}
	s.Add(pokemonType)
}

func (s TypeStoreMock) FindAll() []*pkmn.PokemonType {
	if s.FindAllFunc == nil {
		panic("TypeStoreMock func 'FindAll' not mocked")
	}
	return s.FindAllFunc()
}

func (s TypeStoreMock) FindBySymbol(symbol string) *pkmn.PokemonType {
	if s.FindBySymbolFunc == nil {
		panic("TypeStoreMock func 'FindBySymbol' not mocked")
	}
	return s.FindBySymbolFunc(symbol)
}

// POKEMON STORE
var _ pkmn.PokemonStore = &PokemonStoreMock{}
var _ pkmn.PokemonStore = &inmem.InMemoryPokemonStore{}

type PokemonStoreMock struct {
	AddFunc          func(pokemon *pkmn.Pokemon)
	FindBySymbolFunc func(symbol string) *pkmn.Pokemon
	FindAllFunc      func(pageRequest pagination.PageRequest) pagination.Page[*pkmn.Pokemon]
}

func (s *PokemonStoreMock) Add(pokemon *pkmn.Pokemon) {
	if s.AddFunc == nil {
		panic("PokemonStoreMock func 'Add' not mocked")
	}
	s.AddFunc(pokemon)
}

func (s PokemonStoreMock) FindAll(pageRequest pagination.PageRequest) pagination.Page[*pkmn.Pokemon] {
	if s.FindAllFunc == nil {
		panic("PokemonStoreMock func 'FindAll' not mocked")
	}
	return s.FindAllFunc(pageRequest)
}

func (s PokemonStoreMock) FindBySymbol(symbol string) *pkmn.Pokemon {
	if s.FindBySymbolFunc == nil {
		panic("PokemonStoreMock func 'FindBySymbol' not mocked")
	}
	return s.FindBySymbolFunc(symbol)
}

// POKEMON MAPPER
var _ psapi.PokemonMapper = &psapi.PokemonMapperImpl{}
var _ psapi.PokemonMapper = &PokemonMapperMock{}

type PokemonMapperMock struct {
	PokemonToThumbnailFunc       func(p *pkmn.Pokemon, lang string) *psapigen.PokemonThumbnail
	PokemonToDetailFunc          func(p *pkmn.Pokemon, lang string) *psapigen.PokemonDetails
	FormToPokemonFormDetailsFunc func(f *pkmn.PokemonForm, lang string) *psapigen.FormDetails
}

func (m *PokemonMapperMock) FormToPokemonFormDetails(f *pkmn.PokemonForm, lang string) *psapigen.FormDetails {
	if m.FormToPokemonFormDetailsFunc == nil {
		panic("PokemonMapperMock func 'FormToPokemonFormDetails' not mocked")
	}
	return m.FormToPokemonFormDetailsFunc(f, lang)
}

func (m *PokemonMapperMock) PokemonToDetail(p *pkmn.Pokemon, lang string) *psapigen.PokemonDetails {
	if m.PokemonToDetailFunc == nil {
		panic("PokemonMapperMock func 'PokemonToDetail' not mocked")
	}
	return m.PokemonToDetailFunc(p, lang)
}

func (m *PokemonMapperMock) PokemonToThumbnail(p *pkmn.Pokemon, lang string) *psapigen.PokemonThumbnail {
	if m.PokemonToThumbnailFunc == nil {
		panic("PokemonMapperMock func 'PokemonToThumbnail' not mocked")
	}
	return m.PokemonToThumbnailFunc(p, lang)
}

// TYPE MAPPER
var _ psapi.TypeMapper = &psapi.TypeMapperImpl{}
var _ psapi.TypeMapper = &TypeMapperMock{}

type TypeMapperMock struct {
	toTypeDetailFunc  func(t *pkmn.PokemonType, lang string) *psapigen.TypeDetails
	toTypePartialFunc func(t *pkmn.PokemonType, lang string) *psapigen.TypePartial
}

func (m *TypeMapperMock) ToTypeDetail(t *pkmn.PokemonType, lang string) *psapigen.TypeDetails {
	if m.toTypeDetailFunc == nil {
		panic("TypeMapperMock func 'ToTypeDetail' not mocked")
	}
	return m.toTypeDetailFunc(t, lang)
}

func (m *TypeMapperMock) ToTypePartial(t *pkmn.PokemonType, lang string) *psapigen.TypePartial {
	if m.toTypePartialFunc == nil {
		panic("TypeMapperMock func 'ToTypePartial' not mocked")
	}
	return m.toTypePartialFunc(t, lang)
}
