package palbars

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/reiver/go-img800x418"
	"github.com/reiver/go-imgdye"
)

// Image800x418 returns a 800×418 palette-color-bars visualization of the color palette given by 'colors' as an image.Image.
func Image800x418(colors ...color.Color) image.Image {

	var numColors int = len(colors)

	switch numColors {
	case 0:
		return image800x418_0()
	case 1:
		return image800x418_1(colors[0])
	default:
		return image800x418(colors...)
	}

}

func image800x418_0() image.Image {
	var c color.Color = color.NRGBA{0x16,0x16,0x1D, 0xFF} // eigengrau

	return image800x418_1(c)
}

func image800x418_1(c color.Color) image.Image {
	var palette color.Palette = color.Palette{c}

	img := img800x418.NewPaletted(palette)
	if nil == img {
		return nil
	}

	imgdye.Dye(img, c)

	return img
}

func image800x418(colors ...color.Color) image.Image {

	img := img800x418.NewPaletted(colors)
	if nil == img {
		return nil
	}

	var numColors int = len(colors)
	var numBars  int = numColors    // we have one bar per color, so the number of bars is equal to the number of colors

	var drawer draw.Drawer = draw.Src

	for index:=0; index<numColors; index++ {

		var c color.Color = colors[index]
		coloredImage := image.Uniform{c}

		destinationBounds := destinationbounds(img800x418.Width, img800x418.Height, numBars, index)

		drawer.Draw(img, destinationBounds, &coloredImage, image.ZP)
	}

	return img
}
