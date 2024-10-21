package ps

import (
	"path"
	"slices"
)

const (
	StudioFolder   = "Studio"
	LanguageFolder = "Text/Dialogs"
)

type Studio struct {
	typeStore    TypeStore
	pokemonStore PokemonStore
}

func NewStudio(typeStore TypeStore, pokemonStore PokemonStore) *Studio {
	return &Studio{
		typeStore,
		pokemonStore,
	}
}

func NewInMemoryStudio(folder string) (*Studio, error) {
	translationFolder := path.Join(folder, LanguageFolder)
	studioFolder := path.Join(folder, StudioFolder)

	typeImporter := NewTypeImporter()
	typeIterator, err := typeImporter.Import(studioFolder, translationFolder)
	if err != nil {
		return nil, err
	}

	typeList := slices.Collect(typeIterator)
	typeStore := NewInMemoryTypeStore(typeList)

	pokemonImporter := NewPokemonImporter()
	pokemonIterator, err := pokemonImporter.Import(studioFolder, translationFolder)
	if err != nil {
		return nil, err
	}

	pokemonList := slices.Collect(pokemonIterator)
	pokemonStore := NewInMemoryPokemonStore(pokemonList)

	return &Studio{
		typeStore:    typeStore,
		pokemonStore: pokemonStore,
	}, nil
}

func (a *Studio) TypeStore() TypeStore {
	return a.typeStore
}

func (a *Studio) PokemonStore() PokemonStore {
	return a.pokemonStore
}
