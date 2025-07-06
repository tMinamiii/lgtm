package object

import (
	_ "embed"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Font []byte

var (
	//go:embed data/NotoSansMono-Bold.otf
	NotoSansMono Font
)

func (f Font) FontFace(size float64) (font.Face, error) {
	opts := &opentype.FaceOptions{
		Size:    size,
		DPI:     96,
		Hinting: font.HintingNone,
	}

	otf, err := opentype.Parse(f)
	if err != nil {
		return nil, err
	}
	return opentype.NewFace(otf, opts)
}
