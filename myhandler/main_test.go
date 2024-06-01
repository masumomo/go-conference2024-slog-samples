package main

import (
	"bytes"
	"encoding/json"
	"testing"
	"testing/slogtest"

	myhandler "github.com/masumomo/goconference2024_slog_demos/myhandler/final"
)

func TestMyHandler(t *testing.T) {
	var buf bytes.Buffer
	h := myhandler.New(&buf, nil)
	results := func() []map[string]any {
		ms := []map[string]any{}
		for _, line := range bytes.Split(buf.Bytes(), []byte{'\n'}) {
			if len(line) == 0 {
				continue
			}
			var m map[string]any
			if err := json.Unmarshal(line, &m); err != nil {
				t.Fatal(err)
			}
			ms = append(ms, m)
		}
		return ms
	}
	err := slogtest.TestHandler(h, results)
	if err != nil {
		t.Error(err)
	}
}
