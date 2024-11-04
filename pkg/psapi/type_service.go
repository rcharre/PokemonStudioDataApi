package psapi

import (
	"context"

	"github.com/rcharre/psapi/pkg/pkmn"
	"github.com/rcharre/psapi/pkg/psapi/psapigen"
)

type TypeServiceImpl struct {
	typeStore  pkmn.TypeStore
	typeMapper TypeMapper
}

func NewTypeService(typeStore pkmn.TypeStore, typeMapper TypeMapper) psapigen.TypesAPIServicer {
	return &TypeServiceImpl{
		typeStore,
		typeMapper,
	}
}

func (s TypeServiceImpl) GetTypes(requestCtx context.Context, lang string) (psapigen.ImplResponse, error) {
	types := s.typeStore.FindAll()
	res := make([]*psapigen.TypePartial, len(types))

	for i, t := range types {
		res[i] = s.typeMapper.ToTypePartial(t, lang)
	}
	return psapigen.ImplResponse{Code: 200, Body: res}, nil
}

func (s TypeServiceImpl) GetTypeDetails(requestCtx context.Context, symbol string, lang string) (psapigen.ImplResponse, error) {
	t := s.typeStore.FindBySymbol(symbol)
	if t == nil {
		return psapigen.ImplResponse{Code: 200, Body: nil}, nil
	}
	return psapigen.ImplResponse{Code: 200, Body: s.typeMapper.ToTypeDetail(t, lang)}, nil
}
