package psapicli

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"

	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rcharre/psapi/pkg/psapi"
	"github.com/rcharre/psapi/pkg/storage/inmem"
	"github.com/rcharre/psapi/pkg/studio"
	"github.com/rcharre/psapi/pkg/utils/cli"
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

var ServeCmd = cli.NewCommand(serveFlagSet, run)

func run() error {
	slog.Debug("Flag", "data", serveFlagSet.Lookup(KeyImportDataFolderPath))
	ParseLogLevel(*logLevel)

	inMemStore := inmem.New()

	if err := studio.Import(*dataFolder, inMemStore); err != nil {
		return err
	}
	psapiRouter := psapi.NewPsApiHandler(inMemStore)

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Throttle(100))
	r.Use(middleware.Timeout(5 * time.Second))
	r.Use(middleware.SetHeader("Access-Control-Allow-Origin", *cors))
	r.Mount("/", psapiRouter)

	addr := fmt.Sprintf(":%d", *port)
	slog.Info("Server listening", "addr", addr)

	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	go listenInterrupt(server)
	return server.ListenAndServe()

}

func listenInterrupt(server *http.Server) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan
	slog.Info("Closing server...")
	if err := server.Close(); err != nil {
		slog.Error(err.Error())
	}
}
