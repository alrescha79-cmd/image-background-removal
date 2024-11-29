package utilities

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func SaveImageToFile(inputFileName string, canvas *image.RGBA, outFolder string) error {
	nameNoExtension := strings.TrimSuffix(inputFileName, filepath.Ext(inputFileName))

	outputPath := fmt.Sprintf("%s/%s-no-bg.png", outFolder, nameNoExtension)

	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	if err := png.Encode(outFile, canvas); err != nil {
		return err
	}

	return nil
}

func Transform(baseImage *os.File, threshold uint8, mode interface{}) *image.RGBA {
	imageData, _, err := image.Decode(baseImage)
	if err != nil {
		panic(err)
	}

	rectangle := imageData.Bounds()
	canvas := image.NewRGBA(rectangle)

	for x := 0; x < rectangle.Dx(); x++ {
		for y := 0; y < rectangle.Dy(); y++ {
			oldPixel := imageData.At(x, y)
			r, g, b, _ := oldPixel.RGBA()

			grayscale := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			luminance := uint8(grayscale / 256)

			setAlpha(canvas, threshold, luminance, x, y, oldPixel, mode)
		}
	}

	return canvas
}

func setAlpha(canvas *image.RGBA, threshold uint8, luminance uint8, x int, y int, oldPixel color.Color, mode interface{}) {
	if luminance > threshold {
		canvas.SetRGBA(x, y, color.RGBA{255, 255, 255, 0})
		return
	}

	if mode == "keep" {
		canvas.Set(x, y, oldPixel)
		return
	}

	canvas.SetRGBA(x, y, color.RGBA{0, 0, 0, 255})
}
