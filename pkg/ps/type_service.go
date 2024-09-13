package ps

type TypeService interface{}

type TypeServiceImpl struct{}

func NewTypeService() TypeService {
	return &TypeServiceImpl{}
}
