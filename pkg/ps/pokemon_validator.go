package ps

import (
	"psapi/pkg/utils/validation"
	"strings"
)

const (
	ID_INVALID        = "validations.pokemon.id.invalid"
	DB_SYMBOL_INVALID = "validations.pokemon.db_symbol.invalid"
	FORM_EMPTY        = "validations.pokemon.forms.empty"
	FORM_INVALID      = "validation.pokemon.forms.invalid"

	FORM_TEXT_ID_INVALID    = "validations.pokemon.form_text_id.invalid"
	HEIGHT_INVALID          = "validations.pokemon.height.invalid"
	WEIGHT_INVALID          = "validations.pokemon.weight.invalid"
	TYPE1_INVALID           = "validations.pokemon.type1.invalid"
	TYPE2_INVALID           = "validations.pokemon.type2.invalid"
	BASE_HP_INVALID         = "validations.pokemon.base_hp.invalid"
	BASE_ATK_INVALID        = "validations.pokemon.base_atk.invalid"
	BASE_DFE_INVALID        = "validations.pokemon.base_dfe.invalid"
	BASE_SPD_INVALID        = "validations.pokemon.base_spd.invalid"
	BASE_ATS_INVALID        = "validations.pokemon.base_ats.invalid"
	BASE_DFS_INVALID        = "validations.pokemon.base_dfs.invalid"
	EV_HP_INVALID           = "validations.pokemon.ev_hp.invalid"
	EV_ATK_INVALID          = "validations.pokemon.ev_atk.invalid"
	EV_DFE_INVALID          = "validations.pokemon.ev_dfe.invalid"
	EV_SPD_INVALID          = "validations.pokemon.ev_spd.invalid"
	EV_ATS_INVALID          = "validations.pokemon.ev_ats.invalid"
	EV_DFS_INVALID          = "validations.pokemon.ev_dfs.invalid"
	EXPERIENCE_TYPE_INVALID = "validations.pokemon.experience_type.invalid"
	BASE_EXPERIENCE_INVALID = "validations.pokemon.base_experience.invalid"
	BASE_LOYALTY_INVALID    = "validations.pokemon.base_loyalty.invalid"
	CATCH_RATE_INVALID      = "validations.pokemon.catch_rate.invalid"
	FEMALE_RATE_INVALID     = "validations.pokemon.female_rate.invalid"
	HATCH_STEPS_INVALID     = "validations.pokemon.hatch_steps.invalid"
	BABY_DB_SYMBOL_INVALID  = "validations.pokemon.baby_db_symbol.invalid"
	BABY_FORM_INVALID       = "validations.pokemon.baby_form.invalid"
	FRONT_OFFSET_Y_INVALID  = "validations.pokemon.front_offset_y.invalid"
	ABILITIES_INVALID       = "validations.pokemon.abilities.invalid"
	ITEM_HELD_INVALID       = "validations.pokemon.item_held.invalid"
	EVOLUTIONS_INVALID      = "validations.pokemon.evolutions.invalid"
	BREED_GROUPS_INVALID    = "validations.pokemon.breed_groups.invalid"
	RESOURCES_INVALID       = "validations.pokemon.resources.invalid"

	EVOLUTION_CONDITION_TYPE_INVALID  = "validations.pokemon.evolution_condition_type.invalid"
	EVOLUTION_CONDITION_VALUE_INVALID = "validations.pokemon.evolution_condition_value.invalid"

	ITEM_HELD_DB_SYMBOL_INVALID = "validations.pokemon.item_held_db_symbol.invalid"
	ITEM_HELD_CHANCE_INVALID    = "validations.pokemon.item_held_chance.invalid"
)

type PokemonValidator interface {
	Validate(pokemon *Pokemon) []*validation.Validation
}

type PokemonValidatorImpl struct {
}

func NewPokemonValidator() PokemonValidator {
	return &PokemonValidatorImpl{}
}

