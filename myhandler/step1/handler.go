package step1

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"sync"
	"time"
)

type MyHandler struct {
	opts Options
	out  io.Writer
	mu   *sync.Mutex
}

type Options struct {
	Level slog.Leveler
}

func New(out io.Writer, opts *Options) *MyHandler {
	h := &MyHandler{out: out, mu: &sync.Mutex{}}
	if opts != nil {
		h.opts = *opts
	}
	if h.opts.Level == nil {
		h.opts.Level = slog.LevelInfo
	}
	return h
}

func (h *MyHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}

func (h *MyHandler) Handle(ctx context.Context, r slog.Record) error {
	buf := make([]byte, 0, 1024)
	// TODO: Need to handler zero value

	/** Process Record

	// Record holds time, message, logLebel,
	// and Program Counter
	type Record struct {
		Time time.Time
		Message string
		Level Level
		PC uintptr  // The program counter at the time the record was constructed
	}
	*/
	// Output sample: 2024-05-19 13:55:12 INFO: New payment
	buf = fmt.Appendf(buf, "%s %s: %s\n", r.Time.Format(time.DateTime), r.Level, r.Message)

	// The Attrs function calls argument function
	// on each Attr in the Record.
	// Iteration stops if the func returns false.
	r.Attrs(func(a slog.Attr) bool {
		buf = h.appendAttr(buf, a)
		return true
	})

	// Write the log entry
	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.out.Write(buf)
	return err
}

func (h *MyHandler) appendAttr(buf []byte, a slog.Attr) []byte {
	a.Value = a.Value.Resolve()
	switch a.Value.Kind() {
	case slog.KindString:
		buf = fmt.Appendf(buf, "%s: %q\n", a.Key, a.Value.String())
	case slog.KindTime:
		buf = fmt.Appendf(buf, "%s: %s\n", a.Key, a.Value.Time().Format(time.RFC3339))
	case slog.KindGroup:
		attrs := a.Value.Group()
		if a.Key != "" {
			buf = fmt.Appendf(buf, "%s\n", a.Key)
		}
		for _, ga := range attrs {
			buf = h.appendAttr(buf, ga)
		}
	// TODO: We need to add more case for all Kind
	default:
		buf = fmt.Appendf(buf, "%s: %s\n", a.Key, a.Value)
	}
	return buf
}

func (h *MyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	// Return the handler itself for now
	return h
}

func (h *MyHandler) WithGroup(name string) slog.Handler {
	// Return the handler itself for now
	return h
}
