package logger

var Discard Logger = &discardLogger{}

var _ MessageLogger = (*discardLogger)(nil)

type discardLogger struct{}

func (l *discardLogger) LogMessage(msg *Message) error {
	return nil
}

func (l *discardLogger) Print(v ...interface{}) {
}

func (l *discardLogger) Printf(format string, v ...interface{}) {
}

func (l *discardLogger) Error(v ...interface{}) {
}

func (l *discardLogger) Errorf(format string, v ...interface{}) {
}

func (l *discardLogger) WithFields(fields Fields) Logger {
	return l
}
