package main

import (
	"github.com/johanatandromeda/nftables-exporter/pkg/nftables"
	"log/slog"
	"os"
)

func main() {

	var programLevel = new(slog.LevelVar) // Info by default
	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))
	nftables.Test()
}
