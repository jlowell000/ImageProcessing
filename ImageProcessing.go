package main

import (
	"image"
	"image/png"
	"os"
	"runtime"
)

/* Keep odd for simplicity */
// const kernalSize = 5
const sigma = 1.6
const highThreshold, lowThreshold = 7500, 20000

func main() {

	loadedImage := readFileToImage("images/in.png")

	alteredImage := actOnImagePixel(loadedImage, greyscale, runtime.GOMAXPROCS(0))
	writeImageFile("images/greyscale.png", alteredImage)
	alteredImage = actOnImageKernal(alteredImage, gaussianFilter, 7, runtime.GOMAXPROCS(0)/5)
	writeImageFile("images/gaussian.png", alteredImage)
	alteredImage = actOnImageKernal(alteredImage, sobelFilter, 3, runtime.GOMAXPROCS(0)/5)
	writeImageFile("images/sobel.png", alteredImage)
	alteredImage = actOnImagePixel(alteredImage, doubleThreshold, runtime.GOMAXPROCS(0))
	writeImageFile("images/doubleThreshold.png", alteredImage)
	alteredImage = actOnImagePixel(alteredImage, fillInGaps, runtime.GOMAXPROCS(0))
	writeImageFile("images/fillInGaps.png", alteredImage)
}

func readFileToImage(fileName string) image.Image {
	// Read image from file that already exists
	existingImageFile, _ := os.Open(fileName)
	defer existingImageFile.Close()
	loadedImage, _ := png.Decode(existingImageFile)
	return loadedImage
}

func writeImageFile(fileName string, image image.Image) {
	f, _ := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, image)
}
