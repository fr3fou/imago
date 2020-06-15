package imago

import (
	"image"
	"image/color"
	"math"
)

// Scale takes in an image and scales it "amount" of times
// using a linear interpolation algorithm
func Scale(img image.Image, amount int) image.Image {
	bounds := img.Bounds()

	// Create a new image that is "amount" times bigger than the original one
	scaledImage := image.NewRGBA(image.Rect(0, 0, bounds.Max.X*amount, bounds.Max.Y*amount))
	scaledBounds := scaledImage.Bounds()

	height := scaledBounds.Max.Y
	width := scaledBounds.Max.X

	for y := scaledBounds.Min.Y; y < height; y++ {
		for x := scaledBounds.Min.X; x < width; x++ {
			targetX := float64(x*bounds.Max.X) / float64(width)
			targetY := float64(y*bounds.Max.Y) / float64(height)

			topLeftX, topLeftY := math.Floor(targetX), math.Floor(targetY)
			topRightX, topRightY := math.Ceil(targetX), math.Floor(targetY)
			botRightX, botRightY := math.Ceil(targetX), math.Ceil(targetY)
			botLeftX, botLeftY := math.Floor(targetX), math.Ceil(targetY)

			topLeftColor := img.At(int(topLeftX), int(topLeftY))
			topRightColor := img.At(int(topRightX), int(topRightY))
			botLeftColor := img.At(int(botLeftX), int(botLeftY))
			botRightColor := img.At(int(botRightX), int(botRightY))

			topColor := lerp(topLeftColor, topRightColor,
				int(math.Abs(topLeftX-topRightX)),
				int(math.Abs(topLeftX-targetX)),
			)

			botColor := lerp(botLeftColor, botRightColor,
				int(math.Abs(botLeftX-botRightX)),
				int(math.Abs(botLeftX-targetX)),
			)

			finalColor := lerp(topColor, botColor,
				int(math.Abs(topLeftY-botLeftY)),
				int(math.Abs(topLeftY-targetY)),
			)

			scaledImage.Set(x, y, finalColor)
		}
	}

	return scaledImage
}

func lerp(firstColor color.Color, secondColor color.Color, totalDist int, targetDist int) color.Color {
	fR, fG, fB, fA := firstColor.RGBA()
	sR, sG, sB, sA := secondColor.RGBA()

	return color.RGBA{
		R: uint8((fR) + uint32(targetDist)*(sR-fR)),
		G: uint8((fG) + uint32(targetDist)*(sG-fG)),
		B: uint8((fB) + uint32(targetDist)*(sB-fB)),
		A: uint8((fA) + uint32(targetDist)*(sA-fA)),
	}
}
