package psapi

import (
	"log/slog"

	"github.com/rcharre/psapi/pkg/pkmn"
	"github.com/rcharre/psapi/pkg/psapi/psapigen"
)

type PokemonMapper interface {
	PokemonToThumbnail(p pkmn.Pokemon, lang string) psapigen.PokemonThumbnail
	PokemonToDetail(p pkmn.Pokemon, lang string) psapigen.PokemonDetails
	FormToPokemonFormDetails(f pkmn.PokemonForm, lang string) psapigen.FormDetails
}
type pokemonMapper struct {
	typeMapper TypeMapper
	typeStore  pkmn.TypeStore
}

// NewPokemonMapper Create a new pokemon mapper
// typeMapper the mapper for pokemon types
// typeStore the store for pokemon types
func NewPokemonMapper(typeMapper TypeMapper, typeStore pkmn.TypeStore) PokemonMapper {
	return &pokemonMapper{
		typeMapper,
		typeStore,
	}
}

// PokemonToThumbnail map a pokemon to a thumbnail transfer object
// p the pokemon to map
// lang the language expected
func (m pokemonMapper) PokemonToThumbnail(p pkmn.Pokemon, lang string) psapigen.PokemonThumbnail {
	slog.Debug("Mapping pokemon to thumbnail")
	return psapigen.PokemonThumbnail{
		Symbol: p.DbSymbol,
		Number: p.Id,
		Image:  p.Forms[0].Resources.Front,
		Name:   p.Forms[0].Name[lang],
	}
}

// PokemonToDetail map a pokemon to a details transfer object
// p the pokemon to map
// lang the language expected
func (m pokemonMapper) PokemonToDetail(p pkmn.Pokemon, lang string) psapigen.PokemonDetails {
	slog.Debug("Mapping pokemon to details")
	return psapigen.PokemonDetails{
		Symbol:   p.DbSymbol,
		Number:   p.Id,
		MainForm: m.FormToPokemonFormDetails(p.Forms[0], lang),
	}
}

// FormToPokemonFormDetails map a pokemon form to a form details transfer object
// p the pokemon form to map
// lang the language expected
func (m pokemonMapper) FormToPokemonFormDetails(f pkmn.PokemonForm, lang string) psapigen.FormDetails {
	slog.Debug("Mapping pokemon form to form details")
	var breedGroups []string
	for _, breedGroup := range f.BreedGroups {
		breedGroups = append(breedGroups, pkmn.BreedMap[breedGroup])
	}

	var partialType2Ptr *psapigen.TypePartial

	type1 := m.typeStore.FindBySymbol(f.Type1)
	partialType1 := m.typeMapper.ToTypePartial(*type1, lang)

	if f.Type2 != nil {
		type2 := m.typeStore.FindBySymbol(*f.Type2)
		if type2 != nil {
			partialType2 := m.typeMapper.ToTypePartial(*type2, lang)
			partialType2Ptr = &partialType2
		}
	}

	return psapigen.FormDetails{
		Form:        &f.Form,
		Name:        f.Name[lang],
		Description: f.Description[lang],

		Height: f.Height,
		Weight: f.Weight,

		Type1: &partialType1,
		Type2: partialType2Ptr,

		BaseHp:  f.BaseHp,
		BaseAtk: f.BaseAtk,
		BaseDfe: f.BaseDfe,
		BaseSpd: f.BaseSpd,
		BaseAts: f.BaseAts,
		BaseDfs: f.BaseDfs,

		EvHp:  &f.EvHp,
		EvAtk: &f.EvAtk,
		EvDfe: &f.EvDfe,
		EvSpd: &f.EvSpd,
		EvAts: &f.EvAts,
		EvDfs: &f.EvDfs,

		ExperienceType: pkmn.ExperienceTypeMap[f.ExperienceType],
		BaseExperience: f.BaseExperience,
		BaseLoyalty:    f.BaseLoyalty,
		CatchRate:      f.CatchRate,
		FemaleRate:     f.FemaleRate,
		BreedGroups:    breedGroups,
		HatchSteps:     f.HatchSteps,
		BabyDbSymbol:   f.BabyDbSymbol,
		BabyForm:       &f.BabyForm,
	}
}
