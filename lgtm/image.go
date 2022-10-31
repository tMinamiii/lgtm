package lgtm

import (
	"bytes"
	_ "embed"
	"image"
	"image/png"
)

type EmbedImage []byte

//go:embed embed/gopher.png
var GopherPng EmbedImage

func (e EmbedImage) Image() (image.Image, error) {
	buf := bytes.NewBuffer(GopherPng)
	return png.Decode(buf)
}
