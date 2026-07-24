package captcha

// Noise controls how much distortion is applied to a captcha image.
type Noise int

const (
	NoiseNone   Noise = iota // rotation + baseline only (original look)
	NoiseLow                 // + light dots
	NoiseMedium              // + per-char jitter and extra lines
	NoiseHigh                // + dense dots and stronger jitter
)

// Option configures a Captcha.
type Option func(*Captcha)

// WithNoise sets the distortion level.
func WithNoise(n Noise) Option {
	return func(c *Captcha) { c.Noise = n }
}

// WithSize sets the image dimensions, in pixels.
func WithSize(width, height int) Option {
	return func(c *Captcha) {
		c.Width = width
		c.Height = height
	}
}

// WithFontSize sets the font size of the captcha text.
func WithFontSize(size float64) Option {
	return func(c *Captcha) { c.FontWeight = size }
}
