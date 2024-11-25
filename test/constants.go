package test

import "github.com/rcharre/psapi/pkg/studio"

const (
	TestResourcesFolder = "../test_resources/"
	DataFolder          = TestResourcesFolder + "valid-data/"
	StudioFolder        = DataFolder + "Studio/"
	TranslationFolder   = DataFolder + "Text/Dialogs/"

	TranslationValid   = TranslationFolder + "1000063.csv"
	TranslationInvalid = TestResourcesFolder + "100003-invalid.csv"

	PokemonFolder  = StudioFolder + studio.PokemonFolder
	PokemonValid   = PokemonFolder + "abra.json"
	PokemonInvalid = TestResourcesFolder + "abra-invalid.json"

	TypeFolder  = StudioFolder + studio.TypeFolder
	TypeValid   = TypeFolder + "bug.json"
	TypeInvalid = TestResourcesFolder + "bug-invalid.json"

	InvalidPath = "invalid"
)
