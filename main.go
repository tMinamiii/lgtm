package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
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

func fontSizeMain(imageWidth int, text string) float64 {
	return float64(imageWidth*7) / float64(6*len(text))
}

func fontSizeSub(imageWidth int, text string) float64 {
	return float64(imageWidth*32) / float64(22*len(text))
}

func drawText(img image.Image, text, textColor string, fontSize, x, y float64) (image.Image, error) {
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
		if textColor == "white" {
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
	dc.DrawStringWrapped(text, x, y, 0.5, 0.5, maxWidth, 1.5, gg.AlignCenter)

	return dc.Image(), nil
}

func drawMainText(img image.Image, text, textColor string) (image.Image, error) {
	imgWidth := img.Bounds().Dx()
	imgHeight := img.Bounds().Dy()
	x := float64(imgWidth / 2)
	y := float64((imgHeight / 2) - (imgHeight / 20))
	return drawText(img, text, textColor, fontSizeMain(imgWidth, text), x, y)
}

func drawSubText(img image.Image, text, textColor string) (image.Image, error) {
	imgWidth := img.Bounds().Dx()
	imgHeight := img.Bounds().Dy()
	x := float64(imgWidth / 2)
	y := float64(imgHeight - (imgHeight / 3))
	return drawText(img, text, textColor, fontSizeSub(imgWidth, text), x, y)
}

func format(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", errors.Wrapf(err, "failed to open file: %s", err.Error())
	}

	_, format, err := image.DecodeConfig(f)
	if err != nil {
		return "", errors.Wrapf(err, "failed to read decode config: %s", err.Error())
	}

	return format, nil
}

func readImage(path string) (image.Image, string, error) {
	img, err := imaging.Open(path, imaging.AutoOrientation(true))
	if err != nil {
		return nil, "", errors.Wrapf(err, "failed to open image: %s", err.Error())
	}
	ext, err := format(path)
	if err != nil {
		return nil, "", err
	}
	return img, ext, nil
}

func writeImage(img image.Image, ext, path string) error {
	filename := filepath.Base(path)
	name := strings.Split(filename, ".")[0]
	newFilename := fmt.Sprintf("%s-lgtm.%s", name, ext)

	if err := imaging.Save(img, newFilename); err != nil {
		return errors.Wrapf(err, "failed to save image: %s", err.Error())
	}
	return nil
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

	img, ext, err := readImage(*path)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	img, err = drawMainText(img, *mainText, *textColor)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	img, err = drawSubText(img, *subText, *textColor)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	err = writeImage(img, ext, *path)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
