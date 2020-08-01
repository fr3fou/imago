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

// Conv applies a convolution filter to an image.
func Conv(img image.Image, kernel matrigo.Matrix) image.Image {
	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y

	// https://adeshpande3.github.io/A-Beginner%27s-Guide-To-Understanding-Convolutional-Neural-Networks-Part-2/
	kernelSize := kernel.Rows
	padding := kernelSize / 2
	outputWidth := width + 2*padding
	outputHeight := height + 2*padding

	padded := image.NewRGBA(image.Rect(0, 0, outputWidth, outputHeight))
	output := image.NewRGBA(img.Bounds())

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			padded.Set(x+padding, y+padding, img.At(x, y))
		}
	}

	transposedKernel := kernel.Transpose()
	for x := bounds.Min.X + padding; x < bounds.Max.X-padding; x++ {
		for y := bounds.Min.Y + padding; y < bounds.Max.Y-padding; y++ {
			output.Set(x-padding, y-padding, conv(padded, transposedKernel, x, y))
		}
	}

	return output
}

func conv(img *image.RGBA, kernel matrigo.Matrix, x, y int) color.Color {
	rows := kernel.Rows
	cols := kernel.Columns
	// size := float64(rows * cols)

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
			r, g, b, a := rgba(img, x+i, y+j)
			filter := kernel.Data[endX+i][endY+j]

			rSum += filter * float64(r)
			gSum += filter * float64(g)
			bSum += filter * float64(b)
			aSum += float64(a)
		}
	}

	c := color.RGBA{
		R: clamp(rSum, 255),
		G: clamp(gSum, 255),
		B: clamp(bSum, 255),
		A: 255,
	}

	return c
}

// rgba gets the pixel value at given coordinates
func rgba(img *image.RGBA, x, y int) (r uint8, g uint8, b uint8, a uint8) {
	rect := img.Rect
	stride := img.Stride
	pixelPosition := (y-rect.Min.Y)*stride + (x-rect.Min.X)*4

	r = img.Pix[pixelPosition+0]
	g = img.Pix[pixelPosition+1]
	b = img.Pix[pixelPosition+2]
	a = img.Pix[pixelPosition+3]

	return
}

func clamp(val float64, max uint8) uint8 {
	val = math.Abs(val)
	switch {
	case uint8(val) > max:
		return max
	default:
		return uint8(val)
	}
}
