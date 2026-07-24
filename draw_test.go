package captcha

import "testing"

func TestDrawRejectsInvalidConfig(t *testing.T) {
	cases := []*Captcha{
		New(WithSize(0, 80), WithFontSize(32)),
		New(WithSize(200, 0), WithFontSize(32)),
		New(WithSize(200, 80), WithFontSize(0)),
	}
	for i, c := range cases {
		if _, err := c.Numeric(4); err == nil {
			t.Fatalf("case %d: want error for invalid config", i)
		}
	}
}

func TestDrawDimensionsAndNoiseLevels(t *testing.T) {
	for _, n := range []Noise{NoiseNone, NoiseLow, NoiseMedium, NoiseHigh} {
		c := New(WithSize(200, 80), WithFontSize(32), WithNoise(n))
		im, err := c.draw("aB3")
		if err != nil {
			t.Fatalf("noise %v: %v", n, err)
		}
		if b := im.Bounds(); b.Dx() != 200 || b.Dy() != 80 {
			t.Fatalf("noise %v: got %dx%d, want 200x80", n, b.Dx(), b.Dy())
		}
	}
}
