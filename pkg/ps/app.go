package ps

import (
	"log/slog"
	"path"
	"psapi/pkg/utils/i18n"
)

const (
	StudioFolder = "Studio"

	PokemonFolder = "pokemon"
	TypeFolder    = "type"
	AbilityFolder = "ability"
	MoveFolder    = "move"
	TrainerFolder = "trainer"

	LanguageFolder = "Text/Dialogs"

	PokemonTranslationFileName            = "100067.csv"
	PokemonDescriptionTranslationFileName = "100068.csv"
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
	slog.Info("Import translation file for pokemon")
	pokemonTranslationFilePath := path.Join(folder, LanguageFolder, PokemonTranslationFileName)
	pokemonNameTranslations, err := i18n.ImportTranslations(pokemonTranslationFilePath)
	if err != nil {
		return err
	}

	pokemonDescriptionFilePath := path.Join(folder, LanguageFolder, PokemonDescriptionTranslationFileName)
	pokemonDescriptionTranslations, err := i18n.ImportTranslations(pokemonDescriptionFilePath)
	if err != nil {
		return err
	}

	slog.Info("Importing pokemon folder", "path", folder)
	pokemonFolderPath := path.Join(folder, StudioFolder, PokemonFolder)

	pokemonIterator, err := a.PokemonImporter().ImportPokemonFolder(pokemonFolderPath)
	if err != nil {
		return err
	}

	nameTranslationSize := len(pokemonNameTranslations)
	descriptionTranslationSize := len(pokemonDescriptionTranslations)
	for pokemon := range pokemonIterator {
		for _, form := range pokemon.Forms {
			if form.FormTextId.Name < nameTranslationSize {
				form.Name = pokemonNameTranslations[form.FormTextId.Name]
			}

			if form.FormTextId.Description < descriptionTranslationSize {
				slog.Debug("Form description", "description", form.FormTextId.Description)
				form.Description = pokemonDescriptionTranslations[form.FormTextId.Description]
			}
		}
		a.PokemonStore().Add(pokemon)
	}

	return nil
}
