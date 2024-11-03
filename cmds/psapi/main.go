package main

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/rcharre/psapi/internal/psapicli"
)

func main() {
	if err := psapicli.ServeCmd.Execute(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error(err.Error())
		psapicli.ServeCmd.Usage()
	}
}
