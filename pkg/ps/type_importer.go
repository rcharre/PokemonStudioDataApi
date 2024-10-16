package ps

import (
	"encoding/json"
	"log/slog"
	"path"
	"psapi/pkg/utils/i18n"
	"psapi/pkg/utils/importer"
)

const (
	TypeFolder                     = "types"
	PokemonTypeTranslationFileName = "100003.csv"
)

type TypeImporter interface {
	Import(studioFolder string, translationFolder string) (TypeStore, error)
}

type TypeImporterImpl struct {
}

func NewTypeImporter() *TypeImporterImpl {
	return &TypeImporterImpl{}
}

// Import import a types folder.
func (i *TypeImporterImpl) Import(studioFolder string, translationFolder string) ([]*PokemonType, error) {
	slog.Info("Import translation file for types")
	typeTranslationFilePath := path.Join(translationFolder, PokemonTypeTranslationFileName)
	typeTranslations, err := i18n.ImportTranslations(typeTranslationFilePath)
	if err != nil {
		return nil, err
	}

	typeFolderPath := path.Join(studioFolder, TypeFolder)
	slog.Info("Importing pokemon types folder", "path", typeFolderPath)

	typeContentIterator, err := importer.ImportFolder(typeFolderPath)
	if err != nil {
		return nil, err
	}

	results := make([]*PokemonType, 0)
	for typeContent := range typeContentIterator {
		pokemonType := &PokemonType{}
		if err := json.Unmarshal(typeContent, pokemonType); err != nil {
			slog.Warn("Failed to unmarshal pokemon type", "error", err)
			continue
		}

		if pokemonType.TextId < len(typeTranslations) {
			pokemonType.Name = typeTranslations[pokemonType.TextId]
		}

		results = append(results, pokemonType)
	}

	return results, nil
}
