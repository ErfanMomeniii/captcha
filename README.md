<p align="center">
<img src="./assets/photo/logo.png" width=50% height=50%>
</p>
<p align="center">
<a href="https://pkg.go.dev/github.com/mehditeymorian/koi/v3?tab=doc"target="_blank">
    <img src="https://img.shields.io/badge/Go-1.19+-00ADD8?style=for-the-badge&logo=go" alt="go version" />
</a>

<img src="https://img.shields.io/badge/license-MIT-magenta?style=for-the-badge&logo=none" alt="license" />
<img src="https://img.shields.io/badge/Version-1.0.0-red?style=for-the-badge&logo=none" alt="version" />
</p>

# Captcha
Captcha is a lightweight and powerful package for generating captcha pictures with defined keys.

# Documentation

## Install

```bash
go get github.com/ErfanMomeniii/captcha
```   

Next, include it in your application:

```bash
import "github.com/ErfanMomeniii/captcha"
``` 

## Quick Start

The following examples demonstrates how to generate captcha image with ideal width, height and font weight:

### 1. Numeric Captcha

```go
package main

import (
	"github.com/ErfanMomeniii/captcha"
)

func main() {
	c := captcha.New(300, 400, 40)

	im, _ := c.Numeric(6)
	// im is a numeric captcha that has a numeric word of 
	// length six in it

	c.save("./", im)
	// it saves image (im) in the input path
}

```

### 2. Alphabetical Captcha

```go
package main

import (
	"github.com/ErfanMomeniii/captcha"
)

func main() {
	c := captcha.New(300, 400, 40)

	im, _ := c.Alphabetical(6)
	// im is a Alphabetical captcha that has an alphabetical word of 
	// length six in it

	c.save("./", im)
	// it saves image (im) in the input path
}

```

### 3. Mixed Captcha

```go
package main

import (
	"github.com/ErfanMomeniii/captcha"
)

func main() {
	c := captcha.New(300, 400, 40)

	im, _ := c.Alphabetical(6)
	// im is a Mixed captcha that has a mixed word of 
	// length six in it (combination of alphabets and numbers)

	_ = c.save("./", im)
	// it saves image (im) in the input path
}

```

:warning: This package uses some defined templates for choosing random template from those(you can those see in [this](./template.go)) that can customize by the
following :

```go
package main

import (
	"github.com/ErfanMomeniii/captcha"
)

func main() {
	captcha.Templates = []captcha.Template{
		{
			Background: "#ffffff",
			Color:      "#000000",
		},
	}
}

```

## Usage
#### func New
```go
func New(width int, height int, fontWeight float64) *Captcha
```
New creates a new instance of Captcha.
#### type Captcha
```go
type Captcha struct {
	Width      int     // Width of the generated image
	Height     int     // Height of the generated image
	FontWeight float64 // FontWeight of captcha word
}
```
Captcha is an instantiation used for defining some information of the captcha image.
#### func (*captcha) Numeric
```go
func (c *Captcha) Numeric(length int) (image.Image, error)
```
Numeric generates numeric captcha image with input length.
#### func (*captcha) Alphabetical
```go
func (c *Captcha) Alphabetical(length int) (image.Image, error)
```
Alphabetical generates alphabetical captcha image with input length.
#### func (*captcha) Mixed
```go
func (c *Captcha) Mixed(length int) (image.Image, error)
```
Mixed generates mixed (combination of alphabetical and numeric words) captcha image with input length. 
#### func (*captcha) Save
```go
func (c *Captcha) Save(path string, im image.Image) error
```
Save saves png image in the input path
#### type Template
```go
type Template struct {
	Background string // Background color of captcha image
	Color      string // Color of the captcha word
}
```
Template is an instantiation that used for defining some templates for captcha image.
#### func RandTemplate
```go
func RandTemplate() Template
```
RandTemplate generates a random template from the template array.
## Examples

Here are several examples of generated captcha images by only using default templates.

[<img src="./assets/photo/example1.png">](assets/photo/example1.png)
[<img src="./assets/photo/example2.png">](assets/photo/example2.png)
[<img src="./assets/photo/example3.png">](assets/photo/example3.png)
[<img src="./assets/photo/example4.png">](assets/photo/example4.png)
[<img src="./assets/photo/example5.png">](assets/photo/example5.png)
[<img src="./assets/photo/example6.png">](assets/photo/example6.png)
[<img src="./assets/photo/example7.png">](assets/photo/example7.png)
[<img src="./assets/photo/example8.png">](assets/photo/example8.png)
[<img src="./assets/photo/example9.png">](assets/photo/example9.png)
[<img src="./assets/photo/example10.png">](assets/photo/example10.png)
[<img src="./assets/photo/example11.png">](assets/photo/example11.png)
[<img src="./assets/photo/example12.png">](assets/photo/example12.png)
