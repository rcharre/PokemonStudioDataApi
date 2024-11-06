package test

import "github.com/rcharre/psapi/pkg/studio"

const (
	TestResourcesFolder = "../test_resources/"
	DataFolder          = TestResourcesFolder + "valid-data/"
	StudioFolder        = DataFolder + "/Studio/"
	TranslationFolder   = DataFolder + "Text/Dialogs"

	TranslationValid   = TranslationFolder + "1000063.csv"
	TranslationInvalid = TestResourcesFolder + "100003-invalid.csv"

	PokemonFolder      = StudioFolder + studio.PokemonFolder
	PokemonValid       = PokemonFolder + "abra.json"
	PokemonBadJson     = TestResourcesFolder + "abra-cut.json"
	PokemonInvalidType = TestResourcesFolder + "abra-invalid-primitive.json"

	InvalidPath = "invalid"
)
