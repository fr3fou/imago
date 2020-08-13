package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/fr3fou/imago/imago"
)

func main() {
	img, err := os.Open(os.Args[1])

	if err != nil {
		panic(err)
	}

	filename := path.Base(os.Args[1])
	ext := path.Ext(filename)
	filename = strings.TrimSuffix(filename, ext)

	amount, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

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

	scaled := imago.NNI(loadedImage, amount)

	scaledFile, err := os.Create(fmt.Sprintf("%s_%vx.png", filename, amount))
	if err != nil {
		panic(err)
	}
	defer scaledFile.Close()

	png.Encode(scaledFile, scaled)
}
