package ps

import (
	"path"
)

const (
	StudioFolder   = "Studio"
	LanguageFolder = "Text/Dialogs"
)

type Studio struct {
	TypeStore       TypeStore
	TypeImporter    TypeImporter
	PokemonStore    PokemonStore
	PokemonImporter PokemonImporter
}

func NewStudio(typeStore TypeStore, typeImporter TypeImporter, pokemonStore PokemonStore, pokemonImporter PokemonImporter) *Studio {
	return &Studio{
		typeStore,
		typeImporter,
		pokemonStore,
		pokemonImporter,
	}
}

func NewInMemoryStudio() *Studio {
	return &Studio{
		TypeStore:       NewInMemoryTypeStore(),
		TypeImporter:    NewTypeImporter(),
		PokemonStore:    NewInMemoryPokemonStore(),
		PokemonImporter: NewPokemonImporter(),
	}
}

func (s *Studio) Import(folder string) error {
	translationFolder := path.Join(folder, LanguageFolder)
	studioFolder := path.Join(folder, StudioFolder)

	if err := s.importTypes(studioFolder, translationFolder); err != nil {
		return err
	}
	if err := s.importPokemon(studioFolder, translationFolder); err != nil {
		return err
	}
	return nil
}

func (s *Studio) importTypes(studioFolder, translationFolder string) error {

	typeIterator, err := s.TypeImporter.Import(studioFolder, translationFolder)
	if err != nil {
		return err
	}
	for pokemonType := range typeIterator {
		s.TypeStore.Add(pokemonType)
	}
	return nil
}

func (s *Studio) importPokemon(studioFolder, translationFolder string) error {

	pokemonIterator, err := s.PokemonImporter.Import(studioFolder, translationFolder)
	if err != nil {
		return err
	}
	for pokemon := range pokemonIterator {
		s.PokemonStore.Add(pokemon)
	}
	return nil
}
