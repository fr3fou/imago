package imago

import (
	"image/jpeg"
	"image/png"
	"os"
	"testing"

	"github.com/fr3fou/matrigo"
	"github.com/stretchr/testify/require"
)

func TestConv(t *testing.T) {
	inputPath := "../_examples/cheems.jpg"

	inputFile, err := os.Open(inputPath)
	require.Nil(t, err)

	img, err := jpeg.Decode(inputFile)
	require.Nil(t, err)

	kernel := matrigo.New(3, 3, [][]float64{
		{-1, -1, -1},
		{-1, 8, -1},
		{-1, -1, -1},
	})

	outputFile, err := os.Create("../_examples/cheems_edge.png")
	require.Nil(t, err)
	defer outputFile.Close()

	err = png.Encode(outputFile, Conv(img, kernel))
	require.Nil(t, err)
}
