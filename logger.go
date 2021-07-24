package logger

type Logger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	WithFields(fields Fields) Logger
}
