package ps

var _ TypeStore = &InMemoryTypeStore{}

type TypeStore interface {
	Add(pokemonType *PokemonType)
	FindBySymbol(symbol string) *PokemonType
	FindAll() []*PokemonType
}

type InMemoryTypeStore struct {
	pokemonTypesBySymbol map[string]*PokemonType
	types                []*PokemonType
}

// NewInMemoryTypeStore Create a new in memory type store
// types The list of types to store
func NewInMemoryTypeStore() *InMemoryTypeStore {
	return &InMemoryTypeStore{
		pokemonTypesBySymbol: make(map[string]*PokemonType),
		types:                make([]*PokemonType, 0),
	}
}

// Add add a pokemon type to the store
// pokemonType the type to add
func (s *InMemoryTypeStore) Add(pokemonType *PokemonType) {
	s.types = append(s.types, pokemonType)
	s.pokemonTypesBySymbol[pokemonType.DbSymbol] = pokemonType
}

// FindBySymbol Find a type by its symbol
// symbol The symbol to find
func (s InMemoryTypeStore) FindBySymbol(symbol string) *PokemonType {
	return s.pokemonTypesBySymbol[symbol]
}

// FindAll Find all types in the store
func (s InMemoryTypeStore) FindAll() []*PokemonType {
	return s.types
}
