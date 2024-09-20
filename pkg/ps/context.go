package ps

import (
	"log/slog"
	"path"
)

const (
	StudioFolder = "Studio"

	PokemonFolder = "pokemon"
	TypeFolder    = "type"
	AbilityFolder = "ability"
	MoveFolder    = "move"
	TrainerFolder = "trainer"

	LanguageFolder = "language"
)

type AppContext struct {
	pokemonStore     PokemonStore
	pokemonImporter  PokemonImporter
	pokemonValidator PokemonValidator
}

func NewAppContext(pokemonStore PokemonStore, pokemonImporter PokemonImporter, pokemonValidator PokemonValidator) *AppContext {
	return &AppContext{
		pokemonStore,
		pokemonImporter,
		pokemonValidator,
	}
}

func NewDefaultAppContext() *AppContext {
	pokemonStore := NewPokemonStore()
	pokemonValidator := NewPokemonValidator()
	pokemonImporter := NewPokemonImporter(pokemonStore, pokemonValidator)

	return &AppContext{
		pokemonStore,
		pokemonImporter,
		pokemonValidator,
	}
}

func (c *AppContext) PokemonStore() PokemonStore {
	return c.pokemonStore
}

func (c *AppContext) PokemonImporter() PokemonImporter {
	return c.pokemonImporter
}

func (c *AppContext) PokemonValidator() PokemonValidator {
	return c.pokemonValidator
}

func (c *AppContext) ImportPokemonStudioFolder(folder string) error {
	slog.Info("Importing pokemon studio folder", "path", folder)
	pokemonFolderPath := path.Join(folder, StudioFolder, PokemonFolder)
	if err := c.PokemonImporter().ImportFolder(pokemonFolderPath); err != nil {
		return err
	}

	return nil
}
