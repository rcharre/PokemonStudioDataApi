package studio_test

import (
	"os"
	"testing"

	"github.com/rcharre/psapi/pkg/pkmn"
	"github.com/rcharre/psapi/pkg/studio"
	"github.com/rcharre/psapi/pkg/utils/i18n"
	"github.com/rcharre/psapi/test"
)

func TestUnmarshalType_Error(t *testing.T) {
	content, err := os.ReadFile(test.TypeInvalid)
	if err != nil {
		t.Error("Error when reading test file", "file", test.TypeInvalid)
	}
	_, err = studio.UnmarshalType(content)
	if err == nil {
		t.Error("Unmarshal invalid type should return error")
	}
}

func TestUnmarshalType_Ok(t *testing.T) {
	content, err := os.ReadFile(test.TypeValid)
	if err != nil {
		t.Error("Error when reading test file", "file", test.TypeValid)
	}

	_, err = studio.UnmarshalPokemon(content)
	if err != nil {
		t.Error("Unmarshal valid type should not return error")
	}
}

func TestTranslateType_Oob(t *testing.T) {
	pokemonType := pkmn.PokemonType{
		TextId: 5,
	}

	translations := []i18n.Translation{
		{"en": "test"},
	}

	studio.TranslateType(&pokemonType, translations)

	if pokemonType.Name["en"] != "" {
		t.Error("Translation for pokemon name should be empty")
	}
}

func TestTranslateType_Ok(t *testing.T) {
	pokemonType := pkmn.PokemonType{
		TextId: 0,
	}

	translations := []i18n.Translation{
		{"en": "test"},
	}

	studio.TranslateType(&pokemonType, translations)

	if pokemonType.Name["en"] != "test" {
		t.Error("Translation for pokemon name should not be empty")
	}

}
