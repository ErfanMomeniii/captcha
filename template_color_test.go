package captcha

import (
	"fmt"
	"image/color"
	"testing"
)

func TestTemplateColorsParse(t *testing.T) {
	for i, tpl := range Templates {
		for _, hex := range []string{tpl.Background, tpl.Color} {
			if len(hex) == 0 || hex[0] != '#' {
				t.Fatalf("template %d: color %q must start with '#'", i, hex)
			}
			var r, g, b uint8
			if _, err := fmtSscanHex(hex, &r, &g, &b); err != nil {
				t.Fatalf("template %d: color %q unparseable: %v", i, hex, err)
			}
			_ = color.RGBA{r, g, b, 255}
		}
	}
}

// fmtSscanHex parses "#rrggbb".
func fmtSscanHex(s string, r, g, b *uint8) (int, error) {
	var rr, gg, bb int
	n, err := fmt.Sscanf(s, "#%02x%02x%02x", &rr, &gg, &bb)
	*r, *g, *b = uint8(rr), uint8(gg), uint8(bb)
	return n, err
}
