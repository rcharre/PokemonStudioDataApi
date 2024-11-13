package psapi_test

import (
	"testing"

	"github.com/rcharre/psapi/pkg/pkmn"
	"github.com/rcharre/psapi/pkg/psapi"
	"github.com/rcharre/psapi/pkg/storage/inmem"
	"github.com/rcharre/psapi/pkg/utils/i18n"
)

func TestPokemonToThumbnail(t *testing.T) {
	lang := "test"
	pokemon := pkmn.Pokemon{
		Id:       1,
		DbSymbol: "test",
		Forms: []pkmn.PokemonForm{
			{
				Name: i18n.Translation{
					lang: "testName",
				},
				Resources: pkmn.Resources{
					Front: "testFrontImage",
				},
			},
		},
	}

	typeMapper := psapi.NewTypeMapper()
	typeStore := inmem.NewInMemoryTypeStore()
	pokemonMapper := psapi.NewPokemonMapper(typeMapper, typeStore)

	thumbnail := pokemonMapper.PokemonToThumbnail(pokemon, lang)

	if thumbnail.Image != pokemon.Forms[0].Resources.Front {
		t.Error("Mapper should map image, expected", pokemon.Forms[0].Resources.Front, ", has", thumbnail.Image)
	}

	if thumbnail.Name != pokemon.Forms[0].Name[lang] {
		t.Error("Mapper should map name, expected", pokemon.Forms[0].Name, ", has", thumbnail.Name)
	}

	if thumbnail.Symbol != pokemon.DbSymbol {
		t.Error("Mapper should map db symbol, expected", pokemon.DbSymbol, ", has", thumbnail.Symbol)
	}
	if thumbnail.Number != pokemon.Id {
		t.Error("Mapper should map Id, expected", pokemon.Id, ", has", thumbnail.Number)
	}
}
