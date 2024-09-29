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

	LanguageFolder = "Text/Dialogs"

	PokemonTranslationFileName = "100000.csv"
)

type App struct {
	pokemonStore     PokemonStore
	pokemonImporter  PokemonImporter
	pokemonValidator PokemonValidator
}

func NewApp(pokemonStore PokemonStore, pokemonImporter PokemonImporter, pokemonValidator PokemonValidator) *App {
	return &App{
		pokemonStore,
		pokemonImporter,
		pokemonValidator,
	}
}

func NewDefaultApp() *App {
	pokemonStore := NewPokemonStore()
	pokemonValidator := NewPokemonValidator()
	pokemonImporter := NewPokemonImporter(pokemonValidator)

	return &App{
		pokemonStore,
		pokemonImporter,
		pokemonValidator,
	}
}

func (a *App) PokemonStore() PokemonStore {
	return a.pokemonStore
}

func (a *App) PokemonImporter() PokemonImporter {
	return a.pokemonImporter
}

func (a *App) PokemonValidator() PokemonValidator {
	return a.pokemonValidator
}

func (a *App) ImportData(folder string) error {
	slog.Info("Importing pokemon studio folder", "path", folder)
	pokemonFolderPath := path.Join(folder, StudioFolder, PokemonFolder)
	pokemonList, err := a.PokemonImporter().ImportPokemonFolder(pokemonFolderPath)
	if err != nil {
		return err
	}

	slog.Info("Import translation file for pokemon")
	pokemonTranslationFilePath := path.Join(folder, LanguageFolder, PokemonTranslationFileName)
	translations, err := ImportTranslations(pokemonTranslationFilePath)
	a.pokemonStore.SetPokemonList(pokemonList, translations)
	if err != nil {
		return err
	}

	return nil
}
