package main

import (
	"flag"
	"fmt"
	"github.com/johanatandromeda/nftables-exporter/pkg/nftables"
	"log/slog"
	"os"
)

var version = ""

func main() {

	var programLevel = new(slog.LevelVar) // Info by default
	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))

	slog.Info(fmt.Sprintf("Starting nftables-exporter V %s\n", version))

	debug := flag.Bool("d", false, "Debug")

	flag.Parse()

	if *debug {
		programLevel.Set(slog.LevelDebug)
	}

	_, err := nftables.GetCounter()
	if err != nil {
		panic(err)
	}
}
