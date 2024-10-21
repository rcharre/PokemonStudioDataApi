package ps_test

import (
	"iter"
	"psapi/pkg/ps"
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
	_, err := importer.Import(TestStudioFolder, "invalid_folder")

	if err == nil {
		t.Error("Import with invalid translation folder should give an error")
	}
}

func TestTypeImporter_Import_NoStudioFolder(t *testing.T) {
	importer := ps.NewTypeImporter()
	_, err := importer.Import("invalid_folder", TestTranslationFolder)

	if err == nil {
		t.Error("Import with invalid studio folder should give an error")
	}
}

func TestTypeImporter_Import_ApplyTranslation(t *testing.T) {
	importer := ps.NewTypeImporter()
	iterator, err := importer.Import(TestStudioFolder, TestTranslationFolder)

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
