package studio_test

import (
	"os"
	"testing"

	"github.com/rcharre/psapi/pkg/studio"
	"github.com/rcharre/psapi/pkg/utils/i18n"
	"github.com/rcharre/psapi/test"
)

func TestUnmarshalPokemon_Error(t *testing.T) {
	content, err := os.ReadFile(test.PokemonInvalid)
	if err != nil {
		t.Error("Error when reading test file", "file", test.PokemonInvalid)
	}
	_, err = studio.UnmarshalPokemon(content)
	if err == nil {
		t.Error("Unmarshal invalid pokemon should return error")
	}
}

func TestUnmarshalPokemon_Ok(t *testing.T) {
	content, err := os.ReadFile(test.PokemonValid)
	if err != nil {
		t.Error("Error when reading test file", "file", test.PokemonValid)
	}

	pokemon, err := studio.UnmarshalPokemon(content)
	if err != nil {
		t.Error("Unmarshal valid pokemon should not return error")
	}

	form := pokemon.Forms[0]
	if form.Type2 != nil {
		t.Error("Unmarshal undefined type2 should set type to nil, has", *form.Type2)
	}
}

func TestTranslatePokemon_NameOob(t *testing.T) {
	pokemon := studio.Pokemon{
		Forms: []studio.PokemonForm{
			{
				FormTextId: studio.FormTextId{
					Name:        1000,
					Description: 1000,
				},
			},
		},
	}

	form := pokemon.Forms[0]
	translation := []i18n.Translation{
		{"en": "test"},
	}

	studio.TranslatePokemon(&pokemon, translation, translation)
	if form.Name["en"] != "" {
		t.Error("Translation for pokemon name should be empty")
	}
	if form.Description["en"] != "" {
		t.Error("Translation for pokemon description should be empty")
	}

}

func TestTranslatePokemon_Ok(t *testing.T) {
	pokemon := studio.Pokemon{
		Forms: []studio.PokemonForm{
			{
				FormTextId: studio.FormTextId{
					Name:        0,
					Description: 0,
				},
			},
		},
	}

	translation := []i18n.Translation{
		{"en": "test"},
	}

	studio.TranslatePokemon(&pokemon, translation, translation)
	form := &pokemon.Forms[0]
	if form.Name["en"] != "test" {
		t.Error("Translation for pokemon name should not be empty")
	}
	if form.Description["test"] != "" {
		t.Error("Translation for pokemon description should not be empty")
	}

}
