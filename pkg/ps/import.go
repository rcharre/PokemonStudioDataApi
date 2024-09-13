package ps

import (
	"log/slog"
	"path"
)

const (
	StudioFolder = "Studio"

	PokemonFolder = "pokemon"
	TypeFolder    = "type"
	AbilityFolder = "ability"
	MoveFolder    = "move"
	TrainerFolder = "trainer"

	LanguageFolder = "language"
)

func ImportPokemonStudioFolder(folder string, context *Context) error {
	slog.Info("Importing pokemon studio folder", "path", folder)
	pokemonFolderPath := path.Join(folder, StudioFolder, PokemonFolder)
	if err := context.PokemonImporter().ImportFolder(pokemonFolderPath); err != nil {
		return err
	}

	return nil
}
