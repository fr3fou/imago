package main

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path"

	"github.com/fr3fou/imago/imago"
	"github.com/fr3fou/matrigo"
)

func main() {
	img, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}

	filename := path.Base(os.Args[1])
	ext := path.Ext(filename)

	outputName := os.Args[2]

	defer img.Close()

	var loadedImage image.Image
	if ext == ".jpg" || ext == ".jpeg" {
		loadedImage, err = jpeg.Decode(img)
		if err != nil {
			panic(err)
		}
	} else {
		loadedImage, err = png.Decode(img)
		if err != nil {
			panic(err)
		}
	}

	kernel := matrigo.New(3, 3, [][]float64{
		{-1, -1, -1},
		{-1, 8, -1},
		{-1, -1, -1},
	})

	convolvedFile, err := os.Create(outputName)
	if err != nil {
		panic(err)
	}
	defer convolvedFile.Close()

	err = png.Encode(convolvedFile, imago.Conv(loadedImage, kernel))
	if err != nil {
		panic(err)
	}
}
