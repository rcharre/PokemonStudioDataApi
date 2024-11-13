package psapi_test

import (
	"testing"

	"github.com/rcharre/psapi/pkg/pkmn"
	"github.com/rcharre/psapi/pkg/psapi"
	"github.com/rcharre/psapi/pkg/utils/i18n"
)

func TestToTypeDetail(t *testing.T) {
	lang := "test"
	pokemonType := pkmn.PokemonType{
		DbSymbol: "testDbSymbol",
		Color:    "testColor",
		Name: i18n.Translation{
			lang: "testName",
		},
		DamageTo: []pkmn.TypeDamage{{
			DefensiveType: "testDefType",
			Factor:        .2,
		}},
	}

	typeMapper := psapi.NewTypeMapper()
	typeDetail := typeMapper.ToTypeDetail(pokemonType, lang)

	if typeDetail.Name != pokemonType.Name[lang] {
		t.Error("Mapper should map name, expected", pokemonType.Name[lang], ", has", typeDetail.Name)
	}

	if typeDetail.Color != pokemonType.Color {
		t.Error("Mapper should map color, expected", pokemonType.Color, ", has", typeDetail.Color)
	}

	if typeDetail.Symbol != pokemonType.DbSymbol {
		t.Error("Mapper should map db symbol, expected", pokemonType.DbSymbol, ", has", typeDetail.Symbol)
	}

	for i, typeDamage := range pokemonType.DamageTo {
		result := typeDetail.TypeDamage[i]

		if typeDamage.DefensiveType != result.DefensiveType {
			t.Error("Mapper should map defensive type, expected", typeDamage.DefensiveType, ", has", result.DefensiveType)
		}

		if typeDamage.Factor != *result.Factor {
			t.Error("Mapper should map factor damage, expected", typeDamage.Factor, ", has", result.Factor)
		}
	}
}

func TestToTypePartial(t *testing.T) {
	lang := "test"
	pokemonType := pkmn.PokemonType{
		DbSymbol: "testDbSymbol",
		Color:    "testColor",
		Name: i18n.Translation{
			lang: "testName",
		},
	}

	typeMapper := psapi.NewTypeMapper()
	typePartial := typeMapper.ToTypePartial(pokemonType, lang)

	if typePartial.Name != pokemonType.Name[lang] {
		t.Error("Mapper should map name, expected", pokemonType.Name[lang], ", has", typePartial.Name)
	}

	if typePartial.Color != pokemonType.Color {
		t.Error("Mapper should map color, expected", pokemonType.Color, ", has", typePartial.Color)
	}

	if typePartial.Symbol != pokemonType.DbSymbol {
		t.Error("Mapper should map db symbol, expected", pokemonType.DbSymbol, ", has", typePartial.Symbol)
	}

}
