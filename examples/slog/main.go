package main

import (
	"log/slog"
	"os"

	"github.com/c032/go-logger"
)

func main() {
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	log := logger.FromSlog(l)

	log.Printf("Hello, %s!", "world")

	log.WithFields(logger.Fields{
		"hello": "world",
	}).Print("Example field.")
}
