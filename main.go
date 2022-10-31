package main

import (
	"flag"
	"log"
	"os"

	"github.com/tMinamiii/lgtm/lgtm"
)

func main() {
	mainText := flag.String("main", lgtm.DefaultMainText, "main text")
	subText := flag.String("sub", lgtm.DefaultSubText, "sub text")
	path := flag.String("i", "", "image file path")
	color := flag.String("c", "white", "color 'white' or 'black'")
	line := flag.Bool("line", false, "LINE seed font")
	serif := flag.Bool("serif", false, "Noto serif font")
	gopher := flag.Bool("gopher", false, "embed gopher")
	flag.Parse()

	if *path == "" {
		log.Fatal("no image path")
		os.Exit(1)
	}

	textColor := lgtm.TextColorWhite
	if *color == "black" {
		textColor = lgtm.TextColorBlack
	}

	main := lgtm.NewText(*mainText, lgtm.MessageTypeMain)
	sub := lgtm.NewText(*subText, lgtm.MessageTypeSub)
	font := getFont(*serif, *line)
	d := lgtm.NewTextDrawer(main, sub, textColor, font, *gopher)

	if err := d.Draw(*path); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func getFont(isSerif, isLine bool) lgtm.Font {
	switch {
	case isSerif:
		return lgtm.NotoSerifJP
	case isLine:
		return lgtm.LINESeedJP
	default:
		return lgtm.NotoSansJP
	}
}
