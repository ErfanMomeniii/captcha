package main

import (
	"github.com/ErfanMomeniii/captcha"
)

func main() {
	c := captcha.New(300, 400, 40)
	im, _ := c.Numeric(6)
	c.Save("./out.png", im)
}
