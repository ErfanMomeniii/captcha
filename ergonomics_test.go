package captcha

import (
	"bytes"
	"strings"
	"testing"
)

func TestPackageLevelGenerators(t *testing.T) {
	r, err := Numeric(6)
	if err != nil || len(r.Text) != 6 || r.Image == nil {
		t.Fatalf("Numeric: err=%v r=%+v", err, r)
	}
	if _, err := Alphabetical(4); err != nil {
		t.Fatalf("Alphabetical: %v", err)
	}
	if _, err := Mixed(4); err != nil {
		t.Fatalf("Mixed: %v", err)
	}
	if _, err := Custom(4, "ABC"); err != nil {
		t.Fatalf("Custom: %v", err)
	}
	if _, err := Word([]string{"cat"}); err != nil {
		t.Fatalf("Word: %v", err)
	}
	if _, err := Math(); err != nil {
		t.Fatalf("Math: %v", err)
	}
	a, err := Audio("42")
	if err != nil || len(a.WAV) < 44 {
		t.Fatalf("Audio: err=%v len=%d", err, len(a.WAV))
	}
}

func TestMatch(t *testing.T) {
	cases := []struct {
		expected, input string
		want            bool
	}{
		{"AB12", "ab12", true},    // case-insensitive
		{"AB12", "  AB12 ", true}, // trims whitespace
		{"AB12", "AB13", false},
		{"", "AB12", false}, // empty expected never matches
		{"AB12", "", false},
	}
	for _, tc := range cases {
		if got := Match(tc.expected, tc.input); got != tc.want {
			t.Fatalf("Match(%q,%q)=%v want %v", tc.expected, tc.input, got, tc.want)
		}
	}
}

func TestResultOutputs(t *testing.T) {
	r, err := Numeric(5)
	if err != nil {
		t.Fatal(err)
	}

	png, err := r.PNG()
	if err != nil || !bytes.HasPrefix(png, []byte("\x89PNG\r\n\x1a\n")) {
		t.Fatalf("PNG: err=%v prefix ok=%v", err, bytes.HasPrefix(png, []byte("\x89PNG")))
	}

	var buf bytes.Buffer
	n, err := r.WriteTo(&buf)
	if err != nil || int(n) != buf.Len() || buf.Len() != len(png) {
		t.Fatalf("WriteTo: n=%d err=%v buf=%d png=%d", n, err, buf.Len(), len(png))
	}

	uri, err := r.DataURI()
	if err != nil || !strings.HasPrefix(uri, "data:image/png;base64,") {
		t.Fatalf("DataURI: err=%v uri prefix=%q", err, uri[:min(30, len(uri))])
	}
}
