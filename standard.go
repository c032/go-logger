package logger

import (
	"fmt"
	"log"
)

var _ Logger = (*standardWrapper)(nil)

type standardWrapper struct {
	logger *log.Logger
}

func (sw *standardWrapper) Print(v ...interface{}) {
	sw.logger.Output(2, fmt.Sprint(v...))
}

func (sw *standardWrapper) Printf(format string, v ...interface{}) {
	sw.logger.Output(2, fmt.Sprintf(format, v...))
}

func (sw *standardWrapper) Error(v ...interface{}) {
	sw.logger.Output(2, fmt.Sprint(v...))
}

func (sw *standardWrapper) Errorf(format string, v ...interface{}) {
	sw.logger.Output(2, fmt.Sprintf(format, v...))
}

func (sw *standardWrapper) WithFields(fields Fields) Logger {
	return sw
}

func FromStandard(stdLogger *log.Logger) Logger {
	return &standardWrapper{
		logger: stdLogger,
	}
}
