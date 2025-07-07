package lgtm

import (
	"image"
	"image/color"
	"strings"
	"unicode"
	
	"golang.org/x/image/font"
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

func NewMainText(text string, textColor TextColor) *Text {
	return &Text{
		Text:        PaddingText(text),
		Font:        NotoSansMono,
		MessageType: MessageTypeMain,
		TextColor:   textColor,
	}
}

func NewSubText(text string, textColor TextColor) *Text {
	return &Text{
		Text:        PaddingText(text),
		Font:        NotoSansMono,
		MessageType: MessageTypeSub,
		TextColor:   textColor,
	}
}

func (t *Text) FontSize(img image.Image) float64 {
	imageWidth := img.Bounds().Dx()
	imageHeight := img.Bounds().Dy()
	aspectRatio := float64(imageWidth) / float64(imageHeight)

	// セーフエリアを考慮した利用可能エリア（より保守的に設定）
	safeAreaWidth := float64(imageWidth) * 0.85
	safeAreaHeight := float64(imageHeight) * 0.8

	// 最小・最大フォントサイズの制限
	minFontSize := 8.0
	maxFontSize := 400.0

	// 最適なフォントサイズを二分探索で見つける
	bestFontSize := minFontSize
	left, right := minFontSize, maxFontSize
	
	for right-left > 0.5 {
		mid := (left + right) / 2
		
		// このフォントサイズでテキストが収まるかチェック
		if t.textFitsWithFontSize(mid, safeAreaWidth, safeAreaHeight, aspectRatio) {
			bestFontSize = mid
			left = mid
		} else {
			right = mid
		}
	}

	// 最小フォントサイズでも収まらない場合の最終チェック
	if !t.textFitsWithFontSize(bestFontSize, safeAreaWidth, safeAreaHeight, aspectRatio) {
		// 最小フォントサイズでも収まらない場合は、幅だけを考慮した最小フォントサイズを計算
		face, err := t.Font.FontFace(minFontSize)
		if err == nil {
			textWidth := t.measureTextWidth(face)
			if textWidth > safeAreaWidth {
				// 幅に合わせてフォントサイズを計算
				scaleFactor := safeAreaWidth / textWidth
				bestFontSize = minFontSize * scaleFactor
				if bestFontSize < 6.0 {
					bestFontSize = 6.0 // 絶対最小値
				}
			}
		}
	}

	return bestFontSize
}

// measureTextWidth はテキストの実際の幅を測定する
func (t *Text) measureTextWidth(face font.Face) float64 {
	textWidth := 0.0
	// 平均的な文字幅として'M'の幅を取得
	mAdvance, hasMAdvance := face.GlyphAdvance('M')
	
	for _, r := range t.Text.String() {
		advance, ok := face.GlyphAdvance(r)
		if !ok {
			// グリフが見つからない場合は'M'の幅を使用、それも無い場合は固定値
			if hasMAdvance {
				advance = mAdvance
			} else {
				// 固定のfallback値（フォントサイズに比例した値）
				metrics := face.Metrics()
				advance = metrics.Height / 2 // 高さの半分を幅として使用
			}
		}
		textWidth += float64(advance) / 64.0
	}
	return textWidth
}

// textFitsWithFontSize は指定したフォントサイズでテキストが収まるかどうかを判定
func (t *Text) textFitsWithFontSize(fontSize, safeAreaWidth, safeAreaHeight, aspectRatio float64) bool {
	// フォントフェイスを作成してテキストの実際の幅を計算
	face, err := t.Font.FontFace(fontSize)
	if err != nil {
		return false
	}

	// テキストの実際の幅を計算
	textWidth := t.measureTextWidth(face)
	
	// 幅の制約チェック（少し余裕を持たせる）
	if textWidth > safeAreaWidth * 0.98 {
		return false
	}

	// 高さの制約チェック: メインテキストとサブテキストの両方が収まる必要がある
	metrics := face.Metrics()
	textHeight := float64(metrics.Height) / 64.0
	
	// メインテキストとサブテキストの間隔を考慮
	textSpacing := textHeight * 0.5
	
	// 縦長画像の場合はより多くのスペースを確保
	var totalTextHeight float64
	if aspectRatio < 0.5 {
		// 非常に縦長な画像: より大きな間隔を確保
		totalTextHeight = textHeight + (textHeight * 0.7) + (textSpacing * 2.5)
	} else if aspectRatio < 0.8 {
		// 縦長画像: メインテキストとサブテキストの高さ + 間隔を考慮
		totalTextHeight = textHeight + (textHeight * 0.7) + (textSpacing * 1.5)
	} else {
		// 横長・正方形画像: 標準的な配置
		totalTextHeight = textHeight + (textHeight * 0.6) + textSpacing
	}

	// 利用可能高さ以内に収まるかチェック（縦長画像では保守的に）
	var heightRatio float64
	if aspectRatio < 0.5 {
		// 非常に縦長な画像: より保守的に
		heightRatio = 0.6
	} else if aspectRatio < 0.8 {
		// 縦長画像: やや保守的に
		heightRatio = 0.65
	} else {
		// 横長・正方形画像: 標準
		heightRatio = 0.7
	}
	
	if totalTextHeight > safeAreaHeight * heightRatio {
		return false
	}

	return true
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

	// フォントサイズとテキスト高さを取得
	fontSize := t.FontSize(img)
	face, err := t.Font.FontFace(fontSize)
	if err != nil {
		// フォント作成エラーの場合は従来のロジックを使用
		return t.fallbackPoint(img, marginY, safeHeight, x, aspectRatio)
	}

	metrics := face.Metrics()
	textHeight := float64(metrics.Height) / 64.0

	switch t.MessageType {
	case MessageTypeMain:
		// メインテキスト: 画像の上部に配置
		var yRatio float64
		if aspectRatio > 2.0 {
			// 横長画像: より中央寄りに
			yRatio = 0.35
		} else if aspectRatio < 0.5 {
			// 縦長画像: より上部に（固定位置）
			yRatio = 0.15
		} else {
			// 通常の画像: 上部1/3
			yRatio = 0.3
		}
		y := marginY + (safeHeight * yRatio)
		return &Point{X: x, Y: y}

	case MessageTypeSub:
		// サブテキスト: メインテキストとの重複を確実に避ける
		// メインテキストの情報を直接計算（再帰を避けるため）
		mainText := &Text{
			Text:        t.Text,
			Font:        t.Font,
			MessageType: MessageTypeMain,
			TextColor:   t.TextColor,
		}
		mainFontSize := mainText.FontSize(img)
		mainFace, mainErr := t.Font.FontFace(mainFontSize)
		
		if mainErr != nil {
			// メインテキスト情報取得失敗時は従来ロジック
			return t.fallbackPoint(img, marginY, safeHeight, x, aspectRatio)
		}
		
		mainMetrics := mainFace.Metrics()
		mainTextHeight := float64(mainMetrics.Height) / 64.0
		
		// メインテキストの位置を直接計算（Point()メソッドを呼ばずに）
		var mainYRatio float64
		if aspectRatio > 2.0 {
			mainYRatio = 0.35
		} else if aspectRatio < 0.5 {
			mainYRatio = 0.15
		} else {
			mainYRatio = 0.3
		}
		mainY := marginY + (safeHeight * mainYRatio)
		
		// サブテキストの位置を計算
		var y float64
		if aspectRatio > 2.0 {
			// 横長画像: 固定位置
			y = marginY + (safeHeight * 0.65)
		} else if aspectRatio < 0.5 {
			// 縦長画像: メインテキストの下から十分な間隔を空けて配置
			minSpacing := textHeight * 2.0 // サブテキスト高さの2倍の間隔
			y = mainY + mainTextHeight/2 + minSpacing + textHeight/2
			
			// 画像の下端からも十分な余裕を確保
			maxY := float64(imgHeight) - marginY - textHeight/2
			if y > maxY {
				y = maxY
			}
		} else {
			// 通常の画像: 下部1/3
			y = marginY + (safeHeight * 0.7)
		}
		
		return &Point{X: x, Y: y}
	}

	return &Point{}
}

// fallbackPoint は従来のロジックを使用したポイント計算
func (t *Text) fallbackPoint(img image.Image, marginY, safeHeight, x, aspectRatio float64) *Point {
	switch t.MessageType {
	case MessageTypeMain:
		var yRatio float64
		if aspectRatio > 2.0 {
			yRatio = 0.4
		} else if aspectRatio < 0.5 {
			yRatio = 0.25
		} else {
			yRatio = 1.0 / 3.0
		}
		y := marginY + (safeHeight * yRatio)
		return &Point{X: x, Y: y}

	case MessageTypeSub:
		var yRatio float64
		if aspectRatio > 2.0 {
			yRatio = 0.6
		} else if aspectRatio < 0.5 {
			yRatio = 0.75
		} else {
			yRatio = 2.0 / 3.0
		}
		y := marginY + (safeHeight * yRatio)
		return &Point{X: x, Y: y}
	}

	return &Point{}
}
