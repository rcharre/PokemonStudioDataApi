package ps_test

import (
	"testing"

	"github.com/rcharre/psapi/pkg/studio"
	"github.com/rcharre/psapi/test"
)

func TestImport_Success(t *testing.T) {
	store := test.NewStoreMock()
	err := studio.Import(test.TestDataFolder, store)
	if err != nil {
		t.Error("Import valid studio folder should not returns error")
	}
}

func TestImport_FolderNotExists(t *testing.T) {
	store := test.NewStoreMock()
	err := studio.Import(test.TestDataFolder, store)
	if err != nil {
		t.Error("Import valid studio folder should not returns error")
	}
}

func TestImport_StudioFolderNotExists(t *testing.T) {
	store := test.NewStoreMock()
	err := studio.Import(test.TestDataFolder, store)
	if err != nil {
		t.Error("Import valid studio folder should not returns error")
	}
}

func TestImport_TextFolderNotExists(t *testing.T) {
	store := test.NewStoreMock()
	err := studio.Import(test.TestDataFolder, store)
	if err != nil {
		t.Error("Import valid studio folder should not returns error")
	}
}
