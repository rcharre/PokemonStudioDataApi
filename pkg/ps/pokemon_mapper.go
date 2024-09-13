package ps

import (
	"psapi/pkg/api"
)

type PokemonMapper interface {
	PokemonToThumbnail(p *Pokemon, lang string) *api.PokemonThumbnail
	PokemonToDetail(p *Pokemon, lang string) *api.PokemonDetail
	FormToPokemonForm(p *PokemonForm, lang string) *api.FormDetail
}
type PokemonMapperImpl struct {
}

func NewPokemonMapper() PokemonMapper {
	return &PokemonMapperImpl{}
}

func (m *PokemonMapperImpl) PokemonToThumbnail(p *Pokemon, lang string) *api.PokemonThumbnail {
	return &api.PokemonThumbnail{
		Symbol: p.DbSymbol,
		Number: p.Id,
		Image:  p.Forms[0].Resources.Front,
	}
}

func (m *PokemonMapperImpl) PokemonToDetail(p *Pokemon, lang string) *api.PokemonDetail {
	return &api.PokemonDetail{
		Symbol:   p.DbSymbol,
		Number:   p.Id,
		MainForm: *m.FormToPokemonForm(p.Forms[0], lang),
	}
}

func (m *PokemonMapperImpl) FormToPokemonForm(p *PokemonForm, lang string) *api.FormDetail {

	var breedGroups []string
	for _, breedGroup := range p.BreedGroups {
		breedGroups = append(breedGroups, BreedMap[breedGroup])
	}

	return &api.FormDetail{
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

		ExperienceType: ExperienceTypeMap[p.ExperienceType],
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
