package ps

type PokemonType struct {
	DbSymbol string
	Color    string
	DamageTo []TypeDamage
}

type TypeDamage struct {
	DefensiveType string
	Factor        float64
}
