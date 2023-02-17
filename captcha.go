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

func (c *Captcha) Numeric(length int) image.Image {
	return draw(randstr.Dec(length), c.Width, c.Height, c.FontWeight)
}

func (c *Captcha) Alphabetical(length int) image.Image {
	return draw(randstr.String(length, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"),
		c.Width, c.Height, c.FontWeight)
}

func (c *Captcha) Mixed(length int) image.Image {
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

func draw(text string, width int, height int, fontSize float64) image.Image {
	rand.Seed(time.Now().UnixNano())

	dc := gg.NewContext(width, height)

	font, _ := truetype.Parse(goregular.TTF)
	face := truetype.NewFace(font, &truetype.Options{
		Size: fontSize,
	})
	dc.SetFontFace(face)

	w, h := dc.MeasureString(text)
	dc.DrawRoundedRectangle(0, 0, 2*w, 2*h, 20)
	dc.SetHexColor("#ffffff")
	dc.Fill()

	u := 1
	if rand.Intn(2)%2 == 0 {
		u = -1
	}

	dc.RotateAbout(gg.Radians(float64(u*7)), w/2, h/2)
	dc.SetHexColor("#8AAAE5")
	dc.DrawString(text, w/2, h+h/2)

	dc.Stroke()

	return dc.Image()
}

func New(width int, height int, fontWeight float64) *Captcha {
	return &Captcha{
		Width:      width,
		Height:     height,
		FontWeight: fontWeight,
	}
}
