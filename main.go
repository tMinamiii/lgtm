package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"github.com/pkg/errors"
	"golang.org/x/image/font/gofont/gobold"
)

func fontSizeMain(imageWidth int) float64 {
	return float64(imageWidth) / 6
}

func fontSizeSub(imageWidth int) float64 {
	return float64(imageWidth) / 22
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
	return drawText(img, text, textColor, fontSizeMain(imgWidth), x, y)
}

func drawSubText(img image.Image, text, textColor string) (image.Image, error) {
	imgWidth := img.Bounds().Dx()
	imgHeight := img.Bounds().Dy()
	x := float64(imgWidth / 2)
	y := float64(imgHeight - (imgHeight / 3))
	return drawText(img, text, textColor, fontSizeSub(imgWidth), x, y)
}

func readImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open file: %s", err.Error())
	}
	defer file.Close()

	ext := filepath.Ext(path)
	switch ext {
	case ".jpeg", ".jpg":
		img, err := jpeg.Decode(file)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to decode image: %s", err.Error())
		}
		return img, nil
	case ".png":
		img, err := png.Decode(file)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to decode image: %s", err.Error())
		}
		return img, nil
	}
	return nil, fmt.Errorf("invalid image extension = %s", ext)
}

func writeImage(img image.Image, path string) error {
	filename := filepath.Base(path)
	ext := filepath.Ext(path)
	name := strings.Replace(filename, ext, "", 1)
	newFilename := fmt.Sprintf("%s-lgtm%s", name, ext)

	newFile, err := os.Create(newFilename)
	if err != nil {
		return errors.Wrapf(err, "failed to create file: %s", err.Error())
	}
	defer newFile.Close()

	b := bufio.NewWriter(newFile)

	switch ext {
	case ".jpeg", ".jpg":
		if err := jpeg.Encode(b, img, &jpeg.Options{Quality: 100}); err != nil {
			return errors.Wrapf(err, "failed to encode image: %s", err.Error())
		}
		return nil
	case ".png":
		if err := png.Encode(b, img); err != nil {
			return errors.Wrapf(err, "failed to encode image: %s", err.Error())
		}
		return nil
	}
	return fmt.Errorf("invalid image extension = %s", ext)
}

func main() {
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

	img, err := readImage(*path)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	img, err = drawMainText(img, "L G T M", *textColor)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	img, err = drawSubText(img, "L o o k s   G o o d   T o   M e", *textColor)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	err = writeImage(img, *path)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
