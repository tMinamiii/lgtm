package object

import (
	"image"
	"image/color"
	"strings"
	"unicode"
)

const (
	DefaultMainText string = "LGTM"
	DefaultSubText  string = "Looks Good To Me"
)

type PaddingText string

func (p PaddingText) String() string {
	b := &strings.Builder{}
	space := ' '
	for i, v := range p {
		b.WriteRune(v)
		if len(p)-1 == i {
			break
		}
		b.WriteRune(space)
	}
	return b.String()
}

func (p PaddingText) HasJP() bool {
	for _, v := range p {
		if unicode.In(v, unicode.Hiragana, unicode.Katakana, unicode.Han) {
			return true
		}
	}
	return false
}

type TextColor color.Gray16

var (
	TextColorWhite = TextColor(color.White)
	TextColorBlack = TextColor(color.Black)
)

func (t TextColor) Gray16() color.Gray16 {
	return color.Gray16(t)
}

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
	Text        PaddingText
	Font        Font
	MessageType MessageType
	TextColor   TextColor
}

func NewText(text string, font Font, messageType MessageType, textColor TextColor) *Text {
	return &Text{
		Text:        PaddingText(text),
		Font:        font,
		MessageType: messageType,
		TextColor:   textColor,
	}
}

func (t *Text) toPT(px float64) float64 {
	return px / (float64(96) / float64(72))
}

func (t *Text) FontSizePt(img image.Image, text PaddingText) float64 {
	imageWidthPx := img.Bounds().Dx()
	textAreaWidthPx := float64(imageWidthPx)
	textLength := len(text.String())
	return t.toPT(textAreaWidthPx) / float64(textLength)
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
