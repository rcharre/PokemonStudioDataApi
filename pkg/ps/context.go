package ps

import "psapi/pkg/api"

type Context struct {
	pokemonStore     PokemonStore
	pokemonMapper    PokemonMapper
	pokemonService   api.PokemonsAPIServicer
	pokemonImporter  PokemonImporter
	pokemonValidator PokemonValidator
}

func NewContext() *Context {
	pokemonStore := NewPokemonStore()
	pokemonMapper := NewPokemonMapper()
	pokemonService := NewPokemonService(pokemonStore, pokemonMapper)
	pokemonValidator := NewPokemonValidator()
	pokemonImporter := NewPokemonImporter(pokemonStore, pokemonValidator)

	return &Context{
		pokemonStore,
		pokemonMapper,
		pokemonService,
		pokemonImporter,
		pokemonValidator,
	}
}

func (c *Context) PokemonStore() PokemonStore {
	return c.pokemonStore
}

func (c *Context) PokemonMapper() PokemonMapper {
	return c.pokemonMapper
}

func (c *Context) PokemonService() api.PokemonsAPIServicer {
	return c.pokemonService
}

func (c *Context) PokemonImporter() PokemonImporter {
	return c.pokemonImporter
}

func (c *Context) PokemonValidator() PokemonValidator {
	return c.pokemonValidator
}
