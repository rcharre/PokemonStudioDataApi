package psapi_test

import (
	"testing"

	"github.com/rcharre/psapi/pkg/ps"
	"github.com/rcharre/psapi/pkg/psapi"
	"github.com/rcharre/psapi/pkg/utils/i18n"
)

func TestToTypeDetail(t *testing.T) {
	lang := "test"
	pokemonType := &ps.PokemonType{
		DbSymbol: "testDbSymbol",
		Color:    "testColor",
		Name: i18n.Translation{
			lang: "testName",
		},
		DamageTo: []ps.TypeDamage{{
			DefensiveType: "testDefType",
			Factor:        .2,
		}},
	}

	typeMapper := psapi.NewTypeMapper()
	typeDetail := typeMapper.ToTypeDetail(pokemonType, lang)

	if typeDetail.Name != pokemonType.Name[lang] {
		t.Error("Mapper should map name, expected", pokemonType.Name[lang], ", get", typeDetail.Name)
	}

	if typeDetail.Color != pokemonType.Color {
		t.Error("Mapper should map color, expected", pokemonType.Color, ", get", typeDetail.Color)
	}

	if typeDetail.Symbol != pokemonType.DbSymbol {
		t.Error("Mapper should map db symbol, expected", pokemonType.DbSymbol, ", get", typeDetail.Symbol)
	}

	for i, typeDamage := range pokemonType.DamageTo {
		result := typeDetail.TypeDamage[i]

		if typeDamage.DefensiveType != result.DefensiveType {
			t.Error("Mapper should map defensive type, expected", typeDamage.DefensiveType, ", get", result.DefensiveType)
		}

		if typeDamage.Factor != *result.Factor {
			t.Error("Mapper should map factor damage, expected", typeDamage.Factor, ", get", result.Factor)
		}
	}
}

func TestToTypePartial(t *testing.T) {
	lang := "test"
	pokemonType := &ps.PokemonType{
		DbSymbol: "testDbSymbol",
		Color:    "testColor",
		Name: i18n.Translation{
			lang: "testName",
		},
	}

	typeMapper := psapi.NewTypeMapper()
	typePartial := typeMapper.ToTypePartial(pokemonType, lang)

	if typePartial.Name != pokemonType.Name[lang] {
		t.Error("Mapper should map name, expected", pokemonType.Name[lang], ", get", typePartial.Name)
	}

	if typePartial.Color != pokemonType.Color {
		t.Error("Mapper should map color, expected", pokemonType.Color, ", get", typePartial.Color)
	}

	if typePartial.Symbol != pokemonType.DbSymbol {
		t.Error("Mapper should map db symbol, expected", pokemonType.DbSymbol, ", get", typePartial.Symbol)
	}

}
