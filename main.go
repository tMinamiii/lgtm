package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/pkg/errors"
	"golang.org/x/image/font/gofont/gobold"
)

type point struct {
	x float64
	y float64
}
type TextDrawer struct {
	mainText  string
	subText   string
	textColor string
}

func (t *TextDrawer) Draw(path string) error {
	ext, err := t.extension(path)
	if err != nil {
		return err
	}

	if ext == "gif" {
		return t.drawOnGIF(path)
	}
	return t.drawOnImage(path, ext)
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

func (t *TextDrawer) newFilename(path, ext string) string {
	filename := filepath.Base(path)
	name := strings.Split(filename, ".")[0]
	return fmt.Sprintf("%s-lgtm.%s", name, ext)
}

func (t *TextDrawer) fontSizeMain(img image.Image, text string) float64 {
	imageWidth := img.Bounds().Dx()
	return float64(imageWidth*7) / (5.5 * float64(len(text)))
}

func (t *TextDrawer) pointMain(img image.Image) point {
	imgWidth := img.Bounds().Dx()
	imgHeight := img.Bounds().Dy()

	return point{
		x: float64(imgWidth) / 2,
		y: float64(imgHeight)/2 - float64(imgHeight)/20,
	}
}

func (t *TextDrawer) fontSizeSub(img image.Image, text string) float64 {
	imageWidth := img.Bounds().Dx()
	return float64(imageWidth*32) / float64(22*len(text))
}

func (t *TextDrawer) pointSub(img image.Image) point {
	imgWidth := img.Bounds().Dx()
	imgHeight := img.Bounds().Dy()

	return point{
		x: float64(imgWidth) / 2,
		y: float64(imgHeight) - (float64(imgHeight) / 3.5),
	}
}

func (t *TextDrawer) drawOnGIF(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	orgGif, err := gif.DecodeAll(file)
	if err != nil {
		return err
	}

	newPaletted := make([]*image.Paletted, 0, len(orgGif.Image))
	for _, v := range orgGif.Image {
		img, err := t.drawText(v, t.mainText, t.fontSizeMain(v, t.mainText), t.pointMain(v))
		if err != nil {
			return err
		}

		img, err = t.drawText(img, t.subText, t.fontSizeSub(v, t.subText), t.pointSub(v))
		if err != nil {
			return err
		}

		palettedImage := image.NewPaletted(img.Bounds(), palette.Plan9)
		draw.Draw(palettedImage, palettedImage.Rect, img, img.Bounds().Min, draw.Over)
		newPaletted = append(newPaletted, palettedImage)
	}
	orgGif.Image = newPaletted

	out, err := os.Create(t.newFilename(path, "gif"))
	if err != nil {
		return err
	}

	if err := gif.EncodeAll(out, orgGif); err != nil {
		return err
	}

	return nil
}

func (t *TextDrawer) drawOnImage(path, ext string) error {
	img, err := imaging.Open(path, imaging.AutoOrientation(true))
	if err != nil {
		return err
	}

	img, err = t.drawText(img, t.mainText, t.fontSizeMain(img, t.mainText), t.pointMain(img))
	if err != nil {
		return err
	}

	img, err = t.drawText(img, t.subText, t.fontSizeSub(img, t.subText), t.pointSub(img))
	if err != nil {
		return err
	}

	if err := imaging.Save(img, t.newFilename(path, ext)); err != nil {
		return err
	}

	return nil
}

func (t *TextDrawer) drawText(img image.Image, text string, fontSize float64, p point) (image.Image, error) {
	imgWidth := img.Bounds().Dx()
	imgHeight := img.Bounds().Dy()
	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(img, 0, 0)

	ft, err := truetype.Parse(gobold.TTF)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse font %s", err.Error())
	}

	face := truetype.NewFace(ft, &truetype.Options{Size: fontSize})
	dc.SetFontFace(face)

	c := func() color.Gray16 {
		if t.textColor == "white" {
			return color.White
		}
		return color.Black
	}()
	dc.SetColor(c)

	maxWidth := func() float64 {
		if imgWidth > 640 {
			return float64(imgWidth) - 60.0
		}
		return float64(imgWidth)
	}()

	dc.DrawStringWrapped(text, p.x, p.y, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)

	return dc.Image(), nil
}

func main() {
	mainText := flag.String("main", "L G T M", "main text")
	subText := flag.String("sub", "L o o k s   G o o d   T o   M e", "sub text")
	path := flag.String("i", "", "image path")
	textColor := flag.String("c", "white", "color 'white' or 'black'")
	flag.Parse()

	if *path == "" {
		log.Fatal("no image path")
		os.Exit(1)
	}

	if *textColor != "white" && *textColor != "black" {
		w := "white"
		textColor = &w
	}

	d := &TextDrawer{
		mainText:  *mainText,
		subText:   *subText,
		textColor: *textColor,
	}

	if err := d.Draw(*path); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
