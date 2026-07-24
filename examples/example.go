package main

import (
	"fmt"
	"os"

	"github.com/ErfanMomeniii/captcha/v2"
)

func main() {
	// Simplest path: zero config. r.Text is the answer, r.Image is the picture.
	r, err := captcha.Numeric(6)
	if err != nil {
		panic(err)
	}
	fmt.Println("answer:", r.Text)

	// Verify a user's submission (case- and whitespace-insensitive).
	fmt.Println("correct?", captcha.Match(r.Text, " "+r.Text+" "))

	// Need customization? Build a generator with options and save a PNG.
	c := captcha.New(captcha.WithSize(320, 120), captcha.WithNoise(captcha.NoiseHigh))
	m, _ := c.Math()
	if err := c.Save("./out.png", m.Image); err != nil {
		panic(err)
	}
	uri, _ := m.DataURI() // ready for an HTML <img src="...">
	fmt.Println("math answer:", m.Text, "-", uri[:32], "...")

	// Audio captcha as a WAV.
	a, _ := captcha.Audio("1a9")
	_ = os.WriteFile("./out.wav", a.WAV, 0o644)
}
