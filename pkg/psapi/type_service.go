package psapi

type TypeService interface{}

type TypeServiceImpl struct{}

func NewTypeService() TypeService {
	return &TypeServiceImpl{}
}
