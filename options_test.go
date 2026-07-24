package captcha

import "testing"

func TestWithNoise(t *testing.T) {
	c := New(WithNoise(NoiseHigh))
	if c.Noise != NoiseHigh {
		t.Fatalf("got %v, want NoiseHigh", c.Noise)
	}
	// Default is NoiseNone (preserves current look).
	if def := New(); def.Noise != NoiseNone {
		t.Fatalf("default got %v, want NoiseNone", def.Noise)
	}
	// WithSize / WithFontSize override the defaults.
	sized := New(WithSize(321, 111), WithFontSize(44))
	if sized.Width != 321 || sized.Height != 111 || sized.FontWeight != 44 {
		t.Fatalf("got %dx%d font %v, want 321x111 font 44", sized.Width, sized.Height, sized.FontWeight)
	}
	// New() applies documented defaults.
	if d := New(); d.Width != 240 || d.Height != 80 || d.FontWeight != 36 {
		t.Fatalf("defaults got %dx%d font %v, want 240x80 font 36", d.Width, d.Height, d.FontWeight)
	}
}
