package studio

import (
	"encoding/json"
	"log/slog"
	"path"

	"github.com/rcharre/psapi/pkg/pkmn"
	"github.com/rcharre/psapi/pkg/utils/i18n"
	"github.com/rcharre/psapi/pkg/utils/importer"
)

const (
	StudioFolder   = "Studio"
	LanguageFolder = "Text/Dialogs"

	TypeFolder                     = "types"
	PokemonTypeTranslationFileName = "100003.csv"

	PokemonFolder                         = "pokemon"
	PokemonTranslationFileName            = "100067.csv"
	PokemonDescriptionTranslationFileName = "100068.csv"

	UndefType = "__undef__"
)

func Import(folder string, store pkmn.Store) error {
	translationFolder := path.Join(folder, LanguageFolder)
	studioFolder := path.Join(folder, StudioFolder)

	if err := importTypes(studioFolder, translationFolder, store.GetTypeStore()); err != nil {
		return err
	}
	if err := importPokemon(studioFolder, translationFolder, store.GetPokemonStore()); err != nil {
		return err
	}
	return nil
}

func importTypes(studioFolder, translationFolder string, typeStore pkmn.TypeStore) error {
	slog.Info("Import translation file for types")
	typeTranslationFilePath := path.Join(translationFolder, PokemonTypeTranslationFileName)
	typeTranslations, err := i18n.ImportTranslations(typeTranslationFilePath)
	if err != nil {
		return err
	}

	typeFolderPath := path.Join(studioFolder, TypeFolder)
	slog.Info("Importing pokemon types folder", "path", typeFolderPath)

	typeContentIterator, err := importer.ImportFolder(typeFolderPath)
	if err != nil {
		return err
	}

	for typeContent := range typeContentIterator {
		pokemonType := &pkmn.PokemonType{}
		if err := json.Unmarshal(typeContent, pokemonType); err != nil {
			slog.Warn("Failed to unmarshal pokemon type", "error", err)
			continue
		}

		if pokemonType.TextId < len(typeTranslations) {
			pokemonType.Name = typeTranslations[pokemonType.TextId]
		}

		typeStore.Add(pokemonType)
	}
	return nil
}

func importPokemon(studioFolder, translationFolder string, pokemonStore pkmn.PokemonStore) error {
	slog.Info("Import translation files for pokemon")
	pokemonTranslationFilePath := path.Join(translationFolder, PokemonTranslationFileName)
	pokemonNameTranslations, err := i18n.ImportTranslations(pokemonTranslationFilePath)
	if err != nil {
		return err
	}

	pokemonDescriptionFilePath := path.Join(translationFolder, PokemonDescriptionTranslationFileName)
	pokemonDescriptionTranslations, err := i18n.ImportTranslations(pokemonDescriptionFilePath)
	if err != nil {
		return err
	}

	pokemonFolderPath := path.Join(studioFolder, PokemonFolder)
	slog.Info("Importing pokemon folder", "path", pokemonFolderPath)
	pokemonContentIterator, err := importer.ImportFolder(pokemonFolderPath)
	if err != nil {
		return err
	}

	for pokemonContent := range pokemonContentIterator {
		pokemon := &pkmn.Pokemon{}
		if err = json.Unmarshal(pokemonContent, pokemon); err != nil {
			slog.Warn("Failed to unmarshal pokemon", "error", err)
			continue
		}

		translatePokemon(&pokemonNameTranslations, &pokemonDescriptionTranslations, pokemon)
		pokemonStore.Add(pokemon)
	}
	return nil
}

func translatePokemon(nameTranslations *[]i18n.Translation, descriptionTranslation *[]i18n.Translation, pokemon *pkmn.Pokemon) {
	pokemonNameTranslations := *nameTranslations
	pokemonDescriptionTranslations := *descriptionTranslation

	nameTranslationSize := len(pokemonNameTranslations)
	descriptionTranslationSize := len(pokemonDescriptionTranslations)

	for _, form := range pokemon.Forms {
		if form.Type2 == UndefType {
			form.Type2 = ""
		}

		if form.FormTextId.Name < nameTranslationSize {
			form.Name = pokemonNameTranslations[form.FormTextId.Name]
		}

		if form.FormTextId.Description < descriptionTranslationSize {
			slog.Debug("Form description", "description", form.FormTextId.Description)
			form.Description = pokemonDescriptionTranslations[form.FormTextId.Description]
		}
	}
}
