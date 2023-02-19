package captcha

import (
	"math/rand"
	"time"
)

type Template struct {
	Background string
	Color      string
}

var Templates = []Template{
	{
		Background: "#ffffff",
		Color:      "#8aaae5",
	}, {
		Background: "#fBf8Be",
		Color:      "#234e70",
	}, {
		Background: "#000000",
		Color:      "#f3ca20",
	}, {
		Background: "d9a5b3",
		Color:      "#1868ae",
	}, {
		Background: "fff1e1",
		Color:      "1d3c45",
	}, {
		Background: "#ced7d8",
		Color:      "#f47a60",
	}, {
		Background: "#ffffff",
		Color:      "#cc9999",
	},
}

func RandTemplate() Template {
	rand.Seed(time.Now().UnixNano())

	return Templates[rand.Intn(len(Templates))]
}
