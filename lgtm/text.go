package lgtm

import (
	"image"
	"unicode"
)

type Point struct {
	X float64
	Y float64
}

type Text struct {
	Text     string
	FontSize func(img image.Image, text string) float64
	Point    func(img image.Image) *Point
}

func NewText(text string, fontSize func(image.Image, string) float64, point func(image.Image) *Point) *Text {
	return &Text{
		Text:     text,
		FontSize: fontSize,
		Point:    point,
	}
}

var FontSizeMain = func(img image.Image, text string) float64 {
	imageWidth := img.Bounds().Dx()
	if hasJP(text) {
		return float64(imageWidth*7) / (6 * float64(len(text)) / 1.8)
	}
	return float64(imageWidth*7) / (6 * float64(len(text)))
}

var FontSizeSub = func(img image.Image, text string) float64 {
	imageWidth := img.Bounds().Dx()
	if hasJP(text) {
		return float64(imageWidth*32) / (22 * float64(len(text)) / 1.3)
	}
	return float64(imageWidth*32) / (22 * float64(len(text)))
}

func hasJP(text string) bool {
	for _, v := range text {
		return unicode.In(v, unicode.Hiragana, unicode.Katakana, unicode.Han)
	}
	return false
}

var PointMain = func(img image.Image) *Point {
	imgWidth := img.Bounds().Dx()
	imgHeight := img.Bounds().Dy()

	return &Point{
		X: float64(imgWidth) / 2,
		Y: float64(imgHeight)/2 - float64(imgHeight)/10,
	}
}

var PointSub = func(img image.Image) *Point {
	imgWidth := img.Bounds().Dx()
	imgHeight := img.Bounds().Dy()

	return &Point{
		X: float64(imgWidth) / 2,
		Y: float64(imgHeight) - (float64(imgHeight) / 4),
	}
}
