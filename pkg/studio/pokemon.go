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
	PokemonFolder                         = "pokemon/"
	PokemonTranslationFileName            = "100067.csv"
	PokemonDescriptionTranslationFileName = "100068.csv"
)

// ImportPokemon import a pokemon folder to a store
// studioFolder the pokemon studio folder
// translationFolder the translation folder
// store the store the import is sending data to
func ImportPokemon(studioFolder, translationFolder string, store pkmn.Store) error {
	slog.Info("Importing pokemon name translation")
	pokemonNameFilePath := path.Join(translationFolder, PokemonTranslationFileName)
	pokemonNameTranslations, err := ImportTranslations(pokemonNameFilePath)
	if err != nil {
		return err
	}

	slog.Info("Importing pokemon description translation")
	pokemonDescriptionFilePath := path.Join(translationFolder, PokemonDescriptionTranslationFileName)
	pokemonDescriptionTranslations, err := ImportTranslations(pokemonDescriptionFilePath)
	if err != nil {
		return err
	}

	pokemonFolderPath := path.Join(studioFolder, PokemonFolder)
	slog.Info("Importing pokemon folder", "path", pokemonFolderPath)
	pokemonFileIterator, err := file.ImportFolder(pokemonFolderPath)
	if err != nil {
		return err
	}

	pokemonStore := store.GetPokemonStore()
	for pokemonFile := range pokemonFileIterator {
		pokemon, err := UnmarshalPokemon(pokemonFile.Content)
		if err != nil {
			slog.Warn("Failed to unmarshal pokemon content", "file", pokemonFile.Path)
			continue
		}
		TranslatePokemon(pokemon, pokemonNameTranslations, pokemonDescriptionTranslations)
		pokemonStore.Add(pokemon)
	}
	return nil
}

// UnmarshalPokemon unmarshal a json encoded pokemon to an object
// pokemonContent the encoded pokemon
func UnmarshalPokemon(pokemonContent []byte) (*pkmn.Pokemon, error) {
	pokemon := &pkmn.Pokemon{}
	if err := json.Unmarshal(pokemonContent, pokemon); err != nil {
		return nil, err
	}
	for _, form := range pokemon.Forms {
		if form.Type2 != nil && *form.Type2 == UndefType {
			form.Type2 = nil
		}
	}
	return pokemon, nil
}

// TranslatePokemon add a translation to a pokemon name and description
// pokemon the pokemon to add translation to
// pokemonNameTranslations the datastucture containing all pokemon names translations
// pokemonDescriptionTranslations the datastructure containing all pokemon description translations
func TranslatePokemon(pokemon *pkmn.Pokemon, pokemonNameTranslations, pokemonDescriptionTranslations []i18n.Translation) {
	nameTranslationSize := len(pokemonNameTranslations)
	descriptionTranslationSize := len(pokemonDescriptionTranslations)

	for _, form := range pokemon.Forms {
		if form.FormTextId.Name < nameTranslationSize {
			form.Name = pokemonNameTranslations[form.FormTextId.Name]
		}

		if form.FormTextId.Description < descriptionTranslationSize {
			form.Description = pokemonDescriptionTranslations[form.FormTextId.Description]
		}
	}

}
