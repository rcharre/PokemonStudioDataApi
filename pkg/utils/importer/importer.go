package importer

import (
	"errors"
	"fmt"
	"iter"
	"log/slog"
	"os"
	"path"
)

// ImportFolder Import all files from a given folder path.
func ImportFolder(folder string) (iter.Seq[[]byte], error) {
	slog.Info("Importing folder", "path", folder)

	info, err := os.Stat(folder)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		message := fmt.Sprint("Given path : ", folder, " should be a directory ")
		return nil, errors.New(message)
	}

	files, err := os.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	return func(yield func([]byte) bool) {
		for _, file := range files {
			filePath := path.Join(folder, file.Name())
			content, err := ImportFile(filePath)
			if err == nil {
				if !yield(content) {
					break
				}
			}
		}
	}, nil
}

// ImportFile Import file content from a given file path.
func ImportFile(file string) ([]byte, error) {
	info, err := os.Stat(file)
	if err != nil {
		return nil, err
	}

	if info.IsDir() {
		message := fmt.Sprint("Given path ", file, " should be a file")
		return nil, errors.New(message)
	}

	slog.Debug("Importing file", "path", file)
	return os.ReadFile(file)
}
