package main

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/gobold"
)

func main() {
	file, err := os.Open("sample.png")
	if err != nil {
		log.Fatalf("failed to open file: %s", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	ft, err := truetype.Parse(gobold.TTF)
	if err != nil {
		log.Fatalf("failed to parse font %s", err.Error())
		os.Exit(1)
	}

	img, err := png.Decode(file)
	if err != nil {
		log.Fatalf("failed to decode image: %s", err.Error())
		os.Exit(1)
	}

	dst := image.NewRGBA(img.Bounds())
	draw.Draw(dst, dst.Bounds(), img, image.Point{}, draw.Src)

	text := "LGTM"

}
