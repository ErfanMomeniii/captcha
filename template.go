package captcha

import "math/rand"

// Template defines the color scheme of a captcha image.
type Template struct {
	Background string // Background color of the captcha image
	Color      string // Color of the captcha word
}

var Templates = []Template{
	{Background: "#ffffff", Color: "#8aaae5"},
	{Background: "#fBf8Be", Color: "#234e70"},
	{Background: "#000000", Color: "#f3ca20"},
	{Background: "#d9a5b3", Color: "#1868ae"},
	{Background: "#fff1e1", Color: "#1d3c45"},
	{Background: "#ced7d8", Color: "#f47a60"},
	{Background: "#ffffff", Color: "#cc9999"},
}

// RandTemplate returns a random template from Templates.
func RandTemplate() Template {
	return Templates[rand.Intn(len(Templates))]
}
