package imago

import (
	"errors"
	"image"
	"image/color"
	"math"

	"github.com/fr3fou/matrigo"
)

// ErrKernelSquare is an error for when the kernel shape isn't a square.
var ErrKernelSquare = errors.New("imago: kernel shape must be square")

func Conv(img image.Image, width, height int, kernel matrigo.Matrix, stride int) image.Image {
	output := image.NewNRGBA(image.Rect(0, 0, width, height))
	bounds := img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y += stride {
		for x := bounds.Min.X; x < bounds.Max.X; x += stride {
			output.Set(x, y, conv(img, kernel, x, y))
		}
	}

	return output
}

func conv(img image.Image, kernel matrigo.Matrix, x, y int) color.Color {
	rSum := 0.0
	gSum := 0.0
	bSum := 0.0
	aSum := 0.0

	for i := 0; i < kernel.Rows; i++ {
		for j := 0; j < kernel.Columns; j++ {
			r, g, b, a := img.At(x, y).RGBA()
			filter := kernel.Data[i][j]

			rSum += float64(r) * (filter)
			gSum += float64(g) * (filter)
			bSum += float64(b) * (filter)
			aSum += float64(a) * (filter)
		}
	}

	return color.RGBA{
		R: uint8(math.Round(rSum) / 256),
		G: uint8(math.Round(gSum) / 256),
		B: uint8(math.Round(bSum) / 256),
		A: uint8(math.Round(aSum) / 256),
	}
}
