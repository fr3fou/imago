package imago

import (
	"image/png"
	"os"
	"testing"

	"github.com/fr3fou/matrigo"
	"github.com/stretchr/testify/require"
)

func TestConv(t *testing.T) {
	inputPath := "../_examples/Lenna.png"

	inputFile, err := os.Open(inputPath)
	require.Nil(t, err)

	img, err := png.Decode(inputFile)
	require.Nil(t, err)

	kernel := matrigo.New(3, 3, [][]float64{
		{-1, 1, -1},
		{-1, 1, -1},
		{-1, 1, -1},
	})

	outputFile, err := os.Create("../_examples/Lenna_edge.png")
	require.Nil(t, err)
	defer outputFile.Close()

	err = png.Encode(outputFile, Conv(img, kernel, 1))
	require.Nil(t, err)
}
