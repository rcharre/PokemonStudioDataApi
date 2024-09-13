package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log/slog"
	"net/http"
	"psapi/pkg/api"
	"psapi/pkg/ps"
	"psapi/pkg/utils/config"
	"time"
)

func main() {
	rootCmd := &cobra.Command{
		Short: "pokemon studio data api",
		Long:  `pokemon studio data api is an easy to use api for pokemon studio data. All you need is a pokemon studio data folder and a mongodb server`,
		Run:   run,
	}

	rootCmd.Flags().StringP("log-level", "l", config.DefaultAppLogLevel, "log level")
	rootCmd.Flags().StringP("data", "d", config.DefaultImportDataFolderPath, "data folder")
	rootCmd.Flags().IntP("port", "p", config.DefaultApiPort, "port to serve on")
	rootCmd.Flags().StringP("cors", "C", config.DefaultApiCors, "cors headers")
	rootCmd.InitDefaultHelpCmd()

	_ = viper.BindPFlag(config.KeyApiPort, rootCmd.Flags().Lookup("port"))
	_ = viper.BindPFlag(config.KeyApiCors, rootCmd.Flags().Lookup("cors"))
	_ = viper.BindPFlag(config.KeyImportDataFolderPath, rootCmd.Flags().Lookup("data"))
	_ = viper.BindPFlag(config.KeyAppLogLevel, rootCmd.Flags().Lookup("log-level"))

	cobra.OnInitialize(config.InitConfig)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func run(cmd *cobra.Command, args []string) {
	appConfig := config.GetConfig()
	ctx := ps.NewContext()

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Throttle(100))
	r.Use(middleware.Timeout(5 * time.Second))
	r.Use(middleware.SetHeader("Access-Control-Allow-Origin", appConfig.Api.Cors))

	dataFolderPath := appConfig.Import.DataFolderPath
	if err := ps.ImportPokemonStudioFolder(dataFolderPath, ctx); err != nil {
		panic(err)
	}

	pokemonController := api.NewPokemonsAPIController(ctx.PokemonService())
	r.Mount("/", api.NewRouter(pokemonController))

	addr := fmt.Sprintf(":%d", appConfig.Api.Port)
	slog.Info("Server listening", "addr", addr)
	_ = http.ListenAndServe(addr, r)
}
