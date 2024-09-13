package ps

type TypeValidator interface {
}

type TypeValidatorImpl struct {
}

func NewTypeValidator() TypeValidator {
	return &TypeValidatorImpl{}
}
