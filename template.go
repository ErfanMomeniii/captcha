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
		Background: "",
		Color:      "",
	}, {
		Background: "",
		Color:      "",
	}, {
		Background: "",
		Color:      "",
	},
}
