package captcha

import (
	"github.com/ErfanMomeniii/randstr"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"image/png"
	"math/rand"
	"os"
	"time"
)

type Captcha struct {
	Width      int
	Height     int
	FontWeight float64
}

func (c *Captcha) Numeric(length int) (image.Image, error) {
	return draw(randstr.Dec(length), c.Width, c.Height, c.FontWeight)
}

func (c *Captcha) Alphabetical(length int) (image.Image, error) {
	return draw(randstr.String(length, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"),
		c.Width, c.Height, c.FontWeight)
}

func (c *Captcha) Mixed(length int) (image.Image, error) {
	return draw(randstr.String(length), c.Width, c.Height, c.FontWeight)
}

func (c *Captcha) Save(path string, im image.Image) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, im)
}

func draw(text string, width int, height int, fontSize float64) (image.Image, error) {
	rand.Seed(time.Now().UnixNano())

	template := RandTemplate()

	dc := gg.NewContext(width, height)

	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return nil, err
	}

	face := truetype.NewFace(font, &truetype.Options{
		Size: fontSize,
	})
	dc.SetFontFace(face)

	w, h := dc.MeasureString(text)
	dc.DrawRectangle(0, 0, float64(width), float64(height))
	dc.SetHexColor(template.Background)
	dc.Fill()

	clockwise := 1
	if rand.Intn(2)%2 == 0 {
		clockwise = -1
	}

	dc.RotateAbout(gg.Radians(float64(clockwise*7)), float64(width/2), float64(height/2))
	dc.SetHexColor(template.Color)
	dc.DrawString(text, float64(width/2)-w/2, float64(height/2)+h/2)

	dc.DrawLine(float64(width/2)-w/2, float64(height/2), float64(width/2)+w/2, float64(height/2))

	dc.Stroke()

	return dc.Image(), nil
}

func New(width int, height int, fontWeight float64) *Captcha {
	return &Captcha{
		Width:      width,
		Height:     height,
		FontWeight: fontWeight,
	}
}
