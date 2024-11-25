package psapi

import (
	"context"

	"github.com/rcharre/psapi/pkg/psapi/psapigen"
	"github.com/rcharre/psapi/pkg/studio"
)

type TypeService struct {
	store      *studio.Store
	typeMapper *TypeMapper
}

func NewTypeService(store *studio.Store, typeMapper *TypeMapper) psapigen.TypesAPIServicer {
	return &TypeService{
		store,
		typeMapper,
	}
}

func (s TypeService) GetTypes(requestCtx context.Context, lang string) (psapigen.ImplResponse, error) {
	types := s.store.TypeStore.FindAll()
	res := make([]psapigen.TypePartial, len(types))

	for i, t := range types {
		res[i] = s.typeMapper.ToTypePartial(t, lang)
	}
	return psapigen.ImplResponse{Code: 200, Body: res}, nil
}

func (s TypeService) GetTypeDetails(requestCtx context.Context, symbol string, lang string) (psapigen.ImplResponse, error) {
	t := s.store.TypeStore.FindBySymbol(symbol)
	if t == nil {
		return psapigen.ImplResponse{Code: 200, Body: nil}, nil
	}
	return psapigen.ImplResponse{Code: 200, Body: s.typeMapper.ToTypeDetail(*t, lang)}, nil
}
