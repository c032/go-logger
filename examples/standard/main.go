package main

import (
	stdLog "log"
	"os"

	"github.com/c032/go-logger"
)

func main() {
	stdLogger := stdLog.New(os.Stdout, "", stdLog.Ldate|stdLog.Ltime|stdLog.LUTC|stdLog.Lshortfile)

	log := logger.FromStandard(stdLogger)

	log.Printf("Hello, %s!", "world")

	log.WithFields(logger.Fields{
		"hello": "world",
	}).Print("Fields are ignored.")
}
