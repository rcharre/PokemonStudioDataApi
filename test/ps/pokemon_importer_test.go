package ps_test

import (
	"iter"
	"psapi/pkg/ps"
	"slices"
	"testing"
)

const (
	TestStudioFolder      = "../test_resources/Studio/"
	TestTranslationFolder = "../test_resources/Dialogs"
)

var _ ps.PokemonImporter = &PokemonImporterMock{}

type PokemonImporterMock struct {
	ImportFunc func(studioFolder string, translationFolder string) (iter.Seq[*ps.Pokemon], error)
}

func (i *PokemonImporterMock) Import(studioFolder string, translationFolder string) (iter.Seq[*ps.Pokemon], error) {
	return i.ImportFunc(studioFolder, translationFolder)
}

func TestImport_NoTranslationFolder(t *testing.T) {
	importer := ps.NewPokemonImporter()
	_, err := importer.Import(TestStudioFolder, "invalid_folder")

	if err == nil {
		t.Error("Import with invalid translation folder should give an error")
	}
}

func TestImport_NoStudioFolder(t *testing.T) {
	importer := ps.NewPokemonImporter()
	_, err := importer.Import("invalid_folder", TestTranslationFolder)

	if err == nil {
		t.Error("Import with invalid studio folder should give an error")
	}
}

func TestImport_ApplyTranslation(t *testing.T) {
	importer := ps.NewPokemonImporter()
	iterator, err := importer.Import(TestStudioFolder, TestTranslationFolder)

	if err != nil {
		t.Error("Import with valid folders should not return error")
	}

	pokeList := slices.Collect(iterator)
	expectedLen := 5
	resultLen := len(pokeList)
	if resultLen != expectedLen {
		t.Error("Import expected size", expectedLen, ",has", resultLen)
	}

	for _, pokemon := range pokeList {
		for _, form := range pokemon.Forms {
			if form.Name == nil {
				t.Error("Form name translation map should not be null")
			}

			if form.Description == nil {
				t.Error("Form description translation map should not be null")
			}
		}
	}
}
