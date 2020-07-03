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

// Conv applies a convolution filter to an image
func Conv(img image.Image, kernel matrigo.Matrix, stride int) image.Image {
	bounds := img.Bounds()
	height := bounds.Max.X
	width := bounds.Max.Y

	// https://adeshpande3.github.io/A-Beginner%27s-Guide-To-Understanding-Convolutional-Neural-Networks-Part-2/
	kernelSize := kernel.Rows
	padding := (kernelSize - 1) / 2
	outputWidth := (width-kernelSize+2*padding)/stride + 1
	outputHeight := (height-kernelSize+2*padding)/stride + 1

	output := image.NewNRGBA(image.Rect(0, 0, outputWidth, outputHeight))

	// Apply padding and use new image

	for x := bounds.Min.X; x < bounds.Max.X; x += stride {
		for y := bounds.Min.Y; y < bounds.Max.Y; y += stride {
			output.Set(x, y, conv(img, kernel, x, y))
		}
	}

	return output
}

func conv(img image.Image, kernel matrigo.Matrix, x, y int) color.Color {
	rows := kernel.Rows
	cols := kernel.Columns

	if rows%2 == 0 || cols%2 == 0 {
		panic("imago: kernel shape must consist only of odd numbers")
	}

	rSum := 0.0
	gSum := 0.0
	bSum := 0.0
	aSum := 0.0

	startX := 0 - cols/2
	startY := 0 - rows/2

	endX := cols / 2
	endY := rows / 2

	for i := startX; i <= endX; i++ {
		for j := startY; j <= endY; j++ {
			r, g, b, a := img.At(x+startX, y+startY).RGBA()
			filter := kernel.Data[i][j]

			rSum += float64(r) * filter
			gSum += float64(g) * filter
			bSum += float64(b) * filter
			aSum += float64(a) * filter
		}
	}

	return color.RGBA{
		R: uint8(math.Round(rSum) / 256),
		G: uint8(math.Round(gSum) / 256),
		B: uint8(math.Round(bSum) / 256),
		A: uint8(math.Round(aSum) / 256),
	}
}
