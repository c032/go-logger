package logger

import (
	"os"
	"sync"
)

var (
	defaultLogger      Logger = nil
	defaultLoggerMutex sync.Mutex
)

func Default() Logger {
	defaultLoggerMutex.Lock()
	defer defaultLoggerMutex.Unlock()

	if defaultLogger == nil {
		defaultLogger = NewJSON(os.Stdout)
	}

	return defaultLogger
}
