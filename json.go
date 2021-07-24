package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"
	"time"
)

var (
	_ Logger        = (*jsonLogger)(nil)
	_ MessageLogger = (*jsonLogger)(nil)
)

type jsonLogger struct {
	mu sync.Mutex

	w io.Writer

	fields Fields

	depth int
}

func (l *jsonLogger) LogMessage(msg *Message) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.logMessage(msg)
}

func (l *jsonLogger) logMessage(msg *Message) error {
	var (
		err error

		rawMsg []byte
	)

	rawMsg, err = json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("could not marshal json: %w", err)
	}

	_, _ = fmt.Fprintf(l.w, "%s\n", string(rawMsg))

	return nil
}

func (l *jsonLogger) Print(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	summary := fmt.Sprint(v...)

	msg := &Message{
		Type:      MessageTypeDebug,
		Timestamp: time.Now(),
		Summary:   summary,
		Data:      l.fields,
	}

	msg.autofill(3)

	_ = l.logMessage(msg)
}

func (l *jsonLogger) Printf(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	summary := fmt.Sprintf(format, v...)

	msg := &Message{
		Type:      MessageTypeDebug,
		Timestamp: time.Now(),
		Summary:   summary,
		Data:      l.fields,
	}

	msg.autofill(3)

	_ = l.logMessage(msg)
}

func (l *jsonLogger) Error(v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	summary := fmt.Sprint(v...)

	msg := &Message{
		Type:      MessageTypeError,
		Timestamp: time.Now(),
		Summary:   summary,
		Data:      l.fields,
	}

	msg.autofill(3)

	_ = l.logMessage(msg)
}

func (l *jsonLogger) Errorf(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	summary := fmt.Sprintf(format, v...)

	msg := &Message{
		Type:      MessageTypeError,
		Timestamp: time.Now(),
		Summary:   summary,
		Data:      l.fields,
	}

	msg.autofill(3)

	_ = l.logMessage(msg)
}

func (l *jsonLogger) WithFields(fields Fields) Logger {
	l.mu.Lock()
	defer l.mu.Unlock()

	return &jsonLogger{
		w: l.w,

		fields: fields,
	}
}

func NewJSON(w io.Writer) Logger {
	if w == nil {
		return Discard
	}

	l := &jsonLogger{
		w: w,
	}

	return l
}
