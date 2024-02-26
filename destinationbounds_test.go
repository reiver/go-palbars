package palbars

import (
	"testing"

	"image"
)

func TestDestinationBounds(t *testing.T) {

	tests := []struct {
		Width int
		Height int
		NumBars int
		Index int
		Expected image.Rectangle
	}{
		{
			Width: 800,
			Height: 418,
			NumBars: 0,
			Index:0,
			Expected: image.Rectangle{
				Min:image.Point{0,0},
				Max:image.Point{800,418},
			},
		},


		{
			Width: 800,
			Height: 418,
			NumBars: 1,
			Index:0,
			Expected: image.Rectangle{
				Min:image.Point{0,0},
				Max:image.Point{800,418},
			},
		},



		{
			Width: 800,
			Height: 418,
			NumBars: 2,
			Index:0,
			Expected: image.Rectangle{
				Min:image.Point{0,0},
				Max:image.Point{400,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 2,
			Index:1,
			Expected: image.Rectangle{
				Min:image.Point{400,0},
				Max:image.Point{800,418},
			},
		},



		{
			Width: 800,
			Height: 418,
			NumBars: 3,
			Index:0,
			Expected: image.Rectangle{
				Min:image.Point{0,0},
				Max:image.Point{267,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 3,
			Index:1,
			Expected: image.Rectangle{
				Min:image.Point{267,0},
				Max:image.Point{533,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 3,
			Index:2,
			Expected: image.Rectangle{
				Min:image.Point{533,0},
				Max:image.Point{800,418},
			},
		},



		{
			Width: 800,
			Height: 418,
			NumBars: 4,
			Index:0,
			Expected: image.Rectangle{
				Min:image.Point{0,0},
				Max:image.Point{200,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 4,
			Index:1,
			Expected: image.Rectangle{
				Min:image.Point{200,0},
				Max:image.Point{400,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 4,
			Index:2,
			Expected: image.Rectangle{
				Min:image.Point{400,0},
				Max:image.Point{600,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 4,
			Index:3,
			Expected: image.Rectangle{
				Min:image.Point{600,0},
				Max:image.Point{800,418},
			},
		},



		{
			Width: 800,
			Height: 418,
			NumBars: 5,
			Index:0,
			Expected: image.Rectangle{
				Min:image.Point{0,0},
				Max:image.Point{160,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 5,
			Index:1,
			Expected: image.Rectangle{
				Min:image.Point{160,0},
				Max:image.Point{320,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 5,
			Index:2,
			Expected: image.Rectangle{
				Min:image.Point{320,0},
				Max:image.Point{480,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 5,
			Index:3,
			Expected: image.Rectangle{
				Min:image.Point{480,0},
				Max:image.Point{640,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 5,
			Index:4,
			Expected: image.Rectangle{
				Min:image.Point{640,0},
				Max:image.Point{800,418},
			},
		},



		{
			Width: 800,
			Height: 418,
			NumBars: 6,
			Index:0,
			Expected: image.Rectangle{
				Min:image.Point{0,0},
				Max:image.Point{134,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 6,
			Index:1,
			Expected: image.Rectangle{
				Min:image.Point{134,0},
				Max:image.Point{267,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 6,
			Index:2,
			Expected: image.Rectangle{
				Min:image.Point{267,0},
				Max:image.Point{400,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 6,
			Index:3,
			Expected: image.Rectangle{
				Min:image.Point{400,0},
				Max:image.Point{533,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 6,
			Index:4,
			Expected: image.Rectangle{
				Min:image.Point{533,0},
				Max:image.Point{666,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 6,
			Index:5,
			Expected: image.Rectangle{
				Min:image.Point{666,0},
				Max:image.Point{800,418},
			},
		},



		{
			Width: 800,
			Height: 418,
			NumBars: 7,
			Index:0,
			Expected: image.Rectangle{
				Min:image.Point{0,0},
				Max:image.Point{115,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 7,
			Index:1,
			Expected: image.Rectangle{
				Min:image.Point{115,0},
				Max:image.Point{229,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 7,
			Index:2,
			Expected: image.Rectangle{
				Min:image.Point{229,0},
				Max:image.Point{343,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 7,
			Index:3,
			Expected: image.Rectangle{
				Min:image.Point{343,0},
				Max:image.Point{457,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 7,
			Index:4,
			Expected: image.Rectangle{
				Min:image.Point{457,0},
				Max:image.Point{571,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 7,
			Index:5,
			Expected: image.Rectangle{
				Min:image.Point{571,0},
				Max:image.Point{685,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 7,
			Index:6,
			Expected: image.Rectangle{
				Min:image.Point{685,0},
				Max:image.Point{800,418},
			},
		},



		{
			Width: 800,
			Height: 418,
			NumBars: 8,
			Index:0,
			Expected: image.Rectangle{
				Min:image.Point{0,0},
				Max:image.Point{100,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 8,
			Index:1,
			Expected: image.Rectangle{
				Min:image.Point{100,0},
				Max:image.Point{200,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 8,
			Index:2,
			Expected: image.Rectangle{
				Min:image.Point{200,0},
				Max:image.Point{300,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 8,
			Index:3,
			Expected: image.Rectangle{
				Min:image.Point{300,0},
				Max:image.Point{400,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 8,
			Index:4,
			Expected: image.Rectangle{
				Min:image.Point{400,0},
				Max:image.Point{500,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 8,
			Index:5,
			Expected: image.Rectangle{
				Min:image.Point{500,0},
				Max:image.Point{600,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 8,
			Index:6,
			Expected: image.Rectangle{
				Min:image.Point{600,0},
				Max:image.Point{700,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 8,
			Index:7,
			Expected: image.Rectangle{
				Min:image.Point{700,0},
				Max:image.Point{800,418},
			},
		},



		{
			Width: 800,
			Height: 418,
			NumBars: 9,
			Index:0,
			Expected: image.Rectangle{
				Min:image.Point{0,0},
				Max:image.Point{88+4,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 9,
			Index:1,
			Expected: image.Rectangle{
				Min:image.Point{88+4,0},
				Max:image.Point{88+88+4,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 9,
			Index:2,
			Expected: image.Rectangle{
				Min:image.Point{88+88+4,0},
				Max:image.Point{88+88+88+4,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 9,
			Index:3,
			Expected: image.Rectangle{
				Min:image.Point{88+88+88+4,0},
				Max:image.Point{88+88+88+88+4,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 9,
			Index:4,
			Expected: image.Rectangle{
				Min:image.Point{88+88+88+88+4,0},
				Max:image.Point{88+88+88+88+88+4,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 9,
			Index:5,
			Expected: image.Rectangle{
				Min:image.Point{88+88+88+88+88+4,0},
				Max:image.Point{88+88+88+88+88+88+4,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 9,
			Index:6,
			Expected: image.Rectangle{
				Min:image.Point{88+88+88+88+88+88+4,0},
				Max:image.Point{88+88+88+88+88+88+88+4,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 9,
			Index:7,
			Expected: image.Rectangle{
				Min:image.Point{88+88+88+88+88+88+88+4,0},
				Max:image.Point{88+88+88+88+88+88+88+88+4,418},
			},
		},
		{
			Width: 800,
			Height: 418,
			NumBars: 9,
			Index:8,
			Expected: image.Rectangle{
				Min:image.Point{88+88+88+88+88+88+88+88+4,0},
				Max:image.Point{800,418},
			},
		},
	}

	for testNumber, test := range tests {

		actual := destinationbounds(test.Width, test.Height, test.NumBars, test.Index)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual destination-bounds is not what was expected.", testNumber)
			t.Logf("EXPECTED: %#v", expected)
			t.Logf("ACTUAL:   %#v", actual)
			t.Logf("WIDTH x HEIGHT: %d x %d", test.Width, test.Height)
			t.Logf("NUM-BARS: %d", test.NumBars)
			t.Logf("INDEX: %d", test.Index)
			continue
		}
	}
}
