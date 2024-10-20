package main

import (
	"log/slog"
	"psapi/internal/pscli"
)

func main() {
	serveCmd := pscli.ServeCmd
	if err := serveCmd.Execute(); err != nil {
		slog.Error(err.Error())
		serveCmd.Usage()
	}
}
