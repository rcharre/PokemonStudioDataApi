package ps

import (
	"encoding/json"
	"errors"
	"fmt"
	"iter"
	"log/slog"
	"os"
	"path"
	"psapi/pkg/utils/validation"
)

type PokemonImporter interface {
	ImportPokemonFolder(folder string) (iter.Seq[*Pokemon], error)
	ImportPokemonFile(file string) (*Pokemon, error)
}

type PokemonImporterImpl struct {
	pokemonValidator PokemonValidator
}

func NewPokemonImporter(pokemonValidator PokemonValidator) PokemonImporter {
	return &PokemonImporterImpl{
		pokemonValidator: pokemonValidator,
	}
}

// ImportPokemonFolder Import all pokemon files from a given folder path.
func (i *PokemonImporterImpl) ImportPokemonFolder(folder string) (iter.Seq[*Pokemon], error) {
	slog.Info("Importing pokemon folder", "path", folder)

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

	return func(yield func(*Pokemon) bool) {
		for _, file := range files {
			pokemonPath := path.Join(folder, file.Name())
			pokemon, err := i.ImportPokemonFile(pokemonPath)
			if err == nil {
				if !yield(pokemon) {
					break
				}
			}
		}
	}, nil
}

// ImportPokemonFile Import a pokemon from a given file path.
func (i *PokemonImporterImpl) ImportPokemonFile(file string) (*Pokemon, error) {
	info, err := os.Stat(file)
	if err != nil {
		return nil, err
	}

	if info.IsDir() {
		message := fmt.Sprint("Given path ", file, " should be a file")
		return nil, errors.New(message)
	}

	slog.Debug("Importing pokemon", "path", file)

	content, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	pokemon := &Pokemon{}
	err = json.Unmarshal(content, pokemon)

	if err != nil {
		return nil, err
	}

	validations := i.pokemonValidator.Validate(pokemon)
	if len(validations) > 0 {
		return nil, validation.NewValidationError(validations)
	}

	return pokemon, nil
}
