package psapi

import (
	"log/slog"

	"github.com/rcharre/psapi/pkg/pkmn"
	"github.com/rcharre/psapi/pkg/psapi/psapigen"
)

type TypeMapper interface {
	ToTypeDetail(t *pkmn.PokemonType, lang string) *psapigen.TypeDetails
	ToTypePartial(t *pkmn.PokemonType, lang string) *psapigen.TypePartial
}

type TypeMapperImpl struct {
}

// NewTypeMapper create a new pokemon type mapper
func NewTypeMapper() *TypeMapperImpl {
	return &TypeMapperImpl{}
}

// ToTypeDetail map a type to a type details transfer object
// pokemonType the pokemon type to map
// lang the language expected
func (t TypeMapperImpl) ToTypeDetail(pokemonType *pkmn.PokemonType, lang string) *psapigen.TypeDetails {
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

// ToTypePartial map a type to a type partial transfer object
// pokemonType the pokemon type to map
// lang the language expected
func (t TypeMapperImpl) ToTypePartial(pokemonType *pkmn.PokemonType, lang string) *psapigen.TypePartial {
	slog.Debug("Mapping type to partial")
	return &psapigen.TypePartial{
		Symbol: pokemonType.DbSymbol,
		Name:   pokemonType.Name[lang],
		Color:  pokemonType.Color,
	}
}
