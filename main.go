package main

import (
	"log/slog"

	"github.com/rcharre/psapi/internal/pscli"
)

func main() {
	serveCmd := pscli.ServeCmd
	if err := serveCmd.Execute(); err != nil {
		slog.Error(err.Error())
		serveCmd.Usage()
	}
}
