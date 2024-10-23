package pscli

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"

	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rcharre/psapi/pkg/ps"
	"github.com/rcharre/psapi/pkg/psapi"
)

const (
	AppConfigDefaultName = "psapi"

	KeyAppLogLevel     = "log-level"
	DefaultAppLogLevel = "INFO"

	KeyApiCors     = "cors"
	DefaultApiCors = "*"

	KeyApiPort     = "port"
	DefaultApiPort = 8000

	KeyImportDataFolderPath     = "data"
	DefaultImportDataFolderPath = "data"
)

var serveFlagSet = flag.NewFlagSet("", flag.ExitOnError)
var logLevel = serveFlagSet.String(KeyAppLogLevel, DefaultAppLogLevel, "The log level")
var dataFolder = serveFlagSet.String(KeyImportDataFolderPath, DefaultImportDataFolderPath, "Data folder")
var port = serveFlagSet.Int(KeyApiPort, DefaultApiPort, "port to serve server on")
var cors = serveFlagSet.String(KeyApiCors, DefaultApiCors, "cors header")

var ServeCmd = NewCommand(serveFlagSet, run)

func run() error {
	slog.Info("Flag", "data", serveFlagSet.Lookup(KeyImportDataFolderPath))
	ParseLogLevel(*logLevel)

	studio := ps.NewInMemoryStudio()
	if err := studio.Import(*dataFolder); err != nil {
		return err
	}
	psapiRouter := psapi.NewPsApiHandler(studio)

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Throttle(100))
	r.Use(middleware.Timeout(5 * time.Second))
	r.Use(middleware.SetHeader("Access-Control-Allow-Origin", *cors))
	r.Mount("/", psapiRouter)

	addr := fmt.Sprintf(":%d", *port)
	slog.Info("Server listening", "addr", addr)
	_ = http.ListenAndServe(addr, r)
	return nil
}
