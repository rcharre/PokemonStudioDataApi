package ps

import (
	"context"
	"psapi/pkg/api"
	"psapi/pkg/utils/pagination"
)

type PokemonServiceImpl struct {
	pokemonStore  PokemonStore
	pokemonMapper PokemonMapper
}

func NewPokemonService(pokemonStore PokemonStore, pokemonMapper PokemonMapper) api.PokemonsAPIServicer {
	return &PokemonServiceImpl{
		pokemonStore:  pokemonStore,
		pokemonMapper: pokemonMapper,
	}
}

func (s PokemonServiceImpl) GetPokemon(requestCtx context.Context, symbol string, lang string) (api.ImplResponse, error) {
	pkmn := s.pokemonStore.FindBySymbol(symbol)

	var res *api.PokemonDetail
	if pkmn != nil {
		res = s.pokemonMapper.PokemonToDetail(pkmn, lang)
	}
	return api.ImplResponse{Code: 200, Body: res}, nil
}

func (s PokemonServiceImpl) GetPokemons(requestCtx context.Context, page int32, pageSize int32, lang string) (api.ImplResponse, error) {
	p := int(page)
	size := int(pageSize)

	pr := pagination.NewPageRequest(p, size)

	pkmnIter := s.pokemonStore.FindAll()
	pkmnPage := pagination.ApplyPageRequestToIter(pr, pkmnIter)
	thumbnails := make([]*api.PokemonThumbnail, len(pkmnPage.Content))

	for i, pkmn := range pkmnPage.Content {
		thumbnails[i] = s.pokemonMapper.PokemonToThumbnail(pkmn, lang)
	}

	return api.ImplResponse{Code: 200, Body: pagination.NewPageFromPageRequest(pr, thumbnails, pkmnPage.Total)}, nil
}

func (s PokemonServiceImpl) GetPokemonForm(requestCtx context.Context, symbol string, form int32, lang string) (api.ImplResponse, error) {
	f := int(form)
	pkmn := s.pokemonStore.FindBySymbol(symbol)

	if pkmn == nil {
		return api.ImplResponse{Code: 404, Body: nil}, nil
	}

	if f > len(pkmn.Forms)-1 {
		return api.ImplResponse{Code: 404, Body: nil}, nil
	}

	pkmnForm := pkmn.Forms[f]
	return api.ImplResponse{Code: 200, Body: s.pokemonMapper.FormToPokemonForm(pkmnForm, lang)}, nil
}
