package psapi_test

import (
	"testing"

	"github.com/rcharre/psapi/pkg/ps"
	"github.com/rcharre/psapi/pkg/psapi"
	"github.com/rcharre/psapi/pkg/utils/i18n"
	"github.com/rcharre/psapi/test"
)

func TestPokemonToThumbnail(t *testing.T) {
	lang := "test"
	pokemon := &ps.Pokemon{
		Id:       1,
		DbSymbol: "test",
		Forms: []*ps.PokemonForm{
			{
				Name: i18n.Translation{
					lang: "testName",
				},
				Resources: &ps.Resources{
					Front: "testFrontImage",
				},
			},
		},
	}

	typeMapper := &test.TypeMapperMock{}
	typeStore := &test.TypeStoreMock{}
	pokemonMapper := psapi.NewPokemonMapper(typeMapper, typeStore)

	thumbnail := pokemonMapper.PokemonToThumbnail(pokemon, lang)

	if thumbnail.Image != pokemon.Forms[0].Resources.Front {
		t.Error("Mapper should map image, expected", pokemon.Forms[0].Resources.Front, ", get", thumbnail.Image)
	}

	if thumbnail.Name != pokemon.Forms[0].Name[lang] {
		t.Error("Mapper should map name, expected", pokemon.Forms[0].Name, ", get", thumbnail.Name)
	}

	if thumbnail.Symbol != pokemon.DbSymbol {
		t.Error("Mapper should map db symbol, expected", pokemon.DbSymbol, ", get", thumbnail.Symbol)
	}
	if thumbnail.Number != pokemon.Id {
		t.Error("Mapper should map Id, expected", pokemon.Id, ", get", thumbnail.Number)
	}
}
