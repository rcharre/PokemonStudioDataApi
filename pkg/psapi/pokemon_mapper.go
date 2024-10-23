package psapi

import (
	"log/slog"
	"psapi/pkg/ps"
	"psapi/pkg/psapi/psapigen"
)

type PokemonMapper interface {
	PokemonToThumbnail(p *ps.Pokemon, lang string) *psapigen.PokemonThumbnail
	PokemonToDetail(p *ps.Pokemon, lang string) *psapigen.PokemonDetails
	FormToPokemonFormDetails(p *ps.PokemonForm, lang string) *psapigen.FormDetails
}
type PokemonMapperImpl struct {
	typeMapper TypeMapper
	typeStore  ps.TypeStore
}

func NewPokemonMapper(typeMapper TypeMapper, typeStore ps.TypeStore) *PokemonMapperImpl {
	return &PokemonMapperImpl{
		typeMapper,
		typeStore,
	}
}

func (m *PokemonMapperImpl) PokemonToThumbnail(p *ps.Pokemon, lang string) *psapigen.PokemonThumbnail {
	slog.Debug("Mapping pokemon to thumbnail")
	return &psapigen.PokemonThumbnail{
		Symbol: p.DbSymbol,
		Number: p.Id,
		Image:  p.Forms[0].Resources.Front,
		Name:   p.Forms[0].Name[lang],
	}
}

func (m *PokemonMapperImpl) PokemonToDetail(p *ps.Pokemon, lang string) *psapigen.PokemonDetails {
	slog.Debug("Mapping pokemon to details")
	return &psapigen.PokemonDetails{
		Symbol:   p.DbSymbol,
		Number:   p.Id,
		MainForm: *m.FormToPokemonFormDetails(p.Forms[0], lang),
	}
}

func (m *PokemonMapperImpl) FormToPokemonFormDetails(f *ps.PokemonForm, lang string) *psapigen.FormDetails {
	slog.Debug("Mapping pokemon form to form details")
	var breedGroups []string
	for _, breedGroup := range f.BreedGroups {
		breedGroups = append(breedGroups, ps.BreedMap[breedGroup])
	}

	var partialType1 *psapigen.TypePartial
	var partialType2 *psapigen.TypePartial

	type1 := m.typeStore.FindBySymbol(f.Type1)
	partialType1 = m.typeMapper.ToTypePartial(type1, lang)

	if f.Type2 != "" {
		type2 := m.typeStore.FindBySymbol(f.Type2)
		partialType2 = m.typeMapper.ToTypePartial(type2, lang)
	}

	return &psapigen.FormDetails{
		Form:        &f.Form,
		Name:        f.Name[lang],
		Description: f.Description[lang],

		Height: f.Height,
		Weight: f.Weight,

		Type1: partialType1,
		Type2: partialType2,

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

		ExperienceType: ps.ExperienceTypeMap[f.ExperienceType],
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
