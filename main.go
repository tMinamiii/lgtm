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

	if *gopher {
		drawer := lgtm.NewGopherDrawer()
		if err := drawer.Draw(*path); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		return
	}

	textColor := lgtm.TextColorWhite
	if *color == "black" {
		textColor = lgtm.TextColorBlack
	}

	font := getFont(*serif, *line)
	main := lgtm.NewText(*mainText, font, lgtm.MessageTypeMain, textColor)
	sub := lgtm.NewText(*subText, font, lgtm.MessageTypeSub, textColor)

	drawer := lgtm.NewTextDrawer(main, sub)
	if err := drawer.Draw(*path); err != nil {
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
