package captcha

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestAudioWAVHeader(t *testing.T) {
	a := NewAudio()
	r, err := a.Generate("1a9")
	if err != nil {
		t.Fatal(err)
	}
	if r.Text != "1a9" {
		t.Fatalf("text=%q want 1a9", r.Text)
	}
	if len(r.WAV) < 44 {
		t.Fatalf("WAV too short: %d bytes", len(r.WAV))
	}
	if !bytes.Equal(r.WAV[0:4], []byte("RIFF")) {
		t.Fatalf("missing RIFF marker")
	}
	if !bytes.Equal(r.WAV[8:12], []byte("WAVE")) {
		t.Fatalf("missing WAVE marker")
	}
	// data chunk size in header must match trailing bytes.
	dataSize := binary.LittleEndian.Uint32(r.WAV[40:44])
	if int(dataSize) != len(r.WAV)-44 {
		t.Fatalf("data size %d != payload %d", dataSize, len(r.WAV)-44)
	}
	// sample rate in header must match the default.
	if sr := binary.LittleEndian.Uint32(r.WAV[24:28]); sr != defaultSampleRate {
		t.Fatalf("sample rate=%d want %d", sr, defaultSampleRate)
	}
}

func TestAudioTextIsCaseNormalized(t *testing.T) {
	a := NewAudio()
	r, err := a.Generate("A9z")
	if err != nil {
		t.Fatal(err)
	}
	// Case cannot be conveyed by audio, so Text is the lower-case key.
	if r.Text != "a9z" {
		t.Fatalf("text=%q want a9z", r.Text)
	}
}

func TestAudioRejectsUnknownChar(t *testing.T) {
	a := NewAudio()
	if _, err := a.Generate("!"); err == nil {
		t.Fatal("want error for unsupported character")
	}
}

func TestAudioRejectsEmpty(t *testing.T) {
	a := NewAudio()
	if _, err := a.Generate(""); err == nil {
		t.Fatal("want error for empty text")
	}
}

func TestAudioCustomSampleRate(t *testing.T) {
	a := NewAudio(WithSampleRate(16000), WithTiming(0.2, 0.05))
	r, err := a.Generate("42")
	if err != nil {
		t.Fatal(err)
	}
	if sr := binary.LittleEndian.Uint32(r.WAV[24:28]); sr != 16000 {
		t.Fatalf("sample rate=%d want 16000", sr)
	}
}
