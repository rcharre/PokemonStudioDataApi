package test

import (
	"iter"

	"github.com/rcharre/psapi/pkg/ps"
	"github.com/rcharre/psapi/pkg/psapi"
	"github.com/rcharre/psapi/pkg/psapi/psapigen"
	"github.com/rcharre/psapi/pkg/utils/pagination"
)

// TYPE STORE
var _ ps.TypeStore = &ps.InMemoryTypeStore{}
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

// POKEMON STORE
var _ ps.PokemonStore = &PokemonStoreMock{}
var _ ps.PokemonStore = &ps.InMemoryPokemonStore{}

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

// POKEMON IMPORTER
var _ ps.PokemonImporter = &PokemonImporterMock{}
var _ ps.PokemonImporter = &ps.PokemonImporterImpl{}

type PokemonImporterMock struct {
	ImportFunc func(studioFolder string, translationFolder string) (iter.Seq[*ps.Pokemon], error)
}

func (i *PokemonImporterMock) Import(studioFolder string, translationFolder string) (iter.Seq[*ps.Pokemon], error) {
	return i.ImportFunc(studioFolder, translationFolder)
}

// TYPE IMPORTER

var _ ps.TypeImporter = &TypeImporterMock{}
var _ ps.TypeImporter = &ps.TypeImporterImpl{}

type TypeImporterMock struct {
	ImportFunc func(string, string) (iter.Seq[*ps.PokemonType], error)
}

func (i TypeImporterMock) Import(studioFolder string, translationFolder string) (iter.Seq[*ps.PokemonType], error) {
	return i.ImportFunc(studioFolder, translationFolder)
}

// POKEMON MAPPER
var _ psapi.PokemonMapper = &psapi.PokemonMapperImpl{}
var _ psapi.PokemonMapper = &PokemonMapperMock{}

type PokemonMapperMock struct {
	PokemonToThumbnailFunc       func(p *ps.Pokemon, lang string) *psapigen.PokemonThumbnail
	PokemonToDetailFunc          func(p *ps.Pokemon, lang string) *psapigen.PokemonDetails
	FormToPokemonFormDetailsFunc func(f *ps.PokemonForm, lang string) *psapigen.FormDetails
}

func (m *PokemonMapperMock) FormToPokemonFormDetails(f *ps.PokemonForm, lang string) *psapigen.FormDetails {
	return m.FormToPokemonFormDetailsFunc(f, lang)
}

func (m *PokemonMapperMock) PokemonToDetail(p *ps.Pokemon, lang string) *psapigen.PokemonDetails {
	return m.PokemonToDetailFunc(p, lang)
}

func (m *PokemonMapperMock) PokemonToThumbnail(p *ps.Pokemon, lang string) *psapigen.PokemonThumbnail {
	return m.PokemonToThumbnailFunc(p, lang)
}

// TYPE MAPPER
var _ psapi.TypeMapper = &psapi.TypeMapperImpl{}
var _ psapi.TypeMapper = &TypeMapperMock{}

type TypeMapperMock struct {
	toTypeDetailFunc  func(t *ps.PokemonType, lang string) *psapigen.TypeDetails
	toTypePartialFunc func(t *ps.PokemonType, lang string) *psapigen.TypePartial
}

func (m *TypeMapperMock) ToTypeDetail(t *ps.PokemonType, lang string) *psapigen.TypeDetails {
	return m.toTypeDetailFunc(t, lang)
}

func (m *TypeMapperMock) ToTypePartial(t *ps.PokemonType, lang string) *psapigen.TypePartial {
	return m.toTypePartialFunc(t, lang)
}
