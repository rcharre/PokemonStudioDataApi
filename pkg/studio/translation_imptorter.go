package studio

import (
	"encoding/csv"
	"io"
	"log/slog"
	"os"

	"github.com/rcharre/psapi/pkg/utils/i18n"
)

// ImportTranslations import translations from file
// path the path of the file to import
func ImportTranslations(path string) ([]i18n.Translation, error) {
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

	results := make([]i18n.Translation, 0)
	for {
		records, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		translationMap := make(i18n.Translation)
		for index := range len(records) {
			translationMap[langs[index]] = records[index]
		}

		results = append(results, translationMap)
	}

	return results, nil
}
