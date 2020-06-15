package nni

import (
	"image"
	"math"
)

// NNI takes in an image and scales it using
// Nearest Neighbour Interpolation
func NNI(img image.Image, amount int) image.Image {
	bounds := img.Bounds()

	// Create a new image that is "amount" times bigger than the original one
	scaledImage := image.NewRGBA(image.Rect(0, 0, bounds.Max.X*amount, bounds.Max.Y*amount))
	scaledBounds := scaledImage.Bounds()

	height := scaledBounds.Max.Y
	width := scaledBounds.Max.X

	for y := scaledBounds.Min.Y; y < height; y++ {
		for x := scaledBounds.Min.X; x < width; x++ {
			// \frac{oldWidth}{newWidth} = \frac{newPixel}{currentPixel}
			targetX := int(
				math.Floor(
					float64(x*bounds.Max.X) / float64(width),
				),
			)

			// \frac{oldHeight}{newHeight} = \frac{newPixel}{currentPixel}
			targetY := int(
				math.Floor(
					float64(y*bounds.Max.Y) / float64(height),
				),
			)

			color := img.At(targetX, targetY)
			scaledImage.Set(x, y, color)
		}
	}

	return scaledImage
}
