package ps

type TypeMapper interface {
}

type TypeMapperImpl struct {
}

func NewTypeMapper() TypeMapper {
	return &TypeMapperImpl{}
}
