package ps

import (
	"encoding/csv"
	"io"
	"log/slog"
	"os"
)

type translation map[string]string

// ImportPokemonTranslations import pokemon translations from file
func ImportTranslations(path string) ([]translation, error) {
	slog.Info("Import translation file", "path", path)
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	langs, err := reader.Read()
	if err != nil {
		return nil, err
	}

	results := make([]translation, 0)
	for {
		records, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		translations := make(map[string]string)
		for index := range len(records) {
			translations[langs[index]] = records[index]
		}

		results = append(results, translations)
	}

	return results, nil
}
