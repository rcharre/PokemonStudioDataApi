package ps

type TypeImporter interface{}

type TypeImporterImpl struct{}

func NewTypeImporter() TypeImporter {
	return &TypeImporterImpl{}
}
