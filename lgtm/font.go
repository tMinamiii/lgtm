package lgtm

import (
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Font []byte

var (
	//go:embed embed/NotoSansJP-Bold.otf
	NotoSansJP Font

	//go:embed embed/NotoSerifJP-Bold.otf
	NotoSerifJP Font

	//go:embed embed/LINESeedJP_OTF_Bd.otf
	LINESeedJP Font
)

func (f Font) FontFace(size float64) (font.Face, error) {
	opts := &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingNone,
	}

	otf, err := opentype.Parse(f)
	if err != nil {
		return nil, err
	}
	return opentype.NewFace(otf, opts)
}
