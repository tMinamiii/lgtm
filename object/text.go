package object

import (
	"image"
	"image/color"
	"strings"
	"unicode"
)

const (
	DefaultMainText string = "LGTM"
	DefaultSubText  string = "Looks Good To Me"
)

type PaddingText string

func (p PaddingText) String() string {
	b := &strings.Builder{}
	space := ' '
	for i, v := range p {
		b.WriteRune(v)
		if len(p)-1 == i {
			break
		}
		b.WriteRune(space)
	}
	return b.String()
}

func (p PaddingText) HasJP() bool {
	for _, v := range p {
		if unicode.In(v, unicode.Hiragana, unicode.Katakana, unicode.Han) {
			return true
		}
	}
	return false
}

type TextColor color.Gray16

var (
	TextColorWhite = TextColor(color.White)
	TextColorBlack = TextColor(color.Black)
)

func (t TextColor) Gray16() color.Gray16 {
	return color.Gray16(t)
}

type Point struct {
	X float64
	Y float64
}

type MessageType string

const (
	MessageTypeMain MessageType = "main"
	MessageTypeSub  MessageType = "sub"
)

type Text struct {
	Text        PaddingText
	Font        Font
	MessageType MessageType
	TextColor   TextColor
}

func NewText(text string, messageType MessageType, textColor TextColor) *Text {
	return &Text{
		Text:        PaddingText(text),
		Font:        NotoSansMono,
		MessageType: messageType,
		TextColor:   textColor,
	}
}

func (t *Text) FontSize(img image.Image) float64 {
	imageWidth := img.Bounds().Dx()
	imageHeight := img.Bounds().Dy()
	textLength := len(t.Text.String())

	// セーフエリアを考慮した利用可能エリア
	safeAreaWidth := float64(imageWidth) * 0.9
	safeAreaHeight := float64(imageHeight) * 0.7

	// 幅制約: テキストが収まる最大フォントサイズ
	// 短いテキストにはより寛大な計算を適用
	var maxFontSizeByWidth float64
	if textLength <= 4 {
		// 短いテキスト（4文字以下）: より大きなフォントサイズを許可
		maxFontSizeByWidth = safeAreaWidth / float64(textLength) * 1.5
	} else {
		// 長いテキスト: 従来の計算
		maxFontSizeByWidth = safeAreaWidth / float64(textLength)
	}

	// 高さ制約: メインテキストとサブテキストの両方が収まる最大フォントサイズ
	// メインテキストが大きく、サブテキストがその60%のサイズと仮定
	var maxFontSizeByHeight float64
	if t.MessageType == MessageTypeMain {
		// メインテキスト: 利用可能高さの60%を使用
		maxFontSizeByHeight = safeAreaHeight * 0.6
	} else {
		// サブテキスト: 利用可能高さの36%を使用（メインの60%）
		maxFontSizeByHeight = safeAreaHeight * 0.36
	}

	// 幅と高さの制約の小さい方を採用
	fontSize := maxFontSizeByWidth
	if maxFontSizeByHeight < fontSize {
		fontSize = maxFontSizeByHeight
	}

	// 最小・最大フォントサイズの制限
	minFontSize := 12.0
	maxFontSize := 400.0

	if fontSize < minFontSize {
		fontSize = minFontSize
	} else if fontSize > maxFontSize {
		fontSize = maxFontSize
	}

	return fontSize
}

func (t *Text) Point(img image.Image) *Point {
	imgWidth := img.Bounds().Dx()
	imgHeight := img.Bounds().Dy()
	aspectRatio := float64(imgWidth) / float64(imgHeight)

	// アスペクト比に応じて縦方向のマージンを調整
	var marginY float64
	if aspectRatio > 2.0 {
		// 横長画像: 縦方向のマージンを大きく
		marginY = float64(imgHeight) * 0.15
	} else if aspectRatio < 0.5 {
		// 縦長画像: 縦方向のマージンを小さく
		marginY = float64(imgHeight) * 0.05
	} else {
		// 通常の画像: 標準マージン
		marginY = float64(imgHeight) * 0.1
	}

	// 利用可能エリア
	safeHeight := float64(imgHeight) - (marginY * 2)

	// X座標は常に中央
	x := float64(imgWidth) / 2

	switch t.MessageType {
	case MessageTypeMain:
		// メインテキスト: アスペクト比に応じて位置調整
		var yRatio float64
		if aspectRatio > 2.0 {
			// 横長画像: より中央寄りに
			yRatio = 0.4
		} else if aspectRatio < 0.5 {
			// 縦長画像: より上部に
			yRatio = 0.25
		} else {
			// 通常の画像: 上部1/3
			yRatio = 1.0 / 3.0
		}
		y := marginY + (safeHeight * yRatio)
		return &Point{X: x, Y: y}

	case MessageTypeSub:
		// サブテキスト: アスペクト比に応じて位置調整
		var yRatio float64
		if aspectRatio > 2.0 {
			// 横長画像: より中央寄りに
			yRatio = 0.6
		} else if aspectRatio < 0.5 {
			// 縦長画像: より下部に
			yRatio = 0.75
		} else {
			// 通常の画像: 下部1/3
			yRatio = 2.0 / 3.0
		}
		y := marginY + (safeHeight * yRatio)
		return &Point{X: x, Y: y}
	}

	return &Point{}
}
