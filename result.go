package captcha

import (
	"bytes"
	"encoding/base64"
	"image/png"
	"io"
)

// PNG encodes the captcha image as PNG-formatted bytes.
func (r *Result) PNG() ([]byte, error) {
	var buf bytes.Buffer
	if err := png.Encode(&buf, r.Image); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// WriteTo encodes the captcha image as a PNG and writes it to w. It implements
// io.WriterTo, so it plugs directly into an http.ResponseWriter, file, or buffer.
func (r *Result) WriteTo(w io.Writer) (int64, error) {
	b, err := r.PNG()
	if err != nil {
		return 0, err
	}
	n, err := w.Write(b)
	return int64(n), err
}

// DataURI returns the image as a base64-encoded PNG data URI, ready to drop into
// an HTML <img src="..."> attribute.
func (r *Result) DataURI() (string, error) {
	b, err := r.PNG()
	if err != nil {
		return "", err
	}
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(b), nil
}
