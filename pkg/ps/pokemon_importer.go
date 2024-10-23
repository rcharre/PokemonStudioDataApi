package ps

import (
	"encoding/json"
	"github.com/rcharre/psapi/pkg/utils/i18n"
	"github.com/rcharre/psapi/pkg/utils/importer"
	"iter"
	"log/slog"
	"path"
)

const (
	PokemonFolder                         = "pokemon"
	PokemonTranslationFileName            = "100067.csv"
	PokemonDescriptionTranslationFileName = "100068.csv"

	UndefType = "__undef__"
)

var _ PokemonImporter = &PokemonImporterImpl{}

type PokemonImporter interface {
	Import(studioFolder string, translationFolder string) (iter.Seq[*Pokemon], error)
}

type PokemonImporterImpl struct {
}

func NewPokemonImporter() *PokemonImporterImpl {
	return &PokemonImporterImpl{}
}

// Import import all pokemon files from a given folder path.
func (i PokemonImporterImpl) Import(studioFolder string, translationFolder string) (iter.Seq[*Pokemon], error) {
	slog.Info("Import translation files for pokemon")
	pokemonTranslationFilePath := path.Join(translationFolder, PokemonTranslationFileName)
	pokemonNameTranslations, err := i18n.ImportTranslations(pokemonTranslationFilePath)
	if err != nil {
		return nil, err
	}

	pokemonDescriptionFilePath := path.Join(translationFolder, PokemonDescriptionTranslationFileName)
	pokemonDescriptionTranslations, err := i18n.ImportTranslations(pokemonDescriptionFilePath)
	if err != nil {
		return nil, err
	}

	pokemonFolderPath := path.Join(studioFolder, PokemonFolder)
	slog.Info("Importing pokemon folder", "path", pokemonFolderPath)
	pokemonContentIterator, err := importer.ImportFolder(pokemonFolderPath)
	if err != nil {
		return nil, err
	}

	return func(yield func(*Pokemon) bool) {
		for pokemonContent := range pokemonContentIterator {
			pokemon := &Pokemon{}
			if err = json.Unmarshal(pokemonContent, pokemon); err != nil {
				slog.Warn("Failed to unmarshal pokemon", "error", err)
				continue
			}

			applyTranslation(&pokemonNameTranslations, &pokemonDescriptionTranslations, pokemon)
			if !yield(pokemon) {
				break
			}
		}
	}, nil
}

func applyTranslation(nameTranslations *[]i18n.Translation, descriptionTranslation *[]i18n.Translation, pokemon *Pokemon) {
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
