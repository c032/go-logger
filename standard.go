package logger

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"runtime"
	"sort"
	"sync"
	"time"
)

var (
	_ Logger = (*standardWrapper)(nil)
)

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

type slogWrapper struct {
	mu     sync.Mutex
	ctx    context.Context
	logger *slog.Logger
}

func (sw *slogWrapper) Print(v ...interface{}) {
	sw.mu.Lock()
	defer sw.mu.Unlock()

	h := sw.logger.Handler()

	t := time.Now().UTC()
	level := slog.LevelInfo
	msg := fmt.Sprint(v...)

	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])

	r := slog.NewRecord(t, level, msg, pcs[0])
	_ = h.Handle(sw.ctx, r)
}

func (sw *slogWrapper) Printf(format string, v ...interface{}) {
	sw.mu.Lock()
	defer sw.mu.Unlock()

	h := sw.logger.Handler()

	t := time.Now().UTC()
	level := slog.LevelInfo
	msg := fmt.Sprintf(format, v...)

	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])

	r := slog.NewRecord(t, level, msg, pcs[0])
	_ = h.Handle(sw.ctx, r)
}

func (sw *slogWrapper) Error(v ...interface{}) {
	sw.mu.Lock()
	defer sw.mu.Unlock()

	h := sw.logger.Handler()

	t := time.Now().UTC()
	level := slog.LevelError
	msg := fmt.Sprint(v...)

	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])

	r := slog.NewRecord(t, level, msg, pcs[0])
	_ = h.Handle(sw.ctx, r)
}

func (sw *slogWrapper) Errorf(format string, v ...interface{}) {
	sw.mu.Lock()
	defer sw.mu.Unlock()

	h := sw.logger.Handler()

	t := time.Now().UTC()
	level := slog.LevelError
	msg := fmt.Sprintf(format, v...)

	var pcs [1]uintptr
	runtime.Callers(2, pcs[:])

	r := slog.NewRecord(t, level, msg, pcs[0])
	_ = h.Handle(sw.ctx, r)
}

func (sw *slogWrapper) WithFields(fields Fields) Logger {
	sw.mu.Lock()
	defer sw.mu.Unlock()

	l := sw.logger

	var keys []string
	for k, _ := range fields {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := fields[k]

		l = l.With(k, v)
	}

	return FromSlog(l)
}

func FromSlog(slogger *slog.Logger) Logger {
	return &slogWrapper{
		ctx:    context.Background(),
		logger: slogger,
	}
}
