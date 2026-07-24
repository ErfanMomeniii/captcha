package captcha

import (
	"sync"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

// The bundled Go Regular font is parsed once and reused across every draw call;
// parsing is comparatively expensive and the font never changes at runtime.
var (
	fontOnce sync.Once
	baseFont *truetype.Font
	fontErr  error
)

// loadFont returns the shared parsed font, parsing it on first use.
func loadFont() (*truetype.Font, error) {
	fontOnce.Do(func() {
		baseFont, fontErr = truetype.Parse(goregular.TTF)
	})
	return baseFont, fontErr
}
