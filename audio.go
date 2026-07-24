package captcha

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"strings"
)

// audioAlphabet lists every character an AudioCaptcha can synthesize.
const audioAlphabet = "0123456789abcdefghijklmnopqrstuvwxyz"

// Default audio parameters.
const (
	defaultSampleRate   = 8000 // Hz
	defaultCharDuration = 0.35 // seconds of tone per character
	defaultGapDuration  = 0.15 // seconds of silence between characters
)

// AudioCaptcha renders spoken-style audio captchas. It is independent of the
// image Captcha: audio has no width, height, or font — only timing and sample
// rate — so it carries its own configuration.
type AudioCaptcha struct {
	SampleRate   int     // samples per second
	CharDuration float64 // seconds of tone per character
	GapDuration  float64 // seconds of silence between characters
}

// AudioOption configures an AudioCaptcha.
type AudioOption func(*AudioCaptcha)

// WithSampleRate sets the WAV sample rate in Hz.
func WithSampleRate(hz int) AudioOption {
	return func(a *AudioCaptcha) { a.SampleRate = hz }
}

// WithTiming sets the per-character tone and inter-character gap durations,
// in seconds.
func WithTiming(charSeconds, gapSeconds float64) AudioOption {
	return func(a *AudioCaptcha) {
		a.CharDuration = charSeconds
		a.GapDuration = gapSeconds
	}
}

// NewAudio creates an AudioCaptcha with sensible defaults, overridden by opts.
func NewAudio(opts ...AudioOption) *AudioCaptcha {
	a := &AudioCaptcha{
		SampleRate:   defaultSampleRate,
		CharDuration: defaultCharDuration,
		GapDuration:  defaultGapDuration,
	}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

// Audio synthesizes an audio captcha for text using default settings
// (8000 Hz, 0.35s per character, 0.15s gap). For custom timing or sample rate,
// use NewAudio.
func Audio(text string) (*AudioResult, error) { return NewAudio().Generate(text) }

// AudioResult carries the spoken key and a mono 16-bit PCM WAV.
//
// Phase 1: each character maps to a distinct synthetic tone (deterministic,
// dependency-free). Phase 2 (follow-up) replaces tones with embedded
// public-domain voice samples in assets/audio/.
type AudioResult struct {
	Text string
	WAV  []byte
}

// Generate synthesizes a WAV encoding text. Characters are matched
// case-insensitively against the supported alphabet; because case cannot be
// conveyed by audio, Result.Text is the normalized (lower-case) key to verify
// against.
func (a *AudioCaptcha) Generate(text string) (*AudioResult, error) {
	if a.SampleRate <= 0 {
		return nil, errors.New("captcha: audio sample rate must be positive")
	}
	if a.CharDuration <= 0 || a.GapDuration < 0 {
		return nil, errors.New("captcha: invalid audio timing")
	}
	lower := strings.ToLower(text)
	if lower == "" {
		return nil, errors.New("captcha: audio text must not be empty")
	}

	gap := make([]int16, int(a.GapDuration*float64(a.SampleRate)))
	var samples []int16
	for _, r := range lower {
		idx := strings.IndexRune(audioAlphabet, r)
		if idx < 0 {
			return nil, fmt.Errorf("captcha: unsupported audio character %q", string(r))
		}
		freq := 300.0 + float64(idx)*40.0 // 300..1700 Hz across the alphabet
		samples = append(samples, tone(freq, a.CharDuration, a.SampleRate)...)
		samples = append(samples, gap...)
	}
	return &AudioResult{Text: lower, WAV: encodeWAV(samples, a.SampleRate)}, nil
}

// tone generates seconds of a sine wave at freq Hz for the given sample rate.
func tone(freq, seconds float64, sampleRate int) []int16 {
	n := int(seconds * float64(sampleRate))
	out := make([]int16, n)
	const amp = 0.6 * math.MaxInt16
	for i := 0; i < n; i++ {
		t := float64(i) / float64(sampleRate)
		out[i] = int16(amp * math.Sin(2*math.Pi*freq*t))
	}
	return out
}

// encodeWAV wraps mono 16-bit PCM samples in a RIFF/WAVE container. It writes
// directly into a single preallocated buffer rather than streaming per-sample.
func encodeWAV(samples []int16, sampleRate int) []byte {
	const (
		numChannels   = 1
		bitsPerSample = 16
	)
	dataSize := len(samples) * (bitsPerSample / 8)
	byteRate := sampleRate * numChannels * bitsPerSample / 8
	blockAlign := numChannels * bitsPerSample / 8

	buf := make([]byte, 44+dataSize)
	copy(buf[0:4], "RIFF")
	binary.LittleEndian.PutUint32(buf[4:8], uint32(36+dataSize))
	copy(buf[8:12], "WAVE")
	copy(buf[12:16], "fmt ")
	binary.LittleEndian.PutUint32(buf[16:20], 16) // PCM chunk size
	binary.LittleEndian.PutUint16(buf[20:22], 1)  // PCM format
	binary.LittleEndian.PutUint16(buf[22:24], numChannels)
	binary.LittleEndian.PutUint32(buf[24:28], uint32(sampleRate))
	binary.LittleEndian.PutUint32(buf[28:32], uint32(byteRate))
	binary.LittleEndian.PutUint16(buf[32:34], uint16(blockAlign))
	binary.LittleEndian.PutUint16(buf[34:36], bitsPerSample)
	copy(buf[36:40], "data")
	binary.LittleEndian.PutUint32(buf[40:44], uint32(dataSize))

	off := 44
	for _, s := range samples {
		binary.LittleEndian.PutUint16(buf[off:off+2], uint16(s))
		off += 2
	}
	return buf
}
