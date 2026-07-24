package captcha

import (
	"errors"
	"image"
	"image/png"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/ErfanMomeniii/randstr"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Defaults used by New when no size/font options are given.
const (
	defaultWidth      = 240
	defaultHeight     = 80
	defaultFontWeight = 36
)

// Captcha holds the rendering configuration for captcha images.
type Captcha struct {
	Width      int     // Width of the generated image
	Height     int     // Height of the generated image
	FontWeight float64 // Font size of the captcha word
	Noise      Noise   // Distortion level
}

// Result carries the drawn key and the rendered image. Store Text server-side
// and compare it against the user's answer.
type Result struct {
	Text  string
	Image image.Image
}

// New creates a Captcha. With no options it uses sensible defaults
// (240x80, font 36); customize with WithSize, WithFontSize, and WithNoise.
//
//	c := captcha.New()                                  // defaults
//	c := captcha.New(captcha.WithNoise(captcha.NoiseHigh))
//	c := captcha.New(captcha.WithSize(320, 120), captcha.WithFontSize(48))
func New(opts ...Option) *Captcha {
	c := &Captcha{
		Width:      defaultWidth,
		Height:     defaultHeight,
		FontWeight: defaultFontWeight,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// std is the zero-config generator backing the package-level helpers.
var std = New()

// Numeric generates a numeric captcha of the given length using the default
// generator. For custom sizing or noise, use New.
func Numeric(length int) (*Result, error) { return std.Numeric(length) }

// Alphabetical generates a letters-only captcha using the default generator.
func Alphabetical(length int) (*Result, error) { return std.Alphabetical(length) }

// Mixed generates an alphanumeric captcha using the default generator.
func Mixed(length int) (*Result, error) { return std.Mixed(length) }

// Custom generates a captcha from charset using the default generator.
func Custom(length int, charset string) (*Result, error) { return std.Custom(length, charset) }

// Word renders one random word from words using the default generator.
func Word(words []string) (*Result, error) { return std.Word(words) }

// Math generates an arithmetic captcha using the default generator.
func Math() (*Result, error) { return std.Math() }

// Match reports whether the user's input matches the expected key, ignoring
// case and surrounding whitespace. Use it to verify a submitted answer against
// the Result.Text you stored when the captcha was issued.
func Match(expected, input string) bool {
	expected = strings.TrimSpace(expected)
	return expected != "" && strings.EqualFold(strings.TrimSpace(input), expected)
}

// Numeric generates a numeric captcha of the given length.
func (c *Captcha) Numeric(length int) (*Result, error) {
	return c.render(randstr.Dec(length))
}

// Alphabetical generates a captcha of letters of the given length.
func (c *Captcha) Alphabetical(length int) (*Result, error) {
	return c.render(randstr.String(length, alphabet))
}

// Mixed generates an alphanumeric captcha of the given length.
func (c *Captcha) Mixed(length int) (*Result, error) {
	return c.render(randstr.String(length))
}

// Custom generates a captcha of the given length using only characters from
// charset.
func (c *Captcha) Custom(length int, charset string) (*Result, error) {
	if length <= 0 {
		return nil, errors.New("captcha: length must be positive")
	}
	if charset == "" {
		return nil, errors.New("captcha: charset must not be empty")
	}
	return c.render(randstr.String(length, charset))
}

// Word renders one randomly chosen word from words.
func (c *Captcha) Word(words []string) (*Result, error) {
	if len(words) == 0 {
		return nil, errors.New("captcha: words must not be empty")
	}
	return c.render(words[rand.Intn(len(words))])
}

// render draws text and wraps it in a Result.
func (c *Captcha) render(text string) (*Result, error) {
	im, err := c.draw(text)
	if err != nil {
		return nil, err
	}
	return &Result{Text: text, Image: im}, nil
}

// Save writes a PNG image to path, creating parent directories as needed.
func (c *Captcha) Save(path string, im image.Image) error {
	if !strings.HasSuffix(path, ".png") {
		path += ".png"
	}
	if dir := filepath.Dir(path); dir != "" {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	if err := png.Encode(file, im); err != nil {
		_ = file.Close()
		return err
	}
	return file.Close()
}

// draw renders text onto an image using a random template and the configured
// noise level.
func (c *Captcha) draw(text string) (image.Image, error) {
	if c.Width <= 0 || c.Height <= 0 {
		return nil, errors.New("captcha: width and height must be positive")
	}
	if c.FontWeight <= 0 {
		return nil, errors.New("captcha: font weight must be positive")
	}

	font, err := loadFont()
	if err != nil {
		return nil, err
	}

	tpl := RandTemplate()
	dc := gg.NewContext(c.Width, c.Height)
	face := truetype.NewFace(font, &truetype.Options{Size: c.FontWeight})
	dc.SetFontFace(face)

	// Background.
	dc.DrawRectangle(0, 0, float64(c.Width), float64(c.Height))
	dc.SetHexColor(tpl.Background)
	dc.Fill()

	// Global rotation (original behavior).
	clockwise := 1.0
	if rand.Intn(2) == 0 {
		clockwise = -1
	}
	dc.RotateAbout(gg.Radians(clockwise*7), float64(c.Width)/2, float64(c.Height)/2)

	dc.SetHexColor(tpl.Color)

	if c.Noise >= NoiseMedium {
		c.drawJittered(dc, text)
	} else {
		w, h := dc.MeasureString(text)
		dc.DrawString(text, float64(c.Width)/2-w/2, float64(c.Height)/2+h/2)
	}

	// Baseline strike-through (original behavior).
	w, _ := dc.MeasureString(text)
	dc.DrawLine(float64(c.Width)/2-w/2, float64(c.Height)/2, float64(c.Width)/2+w/2, float64(c.Height)/2)
	dc.Stroke()

	// Reset the transform so noise is scattered in image space, not the
	// rotated coordinate space.
	dc.Identity()
	c.applyNoise(dc, tpl.Color)

	return dc.Image(), nil
}

// drawJittered draws each character with a small random rotation and vertical
// offset.
func (c *Captcha) drawJittered(dc *gg.Context, text string) {
	total, h := dc.MeasureString(text)
	x := float64(c.Width)/2 - total/2
	midY := float64(c.Height)/2 + h/2
	for _, r := range text {
		s := string(r)
		cw, _ := dc.MeasureString(s)
		dy := (rand.Float64()*2 - 1) * h * 0.15
		angle := (rand.Float64()*2 - 1) * 0.35 // radians, ~±20°
		dc.Push()
		dc.RotateAbout(angle, x+cw/2, midY)
		dc.DrawString(s, x, midY+dy)
		dc.Pop()
		x += cw
	}
}

// applyNoise scatters dots and lines according to the noise level.
func (c *Captcha) applyNoise(dc *gg.Context, hexColor string) {
	var dots, lines int
	switch c.Noise {
	case NoiseLow:
		dots, lines = 40, 0
	case NoiseMedium:
		dots, lines = 80, 2
	case NoiseHigh:
		dots, lines = 160, 4
	default:
		return
	}
	dc.SetHexColor(hexColor)
	for i := 0; i < dots; i++ {
		dc.DrawPoint(rand.Float64()*float64(c.Width), rand.Float64()*float64(c.Height), 1+rand.Float64())
		dc.Fill()
	}
	for i := 0; i < lines; i++ {
		dc.SetLineWidth(0.5 + rand.Float64())
		dc.DrawLine(
			rand.Float64()*float64(c.Width), rand.Float64()*float64(c.Height),
			rand.Float64()*float64(c.Width), rand.Float64()*float64(c.Height),
		)
		dc.Stroke()
	}
}