func (v *PokemonValidatorImpl) Validate(pokemon *Pokemon) []*validation.Validation {
	var validations []*validation.Validation

	if pokemon.Id <= 0 {
		validations = append(validations, validation.NewValidation(ID_INVALID, pokemon.Id))
	}
	if strings.TrimSpace(pokemon.DbSymbol) == "" {
		validations = append(validations, validation.NewValidation(DB_SYMBOL_INVALID, pokemon.DbSymbol))
	}

	for _, form := range pokemon.Forms {
		if form.Form < 0 {
			validations = append(validations, validation.NewValidation(FORM_INVALID, form.Form))
		}
		if form.Height <= 0 {
			validations = append(validations, validation.NewValidation(HEIGHT_INVALID, form.Height))
		}
		if form.Weight <= 0 {
			validations = append(validations, validation.NewValidation(WEIGHT_INVALID, form.Weight))
		}
		if form.Type1 == "" {
			validations = append(validations, validation.NewValidation(TYPE1_INVALID, form.Type1))
		}
		if form.BaseHp < 0 {
			validations = append(validations, validation.NewValidation(BASE_HP_INVALID, form.BaseHp))
		}
		if form.BaseAtk < 0 {
			validations = append(validations, validation.NewValidation(BASE_ATK_INVALID, form.BaseAtk))
		}
		if form.BaseDfe < 0 {
			validations = append(validations, validation.NewValidation(BASE_DFE_INVALID, form.BaseDfe))
		}
		if form.BaseSpd < 0 {
			validations = append(validations, validation.NewValidation(BASE_SPD_INVALID, form.BaseSpd))
		}
		if form.BaseAts < 0 {
			validations = append(validations, validation.NewValidation(BASE_ATS_INVALID, form.BaseAts))
		}
		if form.BaseDfs < 0 {
			validations = append(validations, validation.NewValidation(BASE_DFS_INVALID, form.BaseDfs))
		}
		if form.EvHp < 0 {
			validations = append(validations, validation.NewValidation(EV_HP_INVALID, form.EvHp))
		}
		if form.EvAtk < 0 {
			validations = append(validations, validation.NewValidation(EV_ATK_INVALID, form.EvAtk))
		}
		if form.EvDfe < 0 {
			validations = append(validations, validation.NewValidation(EV_DFE_INVALID, form.EvDfe))
		}
		if form.EvSpd < 0 {
			validations = append(validations, validation.NewValidation(EV_SPD_INVALID, form.EvSpd))
		}
		if form.EvAts < 0 {
			validations = append(validations, validation.NewValidation(EV_ATS_INVALID, form.EvAts))
		}
		if form.EvDfs < 0 {
			validations = append(validations, validation.NewValidation(EV_DFS_INVALID, form.EvDfs))
		}
		if form.ExperienceType < 0 || form.ExperienceType > 5 {
			validations = append(validations, validation.NewValidation(EXPERIENCE_TYPE_INVALID, form.ExperienceType))
		}
		if form.BaseExperience < 0 {
			validations = append(validations, validation.NewValidation(BASE_EXPERIENCE_INVALID, form.BaseExperience))
		}
		if form.BaseLoyalty < 0 {
			validations = append(validations, validation.NewValidation(BASE_LOYALTY_INVALID, form.BaseLoyalty))
		}
		if form.CatchRate < 0 {
			validations = append(validations, validation.NewValidation(CATCH_RATE_INVALID, form.CatchRate))
		}
		if form.FemaleRate < -1 || form.FemaleRate > 100 {
			validations = append(validations, validation.NewValidation(FEMALE_RATE_INVALID, form.FemaleRate))
		}
		if len(form.BreedGroups) == 0 {
			validations = append(validations, validation.NewValidation(BREED_GROUPS_INVALID, form.BreedGroups))
		}
		if form.HatchSteps <= 0 {
			validations = append(validations, validation.NewValidation(HATCH_STEPS_INVALID, form.HatchSteps))
		}
		if strings.TrimSpace(form.BabyDbSymbol) == "" {
			validations = append(validations, validation.NewValidation(BABY_DB_SYMBOL_INVALID, form.BabyDbSymbol))
		}
		if form.BabyForm < 0 {
			validations = append(validations, validation.NewValidation(BABY_FORM_INVALID, form.BabyForm))
		}
		if form.FrontOffsetY < 0 {
			validations = append(validations, validation.NewValidation(FRONT_OFFSET_Y_INVALID, form.FrontOffsetY))
		}
		if len(form.Abilities) == 0 {
			validations = append(validations, validation.NewValidation(ABILITIES_INVALID, form.Abilities))
		}

		for _, item := range form.ItemHeld {
			if strings.TrimSpace(item.DbSymbol) == "" {
				validations = append(validations, validation.NewValidation(ITEM_HELD_DB_SYMBOL_INVALID, item.DbSymbol))
			}
			if item.Chance < 0 || item.Chance > 100 {
				validations = append(validations, validation.NewValidation(ITEM_HELD_CHANCE_INVALID, item.Chance))
			}
		}

		for _, evolution := range form.Evolutions {
			if evolution.Form < 0 {
				validations = append(validations, validation.NewValidation(FORM_INVALID, evolution.Form))
			}
			for _, condition := range evolution.Conditions {
				if strings.TrimSpace(condition.Type) == "" {
					validations = append(validations, validation.NewValidation(EVOLUTION_CONDITION_TYPE_INVALID, condition.Type))
				}
			}
		}

		if form.FormTextId == nil {
			validations = append(validations, validation.NewValidation(FORM_TEXT_ID_INVALID, form.FormTextId))
		}
	}

	return validations
}
