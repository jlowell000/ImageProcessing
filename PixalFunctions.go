package main

import (
	"image"
	"image/color"
)

func greyscale(p image.Point, imageOld image.Image) color.Color {
	return color.Gray16Model.Convert(imageOld.At(p.X, p.Y))
}

func invertColor(p image.Point, imageOld image.Image) color.Color {
	red, green, blue, alpha := imageOld.At(p.X, p.Y).RGBA()
	return color.RGBA64{255 - uint16(red), 255 - uint16(green), 255 - uint16(blue), uint16(alpha)}
}

func doubleThreshold(p image.Point, imageOld image.Image) color.Color {
	r, _, _, _ := imageOld.At(p.X, p.Y).RGBA()
	if r > highThreshold {
		// fmt.Println(r)
		return color.Gray16Model.Convert(imageOld.At(p.X, p.Y))
	} else if r > lowThreshold {
		return color.Gray16Model.Convert(imageOld.At(p.X, p.Y))
	}

	return color.Gray16{0}
}

func fillInGaps(p image.Point, imageOld image.Image) color.Color {

	pr, _, _, _ := imageOld.At(p.X, p.Y).RGBA()
	if pr > 0 {
		return imageOld.At(p.X, p.Y)
	}

	up := p.X-1 >= imageOld.Bounds().Min.X
	left := p.Y-1 >= imageOld.Bounds().Min.Y
	down := p.X+1 < imageOld.Bounds().Max.X
	right := p.Y+1 < imageOld.Bounds().Max.Y

	var sumR, countR uint32

	cacl := func(p image.Point, imageIn image.Image, sumRIn, countRIn uint32) (sumR, countR uint32) {
		r, _, _, _ := imageIn.At(p.X, p.Y-1).RGBA()
		sumR += r
		if r > 0 {
			countR++
		}
		return sumRIn, countRIn
	}

	if up {
		sumR, countR = cacl(image.Point{p.X - 1, p.Y}, imageOld, sumR, countR)
	}
	if down {
		sumR, countR = cacl(image.Point{p.X + 1, p.Y}, imageOld, sumR, countR)
	}
	if left {
		sumR, countR = cacl(image.Point{p.X, p.Y - 1}, imageOld, sumR, countR)
	}
	if right {
		sumR, countR = cacl(image.Point{p.X, p.Y + 1}, imageOld, sumR, countR)
	}
	if up && left {
		sumR, countR = cacl(image.Point{p.X - 1, p.Y - 1}, imageOld, sumR, countR)
	}
	if up && right {
		sumR, countR = cacl(image.Point{p.X - 1, p.Y + 1}, imageOld, sumR, countR)
	}
	if down && left {
		sumR, countR = cacl(image.Point{p.X + 1, p.Y - 1}, imageOld, sumR, countR)
	}
	if down && right {
		sumR, countR = cacl(image.Point{p.X + 1, p.Y + 1}, imageOld, sumR, countR)
	}
	if countR < 3 {
		return color.Gray16{0}
	}
	return color.Gray16{uint16(sumR / countR)}
}
