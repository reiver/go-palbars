# go-palbars

Package **palbars** provides tools for creating a **palette color bars** visualization for **color palettes**, for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-palbars

[![GoDoc](https://godoc.org/github.com/reiver/go-palbars?status.svg)](https://godoc.org/github.com/reiver/go-palbars)

## Example

Here is an example:

```golang

import "github.com/reiver/go-palbars"

// ...

color0 := color.NRGBA{0x4e,0x38,0x06, 0xFF}
color1 := color.NRGBA{0xf0,0xa5,0x01, 0xFF},
color2 := color.NRGBA{0xea,0xc0,0x16, 0xFF},
color3 := color.NRGBA{0xf0,0x74,0x00, 0xFF},
color4 := color.NRGBA{0x4e,0x38,0x09, 0xFF},

img = palbars.Image800x418(color0, color1, color2, color3, color4) // <---------

file, err = os.Create(filename)
if nil != err {
	return err
}

var pngEncoder png.Encoder = png.Encoder{
	CompressionLevel:png.BestCompression,
}
err := pngEncoder.Encode(file, img)
if nil != err {
	return err
}
```

## Import

To import package **palbars** use `import` code like the follownig:
```
import "github.com/reiver/go-palbars"
```

## Installation

To install package **palbars** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-palbars
```

## Author

Package **palbars** was written by [Charles Iliya Krempeaux](http://changelog.ca)
