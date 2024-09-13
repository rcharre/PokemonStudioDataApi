package ps

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path"
	"psapi/pkg/utils/validation"
)

type PokemonImporter interface {
	ImportFolder(folder string) error
	ImportFile(file string) error
}

type PokemonImporterImpl struct {
	pokemonStore     PokemonStore
	pokemonValidator PokemonValidator
}

func NewPokemonImporter(store PokemonStore, validator PokemonValidator) PokemonImporter {
	return &PokemonImporterImpl{
		store,
		validator,
	}
}

// ImportFolder Import all pokemon files from a given folder path.
func (i *PokemonImporterImpl) ImportFolder(folder string) error {
	slog.Info("Importing pokemon folder", "path", folder)
	info, err := os.Stat(folder)
	if err != nil {
		return err
	}

	if !info.IsDir() {
		message := fmt.Sprint("Given path : ", folder, " should be a directory ")
		return errors.New(message)
	}

	files, err := os.ReadDir(folder)
	if err != nil {
		return err
	}

	for _, file := range files {
		pokemonPath := path.Join(folder, file.Name())
		err := i.ImportFile(pokemonPath)
		if err != nil {
			slog.Error("Failed to import pokemon", "path", pokemonPath, "error", err.Error())
		}
	}

	return nil
}

// ImportFile Import a pokemon from a given file path.
func (i *PokemonImporterImpl) ImportFile(file string) error {
	info, err := os.Stat(file)
	if err != nil {
		return err
	}

	if info.IsDir() {
		message := fmt.Sprint("Given path ", file, " should be a file")
		return errors.New(message)
	}

	slog.Debug("Importing pokemon", "path", file)

	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	pokemon := &Pokemon{}
	err = json.Unmarshal(content, pokemon)

	if err != nil {
		return err
	}

	validations := i.pokemonValidator.Validate(pokemon)
	if len(validations) > 0 {
		return validation.NewValidationError(validations)
	}

	i.pokemonStore.Add(pokemon)
	return nil
}
