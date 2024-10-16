package psapi

import (
	"context"
	"psapi/pkg/ps"
	"psapi/pkg/psapi/psapigen"
	"psapi/pkg/utils/pagination"
)

type PokemonServiceImpl struct {
	pokemonStore  ps.PokemonStore
	pokemonMapper PokemonMapper
}

func NewPokemonService(pokemonStore ps.PokemonStore, pokemonMapper PokemonMapper) *PokemonServiceImpl {
	return &PokemonServiceImpl{
		pokemonStore:  pokemonStore,
		pokemonMapper: pokemonMapper,
	}
}

func (s PokemonServiceImpl) GetPokemonDetails(requestCtx context.Context, symbol string, lang string) (psapigen.ImplResponse, error) {
	pkmn := s.pokemonStore.FindBySymbol(symbol)

	var res *psapigen.PokemonDetails
	if pkmn != nil {
		res = s.pokemonMapper.PokemonToDetail(pkmn, lang)
	}
	return psapigen.ImplResponse{Code: 200, Body: res}, nil
}

func (s PokemonServiceImpl) GetPokemon(requestCtx context.Context, page int32, pageSize int32, lang string) (psapigen.ImplResponse, error) {
	p := int(page)
	size := int(pageSize)

	pr := pagination.NewPageRequest(p, size)

	pkmnIter := s.pokemonStore.FindAll()
	pkmnPage := pagination.ApplyPageRequestToIter(pr, pkmnIter)
	thumbnails := make([]*psapigen.PokemonThumbnail, len(pkmnPage.Content))

	for i, pkmn := range pkmnPage.Content {
		thumbnails[i] = s.pokemonMapper.PokemonToThumbnail(pkmn, lang)
	}

	return psapigen.ImplResponse{Code: 200, Body: pagination.NewPageFromPageRequest(pr, thumbnails, pkmnPage.Total)}, nil
}

func (s PokemonServiceImpl) GetPokemonForm(requestCtx context.Context, symbol string, form int32, lang string) (psapigen.ImplResponse, error) {
	f := int(form)
	pkmn := s.pokemonStore.FindBySymbol(symbol)

	if pkmn == nil {
		return psapigen.ImplResponse{Code: 404, Body: nil}, nil
	}

	if f > len(pkmn.Forms)-1 {
		return psapigen.ImplResponse{Code: 404, Body: nil}, nil
	}

	pkmnForm := pkmn.Forms[f]
	return psapigen.ImplResponse{Code: 200, Body: s.pokemonMapper.FormToPokemonForm(pkmnForm, lang)}, nil
}
