package drawer

import (
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/tMinamiii/lgtm/object"
)

type GopherDrawer struct {
}

func NewGopherDrawer() Drawer {
	return &GopherDrawer{}
}

func (t *GopherDrawer) Draw(path string) error {
	ext, err := t.extension(path)
	if err != nil {
		return err
	}

	if ext == "gif" {
		return t.drawOnGIF(path)
	}
	return t.drawOnImage(path, ext)
}

func (t *GopherDrawer) extension(path string) (string, error) {
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

func (t *GopherDrawer) newFilename(path, ext string) string {
	filename := filepath.Base(path)
	name := strings.Split(filename, ".")[0]
	suffix := "gopher"
	return fmt.Sprintf("%s-%s.%s", name, suffix, ext)
}

func (t *GopherDrawer) drawOnGIF(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	orgGif, err := gif.DecodeAll(file)
	if err != nil {
		return err
	}

	newImage := make([]*image.Paletted, 0, len(orgGif.Image))
	for i, v := range orgGif.Image {
		v := v
		var img image.Image
		img, err = t.embedGopher(v, i%2 == 0)
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

	out, err := os.Create(t.newFilename(path, "gif"))
	if err != nil {
		return err
	}
	defer out.Close()

	if err := gif.EncodeAll(out, orgGif); err != nil {
		return err
	}

	return nil
}

func (t *GopherDrawer) drawOnImage(path, ext string) error {
	img, err := imaging.Open(path, imaging.AutoOrientation(true))
	if err != nil {
		return err
	}

	img, err = t.embedGopher(img, false)
	if err != nil {
		return err
	}

	if err := imaging.Save(img, t.newFilename(path, ext)); err != nil {
		return err
	}

	return nil
}

func (t *GopherDrawer) embedGopher(src image.Image, shake bool) (image.Image, error) {
	gopher, err := object.GopherPng.Image()
	if err != nil {
		return nil, err
	}

	// if gopher image is larger than src image, resize gopher image to half size.
	if src.Bounds().Dx() <= gopher.Bounds().Dx() || src.Bounds().Dy() <= gopher.Bounds().Dy() {
		gopher = imaging.Resize(gopher, gopher.Bounds().Dx()/2, gopher.Bounds().Dy()/2, imaging.NearestNeighbor)
	}

	x := -((src.Bounds().Dx() - gopher.Bounds().Dx()) / 2)
	y := -(src.Bounds().Dy() - gopher.Bounds().Dy()) / 2
	if shake {
		x -= 3
	}

	center := image.Point{x, y}
	newImg := image.NewRGBA(src.Bounds())
	draw.Draw(newImg, newImg.Bounds(), src, image.Point{0, 0}, draw.Src)
	draw.Draw(newImg, newImg.Bounds(), gopher, center, draw.Over)

	return newImg, nil
}
