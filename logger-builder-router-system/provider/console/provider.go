package console

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/example/logger-builder-router-system/logger"
)

type Provider struct {
	out        io.Writer
	timeFormat string
	mu         sync.Mutex
}

func New(timeFormat string) *Provider {
	return &Provider{out: os.Stdout, timeFormat: timeFormat}
}

func NewWithWriter(out io.Writer, timeFormat string) *Provider {
	if out == nil {
		out = io.Discard
	}
	return &Provider{out: out, timeFormat: timeFormat}
}

func (p *Provider) Name() string { return "console" }

func (p *Provider) Write(_ context.Context, entry logger.Entry) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	_, err := fmt.Fprintln(p.out, Format(entry, p.timeFormat))
	return err
}

func (p *Provider) Close() error { return nil }

func Format(entry logger.Entry, timeFormat string) string {
	var b strings.Builder
	b.WriteString(entry.Timestamp.Format(timeFormat))
	b.WriteString(" [")
	b.WriteString(entry.Level.String())
	b.WriteString("] ")
	b.WriteString(entry.Message)
	if len(entry.Fields) == 0 {
		return b.String()
	}

	keys := make([]string, 0, len(entry.Fields))
	for key := range entry.Fields {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	ordered := make(map[string]any, len(entry.Fields))
	for _, key := range keys {
		ordered[key] = entry.Fields[key]
	}
	payload, _ := json.Marshal(ordered)
	b.WriteString(" ")
	b.Write(payload)
	return b.String()
}
