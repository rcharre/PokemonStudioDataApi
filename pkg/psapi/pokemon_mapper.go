package psapi

import (
	"log/slog"
	"psapi/pkg/ps"
	"psapi/pkg/psapi/psapigen"
)

type PokemonMapper interface {
	PokemonToThumbnail(p *ps.Pokemon, lang string) *psapigen.PokemonThumbnail
	PokemonToDetail(p *ps.Pokemon, lang string) *psapigen.PokemonDetail
	FormToPokemonForm(p *ps.PokemonForm, lang string) *psapigen.FormDetail
}
type PokemonMapperImpl struct {
}

func NewPokemonMapper() PokemonMapper {
	return &PokemonMapperImpl{}
}

func (m *PokemonMapperImpl) PokemonToThumbnail(p *ps.Pokemon, lang string) *psapigen.PokemonThumbnail {
	return &psapigen.PokemonThumbnail{
		Symbol: p.DbSymbol,
		Number: p.Id,
		Image:  p.Forms[0].Resources.Front,
		Name:   p.Forms[0].Name[lang],
	}
}

func (m *PokemonMapperImpl) PokemonToDetail(p *ps.Pokemon, lang string) *psapigen.PokemonDetail {
	return &psapigen.PokemonDetail{
		Symbol:   p.DbSymbol,
		Number:   p.Id,
		MainForm: *m.FormToPokemonForm(p.Forms[0], lang),
	}
}

func (m *PokemonMapperImpl) FormToPokemonForm(f *ps.PokemonForm, lang string) *psapigen.FormDetail {

	var breedGroups []string
	for _, breedGroup := range f.BreedGroups {
		breedGroups = append(breedGroups, ps.BreedMap[breedGroup])
	}

	slog.Info("Description", "description", f.Description[lang])

	return &psapigen.FormDetail{
		Form:        &f.Form,
		Name:        f.Name[lang],
		Description: f.Description[lang],

		Height: f.Height,
		Weight: f.Weight,

		Type1: f.Type1,
		Type2: &f.Type2,

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
