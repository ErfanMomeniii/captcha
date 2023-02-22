package captcha_test

import (
	"github.com/ErfanMomeniii/captcha.git"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Numeric_Captcha(t *testing.T) {
	c := captcha.New(300, 400, 40)

	_, err := c.Numeric(6)
	assert.NoError(t, err)

	_, err = c.Numeric(100)
	assert.NoError(t, err)
}

func Test_Alphabetical_Captcha(t *testing.T) {
	c := captcha.New(300, 400, 40)

	_, err := c.Alphabetical(6)
	assert.NoError(t, err)

	_, err = c.Alphabetical(100)
	assert.NoError(t, err)
}

func Test_Mixed_Captcha(t *testing.T) {
	c := captcha.New(300, 400, 40)

	_, err := c.Mixed(6)
	assert.NoError(t, err)

	_, err = c.Mixed(100)
	assert.NoError(t, err)
}

func Test_Save_Image(t *testing.T) {
	c := captcha.New(300, 400, 40)

	im, err := c.Numeric(6)
	assert.NoError(t, err)
	err = c.Save("./test/a.png", im)
	assert.NoError(t, err)

	im, err = c.Alphabetical(6)
	assert.NoError(t, err)
	err = c.Save("./test/b.png", im)
	assert.NoError(t, err)

	im, err = c.Mixed(6)
	assert.NoError(t, err)
	err = c.Save("./test/c.png", im)
	assert.NoError(t, err)

	err = c.Save("./test/d", im)
	assert.NoError(t, err)
}
