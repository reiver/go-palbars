package palbars

import (
	"image"
)

// destinationbounds return the bounding image.Rectangle for bar 'index' of a 'width' × 'height' color palette bars visualization with 'numBar' bars.
func destinationbounds(width int, height int, numBars int, index int) image.Rectangle {
	if 0 == numBars {
		return destinationbounds(width, height, 1, index)
	}

	var idealBarWidth  int = width / numBars
	var remainderWidth int = width % numBars

	var left  int = index * idealBarWidth
	var right int = left  + idealBarWidth

	// if the remainder-width is 2 pixels, then we add 1 extra pixel width to the far-left bar and the far-right bar.
	{
		if 2 <= remainderWidth {
			var barZeroExtraWidth int = remainderWidth / 2

			// make sure that the far-left bar always starts at the left end of the image.
			if 0 != left {
				left += barZeroExtraWidth
			}
			right += barZeroExtraWidth
		}
	}

	// make sure that the far-right bar always goes to the right end of the image.
	{

		if index == numBars - 1 {
			right = width
		}
	}

	return image.Rectangle{
		Min:image.Point{left,  0},
		Max:image.Point{right, height},
	}
}
