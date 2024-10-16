package main

import (
	"psapi/internal/pscli"
)

func main() {
	serveCmd := pscli.NewServeCmd()
	if err := serveCmd.Execute(); err != nil {
		panic(err)
	}
}
