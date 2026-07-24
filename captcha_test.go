package captcha_test

import (
	"testing"

	"github.com/ErfanMomeniii/captcha/v2"
	"github.com/stretchr/testify/assert"
)

func Test_Numeric_Captcha(t *testing.T) {
	c := captcha.New()
	r, err := c.Numeric(6)
	assert.NoError(t, err)
	assert.Len(t, r.Text, 6)

	_, err = c.Numeric(100)
	assert.NoError(t, err)
}

func Test_Alphabetical_Captcha(t *testing.T) {
	c := captcha.New()
	r, err := c.Alphabetical(6)
	assert.NoError(t, err)
	assert.Len(t, r.Text, 6)
}

func Test_Mixed_Captcha(t *testing.T) {
	c := captcha.New()
	r, err := c.Mixed(6)
	assert.NoError(t, err)
	assert.Len(t, r.Text, 6)
}

func Test_Save_Image(t *testing.T) {
	c := captcha.New()
	r, err := c.Numeric(6)
	assert.NoError(t, err)
	assert.NoError(t, c.Save("./test/a.png", r.Image))
	assert.NoError(t, c.Save("./test/d", r.Image)) // .png appended
}
