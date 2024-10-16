package pscli

import (
	"fmt"
	"log/slog"
	"net/http"
	"psapi/pkg/ps"
	"psapi/pkg/psapi"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewServeCmd() *cobra.Command {
	serveCmd := &cobra.Command{
		Short: "pokemon studio data api",
		Long:  `pokemon studio data api is an easy to use api for pokemon studio data. All you need is a pokemon studio data folder and a mongodb server`,
		Run:   run,
	}

	serveCmd.Flags().StringP("log-level", "l", DefaultAppLogLevel, "log level")
	serveCmd.Flags().StringP("data", "d", DefaultImportDataFolderPath, "data folder")
	serveCmd.Flags().IntP("port", "p", DefaultApiPort, "port to serve on")
	serveCmd.Flags().StringP("cors", "C", DefaultApiCors, "cors headers")
	serveCmd.InitDefaultHelpCmd()

	_ = viper.BindPFlag(KeyApiPort, serveCmd.Flags().Lookup("port"))
	_ = viper.BindPFlag(KeyApiCors, serveCmd.Flags().Lookup("cors"))
	_ = viper.BindPFlag(KeyImportDataFolderPath, serveCmd.Flags().Lookup("data"))
	_ = viper.BindPFlag(KeyAppLogLevel, serveCmd.Flags().Lookup("log-level"))

	cobra.OnInitialize(InitConfig)
	return serveCmd
}

func run(cmd *cobra.Command, args []string) {
	appConfig := GetConfig()
	dataFolderPath := appConfig.Import.DataFolderPath
	studio, err := ps.NewInMemoryStudio(dataFolderPath)
	if err != nil {
		panic(err)
	}

	psapiRouter := psapi.NewPsApiHandler(studio)

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
