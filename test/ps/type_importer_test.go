package ps_test

import (
	"github.com/rcharre/psapi/pkg/ps"
	"github.com/rcharre/psapi/test"
	"iter"
	"slices"
	"testing"
)

var _ ps.TypeImporter = &TypeImporterMock{}

type TypeImporterMock struct {
	ImportFunc func(string, string) (iter.Seq[*ps.PokemonType], error)
}

func (i TypeImporterMock) Import(studioFolder string, translationFolder string) (iter.Seq[*ps.PokemonType], error) {
	return i.ImportFunc(studioFolder, translationFolder)
}

func TestTypeImporter_Import_NoTranslationFolder(t *testing.T) {
	importer := ps.NewTypeImporter()
	_, err := importer.Import(test.TestStudioFolder, "invalid_folder")

	if err == nil {
		t.Error("Import with invalid translation folder should give an error")
	}
}

func TestTypeImporter_Import_NoStudioFolder(t *testing.T) {
	importer := ps.NewTypeImporter()
	_, err := importer.Import("invalid_folder", test.TestTranslationFolder)

	if err == nil {
		t.Error("Import with invalid studio folder should give an error")
	}
}

func TestTypeImporter_Import_ApplyTranslation(t *testing.T) {
	importer := ps.NewTypeImporter()
	iterator, err := importer.Import(test.TestStudioFolder, test.TestTranslationFolder)

	if err != nil {
		t.Error("Import with valid folders should not return error")
	}

	typeList := slices.Collect(iterator)
	expectedLen := 18
	resultLen := len(typeList)
	if resultLen != expectedLen {
		t.Error("Import expected size", expectedLen, ",has", resultLen)
	}

	for _, pokeType := range typeList {
		if pokeType.Name == nil {
			t.Error("Type name translation map should not be null")
		}
	}
}
