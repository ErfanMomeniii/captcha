package captcha

type Template struct {
	Background string
	Color      string
}

var Templates = []Template{
	{
		Background: "#FFFFFF",
		Color:      "#8AAAE5",
	}, {
		Background: "#FBF8BE",
		Color:      "#234E70",
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
	},
}
