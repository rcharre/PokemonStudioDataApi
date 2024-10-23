package psapi

import (
	"github.com/rcharre/psapi/pkg/ps"
	"github.com/rcharre/psapi/pkg/psapi/psapigen"
	"log/slog"
)

type TypeMapper interface {
	ToTypeDetail(t *ps.PokemonType, lang string) *psapigen.TypeDetails
	ToTypePartial(t *ps.PokemonType, lang string) *psapigen.TypePartial
}

type TypeMapperImpl struct {
}

func NewTypeMapper() *TypeMapperImpl {
	return &TypeMapperImpl{}
}

func (t TypeMapperImpl) ToTypeDetail(pokemonType *ps.PokemonType, lang string) *psapigen.TypeDetails {
	slog.Debug("Mapping type to details")
	typeDamage := make([]psapigen.TypeDamage, len(pokemonType.DamageTo))
	for i, damage := range pokemonType.DamageTo {
		typeDamage[i] = psapigen.TypeDamage{
			DefensiveType: damage.DefensiveType,
			Factor:        &damage.Factor,
		}
		typeDamage[i].Factor = &damage.Factor
	}
	return &psapigen.TypeDetails{
		Symbol:     pokemonType.DbSymbol,
		Name:       pokemonType.Name[lang],
		Color:      pokemonType.Color,
		TypeDamage: typeDamage,
	}
}

func (t TypeMapperImpl) ToTypePartial(pokemonType *ps.PokemonType, lang string) *psapigen.TypePartial {
	slog.Debug("Mapping type to partial")
	return &psapigen.TypePartial{
		Symbol: pokemonType.DbSymbol,
		Name:   pokemonType.Name[lang],
		Color:  pokemonType.Color,
	}
}
