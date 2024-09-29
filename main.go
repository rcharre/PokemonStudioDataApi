package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"psapi/internal/pscli"
	"psapi/pkg/ps"
	"psapi/pkg/psapi"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	rootCmd := &cobra.Command{
		Short: "pokemon studio data api",
		Long:  `pokemon studio data api is an easy to use api for pokemon studio data. All you need is a pokemon studio data folder and a mongodb server`,
		Run:   run,
	}

	rootCmd.Flags().StringP("log-level", "l", pscli.DefaultAppLogLevel, "log level")
	rootCmd.Flags().StringP("data", "d", pscli.DefaultImportDataFolderPath, "data folder")
	rootCmd.Flags().IntP("port", "p", pscli.DefaultApiPort, "port to serve on")
	rootCmd.Flags().StringP("cors", "C", pscli.DefaultApiCors, "cors headers")
	rootCmd.InitDefaultHelpCmd()

	_ = viper.BindPFlag(pscli.KeyApiPort, rootCmd.Flags().Lookup("port"))
	_ = viper.BindPFlag(pscli.KeyApiCors, rootCmd.Flags().Lookup("cors"))
	_ = viper.BindPFlag(pscli.KeyImportDataFolderPath, rootCmd.Flags().Lookup("data"))
	_ = viper.BindPFlag(pscli.KeyAppLogLevel, rootCmd.Flags().Lookup("log-level"))

	cobra.OnInitialize(pscli.InitConfig)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func run(cmd *cobra.Command, args []string) {
	appConfig := pscli.GetConfig()
	app := ps.NewDefaultApp()

	dataFolderPath := appConfig.Import.DataFolderPath
	if err := app.ImportData(dataFolderPath); err != nil {
		panic(err)
	}

	psapiRouter := psapi.NewPsApiHandler(app)

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Throttle(100))
	r.Use(middleware.Timeout(5 * time.Second))
	r.Use(middleware.SetHeader("Access-Control-Allow-Origin", appConfig.Api.Cors))
	r.Mount("/", psapiRouter)

	addr := fmt.Sprintf(":%d", appConfig.Api.Port)
	slog.Info("Server listening", "addr", addr)
	_ = http.ListenAndServe(addr, r)
}
