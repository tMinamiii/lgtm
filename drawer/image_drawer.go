package drawer

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
	"github.com/tMinamiii/lgtm/object"
)

type TextDrawer struct {
	MainText *object.Text
	SubText  *object.Text
}

func NewTextDrawer(main, sub *object.Text) Drawer {
	return &TextDrawer{
		MainText: main,
		SubText:  sub,
	}
}

func (t *TextDrawer) Draw(inputPath, outputPath string) error {
	ext, err := t.extension(inputPath)
	if err != nil {
		return err
	}

	if ext == "gif" {
		return t.drawOnGIF(inputPath, outputPath)
	}
	return t.drawOnImage(inputPath, outputPath, ext)
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
	return fmt.Sprintf("%s-%s.%s", name, suffix, ext)
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

	img, err = t.embedString(img, t.SubText)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func (t *TextDrawer) embedString(img image.Image, text *object.Text) (image.Image, error) {
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

	maxWidth := func() float64 {
		if imgWidth > 640 {
			return float64(imgWidth) - 60.0
		}
		return float64(imgWidth)
	}()

	pt := text.Point(img)
	dc.DrawStringWrapped(text.Text.String(), pt.X, pt.Y, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)

	return dc.Image(), nil
}
