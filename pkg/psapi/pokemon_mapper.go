package psapi

import (
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
	}
}

func (m *PokemonMapperImpl) PokemonToDetail(p *ps.Pokemon, lang string) *psapigen.PokemonDetail {
	return &psapigen.PokemonDetail{
		Symbol:   p.DbSymbol,
		Number:   p.Id,
		MainForm: *m.FormToPokemonForm(p.Forms[0], lang),
	}
}

func (m *PokemonMapperImpl) FormToPokemonForm(p *ps.PokemonForm, lang string) *psapigen.FormDetail {

	var breedGroups []string
	for _, breedGroup := range p.BreedGroups {
		breedGroups = append(breedGroups, ps.BreedMap[breedGroup])
	}

	return &psapigen.FormDetail{
		Form: &p.Form,

		Height: p.Height,
		Weight: p.Weight,

		Type1: p.Type1,
		Type2: &p.Type2,

		BaseHp:  p.BaseHp,
		BaseAtk: p.BaseAtk,
		BaseDfe: p.BaseDfe,
		BaseSpd: p.BaseSpd,
		BaseAts: p.BaseAts,
		BaseDfs: p.BaseDfs,

		EvHp:  &p.EvHp,
		EvAtk: &p.EvAtk,
		EvDfe: &p.EvDfe,
		EvSpd: &p.EvSpd,
		EvAts: &p.EvAts,
		EvDfs: &p.EvDfs,

		ExperienceType: ps.ExperienceTypeMap[p.ExperienceType],
		BaseExperience: p.BaseExperience,
		BaseLoyalty:    p.BaseLoyalty,
		CatchRate:      p.CatchRate,
		FemaleRate:     p.FemaleRate,
		BreedGroups:    breedGroups,
		HatchSteps:     p.HatchSteps,
		BabyDbSymbol:   p.BabyDbSymbol,
		BabyForm:       &p.BabyForm,
	}

}
