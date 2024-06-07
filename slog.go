package logger

import (
	"context"
	"fmt"
	"log/slog"
	"runtime"
	"slices"
	"sync"
	"time"
)

var _ Logger = (*slogWrapper)(nil)

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
	slices.Sort(keys)

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
