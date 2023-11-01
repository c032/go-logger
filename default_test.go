package logger_test

import (
	"testing"

	"github.com/c032/go-logger"
)

func TestDefault(t *testing.T) {
	a := logger.Default()
	b := logger.Default()
	if got, want := a, b; got != want {
		t.Fatal("logger.Default() should always return the same value")
	}
}
