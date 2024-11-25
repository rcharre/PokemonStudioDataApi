package studio_test

import (
	"testing"

	"github.com/rcharre/psapi/pkg/studio"
	"github.com/rcharre/psapi/test"
)

func TestImportTranslations_FileNotFound(t *testing.T) {
	_, err := studio.ImportTranslations(test.InvalidPath)
	if err == nil {
		t.Error("Import translation at invalid path should return error")
	}
}

func TestImportTranslations_InvalidFormat(t *testing.T) {
	_, err := studio.ImportTranslations(test.TranslationInvalid)
	if err == nil {
		t.Error("Import translation with invalid format should return error")
	}

}

func TestImportTranslations_Ok(t *testing.T) {
	_, err := studio.ImportTranslations(test.TranslationValid)
	if err == nil {
		t.Error("Import translation with valid file should not return error")
	}
}
