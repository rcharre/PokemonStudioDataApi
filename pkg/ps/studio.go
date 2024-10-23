package ps

import (
	"path"
)

const (
	StudioFolder   = "Studio"
	LanguageFolder = "Text/Dialogs"
)

type Studio struct {
	TypeStore    TypeStore
	PokemonStore PokemonStore
}

func NewStudio(typeStore TypeStore, pokemonStore PokemonStore) *Studio {
	return &Studio{
		typeStore,
		pokemonStore,
	}
}

func NewInMemoryStudio() *Studio {
	typeStore := NewInMemoryTypeStore()
	pokemonStore := NewInMemoryPokemonStore()
	return &Studio{
		typeStore,
		pokemonStore,
	}
}

func (s *Studio) Import(folder string) error {
	translationFolder := path.Join(folder, LanguageFolder)
	studioFolder := path.Join(folder, StudioFolder)

	typeImporter := NewTypeImporter()
	typeIterator, err := typeImporter.Import(studioFolder, translationFolder)
	if err != nil {
		return err
	}
	for pokemonType := range typeIterator {
		s.TypeStore.Add(pokemonType)
	}

	pokemonImporter := NewPokemonImporter()
	pokemonIterator, err := pokemonImporter.Import(studioFolder, translationFolder)
	if err != nil {
		return err
	}
	for pokemon := range pokemonIterator {
		s.PokemonStore.Add(pokemon)
	}
	return nil
}
