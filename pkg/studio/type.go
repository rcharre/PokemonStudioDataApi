package studio

import "github.com/rcharre/psapi/pkg/utils/i18n"

type PokemonType struct {
	DbSymbol string
	Color    string
	TextId   int
	Name     i18n.Translation
	DamageTo []TypeDamage
}

type TypeDamage struct {
	DefensiveType string
	Factor        float32
}
