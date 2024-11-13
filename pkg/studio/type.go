package studio

import (
	"encoding/json"
	"log/slog"
	"path"

	"github.com/rcharre/psapi/pkg/pkmn"
	"github.com/rcharre/psapi/pkg/utils/file"
	"github.com/rcharre/psapi/pkg/utils/i18n"
)

const (
	TypeFolder                     = "types"
	PokemonTypeTranslationFileName = "100003.csv"
)

// ImportTypes import a type folder to a store
// studioFolder the pokemon studio folder
// translationFolder the translation folder
// store the store the import send data to
func ImportTypes(studioFolder, translationFolder string, store pkmn.Store) error {
	slog.Info("Importing translation file for types")
	typeNameFilePath := path.Join(translationFolder, PokemonTypeTranslationFileName)
	typeNameTranslations, err := ImportTranslations(typeNameFilePath)
	if err != nil {
		return err
	}

	typeFolderPath := path.Join(studioFolder, TypeFolder)
	slog.Info("Importing pokemon types folder", "path", typeFolderPath)
	typeFileIterator, err := file.ImportFolder(typeFolderPath)
	if err != nil {
		return err
	}

	typeStore := store.GetTypeStore()
	for typeFile := range typeFileIterator {
		pokemonType, err := UnmarshalType(typeFile.Content)
		if err != nil {
			slog.Warn("Failed to unmarshal type content", "file", typeFile.Path)
			continue
		}
		TranslateType(pokemonType, typeNameTranslations)
		typeStore.Add(*pokemonType)
	}
	return nil
}

// UnmarshalType unmarshal a json encoded pokemon type to an object
// typeContent the json encoded type
func UnmarshalType(typeContent []byte) (*pkmn.PokemonType, error) {
	pokemonType := &pkmn.PokemonType{}
	if err := json.Unmarshal(typeContent, pokemonType); err != nil {
		return nil, err
	}
	return pokemonType, nil
}

// TranslateType translate a pokemon type
// pokemonType the pokemon type to translate
// typeTranslations the data structure containing all pokemon type name translations
func TranslateType(pokemonType *pkmn.PokemonType, typeTranslations []i18n.Translation) {
	if pokemonType.TextId < len(typeTranslations) {
		pokemonType.Name = typeTranslations[pokemonType.TextId]
	}
}
