package object

import (
	"bytes"
	_ "embed"
	"image"
	"image/png"
)

type EmbedImage []byte

//go:embed data/gopher.png
var GopherPng EmbedImage

func (e EmbedImage) Image() (image.Image, error) {
	buf := bytes.NewBuffer(GopherPng)
	return png.Decode(buf)
}
