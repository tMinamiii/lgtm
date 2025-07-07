package lgtm

import (
	_ "embed"
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/pkg/errors"
)

type TextDrawer struct {
	MainText   *Text
	SubText    *Text
	InputPath  string
	OutputPath string
}

func NewTextDrawer(main, sub *Text, inputPath, outputPath string) Drawer {
	return &TextDrawer{
		MainText:   main,
		SubText:    sub,
		InputPath:  inputPath,
		OutputPath: outputPath,
	}
}

func (t *TextDrawer) Draw() error {
	ext, err := t.extension(t.InputPath)
	if err != nil {
		return err
	}

	if ext == "gif" {
		return t.drawOnGIF(t.InputPath, t.OutputPath)
	}
	return t.drawOnImage(t.InputPath, t.OutputPath, ext)
}

func (t *TextDrawer) extension(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}

	_, format, err := image.DecodeConfig(f)
	if err != nil {
		return "", err
	}

	return format, nil
}

func (t *TextDrawer) newFilename(inputPath, outputPath, ext string) string {
	if outputPath != "" {
		return outputPath
	}
	filename := filepath.Base(inputPath)
	name := strings.Split(filename, ".")[0]
	suffix := "lgtm"
	return filepath.Join(".", fmt.Sprintf("%s-%s.%s", name, suffix, ext))
}

func (t *TextDrawer) drawOnGIF(inputPath, outputPath string) error {
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	orgGif, err := gif.DecodeAll(file)
	if err != nil {
		return err
	}

	newImage := make([]*image.Paletted, 0, len(orgGif.Image))
	for _, v := range orgGif.Image {
		v := v
		var img image.Image
		img, err = t.embedTexts(v)
		if err != nil {
			return err
		}

		palettedImage := &image.Paletted{
			Pix:     v.Pix,
			Stride:  v.Stride,
			Rect:    v.Bounds(),
			Palette: v.Palette,
		}
		draw.Draw(palettedImage, palettedImage.Rect, img, img.Bounds().Min, draw.Over)
		newImage = append(newImage, palettedImage)
	}
	orgGif.Image = newImage

	out, err := os.Create(t.newFilename(inputPath, outputPath, "gif"))
	if err != nil {
		return err
	}
	defer out.Close()

	if err := gif.EncodeAll(out, orgGif); err != nil {
		return err
	}

	return nil
}

func (t *TextDrawer) drawOnImage(inputPath, outputPath, ext string) error {
	img, err := imaging.Open(inputPath, imaging.AutoOrientation(true))
	if err != nil {
		return err
	}

	img, err = t.embedTexts(img)
	if err != nil {
		return err
	}

	if err := imaging.Save(img, t.newFilename(inputPath, outputPath, ext)); err != nil {
		return err
	}

	return nil
}

func (t *TextDrawer) embedTexts(i image.Image) (image.Image, error) {
	img, err := t.embedString(i, t.MainText)
	if err != nil {
		return nil, err
	}

	// サブテキストが空でない場合のみ描画
	if t.SubText.Text.String() != "" {
		img, err = t.embedString(img, t.SubText)
		if err != nil {
			return nil, err
		}
	}

	return img, nil
}

func (t *TextDrawer) embedString(img image.Image, text *Text) (image.Image, error) {
	imgWidth := img.Bounds().Dx()
	imgHeight := img.Bounds().Dy()
	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(img, 0, 0)

	fontSize := text.FontSize(img)
	face, err := text.Font.FontFace(fontSize)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse font %s", err.Error())
	}
	dc.SetFontFace(face)

	dc.SetColor(text.TextColor.Gray16())

	pt := text.Point(img)
	// 1行制限: DrawStringAnchoredを使用して改行を防ぐ
	// 中央揃えで描画（0.5, 0.5 = 中央基準点）
	dc.DrawStringAnchored(text.Text.String(), pt.X, pt.Y, 0.5, 0.5)

	return dc.Image(), nil
}
