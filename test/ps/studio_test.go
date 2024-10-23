package ps_test

import (
	"errors"
	"github.com/rcharre/psapi/pkg/ps"
	"github.com/rcharre/psapi/test"
	"iter"
	"slices"
	"testing"
)

func TestIntegration_NewInMemoryStudio(t *testing.T) {
	studio := ps.NewInMemoryStudio()
	err := studio.Import(test.TestDataFolder)
	if err != nil {
		t.Error("Import valid studio folder should not returns error")
	}
}

func TestImport_FailTypeImport(t *testing.T) {
	pokemonStore := &PokemonStoreMock{}
	pokemonImporter := &PokemonImporterMock{}

	typeStore := &TypeStoreMock{}
	typeImporter := &TypeImporterMock{}

	studio := ps.NewStudio(typeStore, typeImporter, pokemonStore, pokemonImporter)

	typeImporter.ImportFunc = func(s1, s2 string) (iter.Seq[*ps.PokemonType], error) {
		return nil, errors.New("import failed")
	}

	err := studio.Import(test.TestDataFolder)
	if err == nil {
		t.Error("Studio import should return error when type import failed")
	}
}

func TestImport_FailPokemonImport(t *testing.T) {
	pokemonStore := &PokemonStoreMock{}
	pokemonImporter := &PokemonImporterMock{}

	typeStore := &TypeStoreMock{}
	typeImporter := &TypeImporterMock{}

	studio := ps.NewStudio(typeStore, typeImporter, pokemonStore, pokemonImporter)

	typeImporter.ImportFunc = func(s1, s2 string) (iter.Seq[*ps.PokemonType], error) {
		return slices.Values([]*ps.PokemonType{}), nil
	}

	typeStore.AddFunc = func(pokemonType *ps.PokemonType) {}

	pokemonImporter.ImportFunc = func(studioFolder, translationFolder string) (iter.Seq[*ps.Pokemon], error) {
		return nil, errors.New("Import failed")
	}

	err := studio.Import(test.TestDataFolder)
	if err == nil {
		t.Error("Studio import should return error when pokemon import failed")
	}
}

func TestImport_Success(t *testing.T) {
	{
		pokemonStore := &PokemonStoreMock{}
		pokemonImporter := &PokemonImporterMock{}

		typeStore := &TypeStoreMock{}
		typeImporter := &TypeImporterMock{}

		studio := ps.NewStudio(typeStore, typeImporter, pokemonStore, pokemonImporter)

		typeImporter.ImportFunc = func(s1, s2 string) (iter.Seq[*ps.PokemonType], error) {
			return slices.Values([]*ps.PokemonType{}), nil
		}

		typeStore.AddFunc = func(pokemonType *ps.PokemonType) {}

		pokemonImporter.ImportFunc = func(studioFolder, translationFolder string) (iter.Seq[*ps.Pokemon], error) {
			return slices.Values([]*ps.Pokemon{}), nil
		}

		err := studio.Import(test.TestDataFolder)
		if err != nil {
			t.Error("Studio import should not return error when imports succeeded")
		}
	}
}
