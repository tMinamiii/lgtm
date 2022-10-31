package lgtm

import (
	"image"
	"unicode"
)

const (
	DefaultMainText = "L G T M"
	DefaultSubText  = "L o o k s   G o o d   T o   M e"
)

type Point struct {
	X float64
	Y float64
}

type MessageType string

const (
	MessageTypeMain MessageType = "main"
	MessageTypeSub  MessageType = "sub"
)

type Text struct {
	Text        string
	MessageType MessageType
}

func NewText(text string, messageType MessageType) *Text {
	return &Text{
		Text:        text,
		MessageType: messageType,
	}
}

func (t *Text) FontSize(img image.Image, text string) float64 {

	switch t.MessageType {
	case MessageTypeMain:
		textLength := len(DefaultMainText)
		if len(text) > len(DefaultMainText) {
			textLength = len(text)
		}
		imageWidth := img.Bounds().Dx()
		if t.hasJP(text) {
			return float64(imageWidth*7) / (6 * float64(textLength) / 1.8)
		}
		return float64(imageWidth*7) / (6 * float64(textLength))
	case MessageTypeSub:
		textLength := len(DefaultSubText)
		if len(text) > len(DefaultSubText) {
			textLength = len(text)
		}
		imageWidth := img.Bounds().Dx()
		if t.hasJP(text) {
			return float64(imageWidth*32) / (22 * float64(textLength) / 1.3)
		}
		return float64(imageWidth*32) / (22 * float64(textLength))
	}
	return 0
}

func (t *Text) hasJP(text string) bool {
	for _, v := range text {
		return unicode.In(v, unicode.Hiragana, unicode.Katakana, unicode.Han)
	}
	return false
}

func (t *Text) Point(img image.Image) *Point {
	switch t.MessageType {
	case MessageTypeMain:
		imgWidth := img.Bounds().Dx()
		imgHeight := img.Bounds().Dy()

		return &Point{
			X: float64(imgWidth) / 2,
			Y: float64(imgHeight)/2 - float64(imgHeight)/10,
		}
	case MessageTypeSub:
		imgWidth := img.Bounds().Dx()
		imgHeight := img.Bounds().Dy()

		return &Point{
			X: float64(imgWidth) / 2,
			Y: float64(imgHeight) - (float64(imgHeight) / 4),
		}
	}
	return &Point{}
}
