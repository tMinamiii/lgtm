package lgtm

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
)

type ConcentrationLinesDrawer struct {
	InputPath  string
	OutputPath string
	LineCount  int         // 集中線の本数
	LineColor  color.Color // 線の色
}

func NewConcentrationLinesDrawer(inputPath, outputPath string) Drawer {
	return &ConcentrationLinesDrawer{
		InputPath:  inputPath,
		OutputPath: outputPath,
		LineCount:  200,         // デフォルトの線の本数（密度をさらに上げる）
		LineColor:  color.Black, // デフォルトは黒
	}
}

func (c *ConcentrationLinesDrawer) Draw() error {
	ext, err := c.extension(c.InputPath)
	if err != nil {
		return err
	}

	if ext == "gif" {
		return c.drawOnGIF(c.InputPath, c.OutputPath)
	}
	return c.drawOnImage(c.InputPath, c.OutputPath, ext)
}

func (c *ConcentrationLinesDrawer) extension(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, format, err := image.DecodeConfig(f)
	if err != nil {
		return "", err
	}

	return format, nil
}

func (c *ConcentrationLinesDrawer) newFilename(inputPath, outputPath, ext string) string {
	if outputPath != "" {
		return outputPath
	}
	filename := filepath.Base(inputPath)
	name := strings.Split(filename, ".")[0]
	suffix := "concentration"
	return filepath.Join(".", fmt.Sprintf("%s-%s.%s", name, suffix, ext))
}

func (c *ConcentrationLinesDrawer) drawOnGIF(inputPath, outputPath string) error {
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
		img := c.drawConcentrationLines(v)

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

	out, err := os.Create(c.newFilename(inputPath, outputPath, "gif"))
	if err != nil {
		return err
	}
	defer out.Close()

	if err := gif.EncodeAll(out, orgGif); err != nil {
		return err
	}

	return nil
}

func (c *ConcentrationLinesDrawer) drawOnImage(inputPath, outputPath, ext string) error {
	img, err := imaging.Open(inputPath, imaging.AutoOrientation(true))
	if err != nil {
		return err
	}

	result := c.drawConcentrationLines(img)

	if err := imaging.Save(result, c.newFilename(inputPath, outputPath, ext)); err != nil {
		return err
	}

	return nil
}

func (c *ConcentrationLinesDrawer) drawConcentrationLines(img image.Image) image.Image {
	imgWidth := img.Bounds().Dx()
	imgHeight := img.Bounds().Dy()
	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(img, 0, 0)

	// 画像の中心点
	centerX := float64(imgWidth) / 2
	centerY := float64(imgHeight) / 2

	// 画像の対角線の長さ（線が画像全体をカバーするため）
	maxDistance := math.Sqrt(float64(imgWidth*imgWidth + imgHeight*imgHeight))

	// ランダムシードを初期化
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 角度をランダムに生成
	angles := make([]float64, c.LineCount)
	for i := 0; i < c.LineCount; i++ {
		angles[i] = rng.Float64() * 2 * math.Pi
	}

	for i := 0; i < c.LineCount; i++ {
		angle := angles[i]

		// 外側は必ず画像の端から（maxDistanceを使用）
		outerDistance := maxDistance * 0.6 // 画像の端まで到達
		outerX := centerX + math.Cos(angle)*outerDistance
		outerY := centerY + math.Sin(angle)*outerDistance

		// 内側の長さをランダムに（0.15 ~ 0.35の範囲で変化）
		innerDistanceRatio := 0.15 + rng.Float64()*0.2 // 0.15 ~ 0.35
		innerDistance := maxDistance * innerDistanceRatio
		innerX := centerX + math.Cos(angle)*innerDistance
		innerY := centerY + math.Sin(angle)*innerDistance

		// 三角形の幅を完全にランダムに設定（細くする）
		minWidth := math.Min(float64(imgWidth), float64(imgHeight)) * 0.003
		maxWidth := math.Min(float64(imgWidth), float64(imgHeight)) * 0.015
		baseWidth := minWidth + rng.Float64()*(maxWidth-minWidth)

		// 外側の2点を計算（角度に垂直な方向）
		perpAngle := angle + math.Pi/2
		outerLeft1X := outerX + math.Cos(perpAngle)*baseWidth
		outerLeft1Y := outerY + math.Sin(perpAngle)*baseWidth
		outerLeft2X := outerX - math.Cos(perpAngle)*baseWidth
		outerLeft2Y := outerY - math.Sin(perpAngle)*baseWidth

		// 黒い三角形を描画
		dc.SetColor(c.LineColor)
		dc.NewSubPath()
		dc.MoveTo(innerX, innerY)
		dc.LineTo(outerLeft1X, outerLeft1Y)
		dc.LineTo(outerLeft2X, outerLeft2Y)
		dc.ClosePath()
		dc.Fill()
	}

	return dc.Image()
}
