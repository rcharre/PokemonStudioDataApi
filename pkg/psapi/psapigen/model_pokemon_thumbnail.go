// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Pokemon Studio API
 *
 * API for the Pokemon Studio
 *
 * API version: 0.0.1
 */

package psapigen

// PokemonThumbnail - An object containing base information to display a pokemon
type PokemonThumbnail struct {

	// The symbol of the pokemon
	Symbol string `json:"symbol,omitempty"`

	// The translated name of the first pokemon's form
	Name string `json:"name,omitempty"`

	// The number of the pokemon
	Number int32 `json:"number,omitempty"`

	// The image symbol of the pokemon
	Image string `json:"image,omitempty"`
}

// AssertPokemonThumbnailRequired checks if the required fields are not zero-ed
func AssertPokemonThumbnailRequired(obj PokemonThumbnail) error {
	return nil
}

// AssertPokemonThumbnailConstraints checks if the values respects the defined constraints
func AssertPokemonThumbnailConstraints(obj PokemonThumbnail) error {
	return nil
}
