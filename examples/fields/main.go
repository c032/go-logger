package main

import (
	"time"

	"github.com/c032/go-logger"
)

func main() {
	log := logger.Default

	log.Print("No fields.")

	log.WithFields(logger.Fields{
		"example_name":   "fields",
		"unix_timestamp": time.Now().Unix(),
	}).Printf("Hello, %s.", "world")

	log.Print("No fields again.")
}
