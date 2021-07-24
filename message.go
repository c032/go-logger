package logger

import (
	"runtime"
	"time"
)

type MessageType string

const (
	MessageTypeDebug   MessageType = "debug"
	MessageTypeWarning MessageType = "warning"
	MessageTypeError   MessageType = "error"
)

type Message struct {
	Type      MessageType `json:"type"`
	Timestamp time.Time   `json:"timestamp"`

	File     string `json:"file"`
	Line     int    `json:"line"`
	Function string `json:"function"`

	Summary string `json:"summary"`
	Data    Fields `json:"data"`
}

func (msg *Message) autofill(calldepth int) {
	var (
		ok bool

		file     string
		line     int
		function string
	)

	pc := make([]uintptr, 15)
	n := runtime.Callers(calldepth, pc)
	frames := runtime.CallersFrames(pc[:n])
	if n > 0 {
		ok = true

		var frame runtime.Frame
		frame, _ = frames.Next()

		file = frame.File
		line = frame.Line
		function = frame.Function
	} else {
		_, file, line, ok = runtime.Caller(calldepth)
	}

	if ok {
		short := file
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				short = file[i+1:]

				break
			}
		}

		file = short
	} else {
		file = ""
		line = 0
		function = ""
	}

	msg.File = file
	msg.Line = line
	msg.Function = function
}

type MessageLogger interface {
	LogMessage(msg *Message) error
}
