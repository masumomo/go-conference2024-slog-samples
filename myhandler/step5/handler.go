package step5

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"sync"
	"time"
)

type MyHandler struct {
	opts   Options
	groups []string
	attrs  []slog.Attr
	out    io.Writer
	mu     *sync.Mutex
}

type Options struct {
	Level    slog.Leveler
	TimeZone *time.Location
}

func New(out io.Writer, opts *Options) *MyHandler {
	h := &MyHandler{out: out, mu: &sync.Mutex{}}
	if opts != nil {
		h.opts = *opts
	}
	if h.opts.Level == nil {
		h.opts.Level = slog.LevelInfo
	}

	if h.opts.TimeZone == nil {
		h.opts.TimeZone = time.FixedZone("Asia/Tokyo", 9*60*60)
	}
	return h
}

func (h *MyHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}

func (h *MyHandler) Handle(ctx context.Context, r slog.Record) error {
	buf := make([]byte, 0, 1024)
	// TODO: Need to handler zero value

	buf = fmt.Appendf(buf, "╔--------------------------------------╗\n")
	buf = fmt.Appendf(buf, " %s%s: %s\n", decolateLogLevel(r.Level), r.Level, r.Message)
	buf = fmt.Appendf(buf, " UTC: %s\n", r.Time.In(time.UTC).Format(time.DateTime))
	buf = fmt.Appendf(buf, " %s: %s\n", h.opts.TimeZone.String(), r.Time.In(h.opts.TimeZone).Format(time.DateTime))
	buf = fmt.Appendf(buf, " -------------------------------------- \n")

	nestLevel := 0

	for _, a := range h.attrs {
		buf = h.appendAttr(buf, a, nestLevel)
	}
	if nestLevel > 0 {
		buf = fmt.Appendf(buf, "%*s└──", (nestLevel-1)*4, "")
	}
	for _, g := range h.groups {
		buf = fmt.Appendf(buf, "  %s\n", g)
		nestLevel++
	}

	r.Attrs(func(a slog.Attr) bool {
		buf = h.appendAttr(buf, a, nestLevel)
		return true
	})

	buf = fmt.Appendf(buf, "╚--------------------------------------╝\n")

	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.out.Write(buf)
	return err
}

func (h *MyHandler) appendAttr(buf []byte, a slog.Attr, nestLevel int) []byte {
	a.Value = a.Value.Resolve()
	buf = fmt.Appendf(buf, "  ")
	if nestLevel > 0 {
		buf = fmt.Appendf(buf, "%*s└──", (nestLevel-1)*4, "")
	}
	switch a.Value.Kind() {
	case slog.KindString:
		buf = fmt.Appendf(buf, "%s: %q\n", a.Key, a.Value.String())
	case slog.KindTime:
		buf = fmt.Appendf(buf, "%s: %s\n", a.Key, a.Value.Time().Format(time.RFC3339))
	case slog.KindGroup:
		attrs := a.Value.Group()
		if a.Key != "" {
			buf = fmt.Appendf(buf, "%s\n", a.Key)
			nestLevel++
		}
		for _, ga := range attrs {
			buf = h.appendAttr(buf, ga, nestLevel)
		}
	// TODO: We need to add more case for all Kind
	default:
		buf = fmt.Appendf(buf, "%s: %s\n", a.Key, a.Value)
	}
	return buf
}

func (h *MyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	if len(attrs) == 0 {
		return h
	}
	return &MyHandler{
		opts:   h.opts,
		out:    h.out,
		groups: h.groups,
		attrs:  append(h.attrs, attrs...),
		mu:     h.mu,
	}
}

func (h *MyHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	return &MyHandler{
		opts:   h.opts,
		out:    h.out,
		groups: append(h.groups, name),
		attrs:  h.attrs,
		mu:     h.mu,
	}
}

func decolateLogLevel(l slog.Level) string {
	switch l {
	case slog.LevelInfo:
		return "🔍"
	case slog.LevelWarn:
		return "🐛"
	case slog.LevelError:
		return "❗"
	default:
		return "🪵"
	}
}
