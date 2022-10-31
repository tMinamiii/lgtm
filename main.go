package main

import (
	"flag"
	"log"
	"os"

	drawer "github.com/tMinamiii/lgtm/lgtm"
	"github.com/tMinamiii/lgtm/object"
)

func main() {
	mainText := flag.String("main", object.DefaultMainText, "main text")
	subText := flag.String("sub", object.DefaultSubText, "sub text")
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
		d := drawer.NewGopherDrawer()
		if err := d.Draw(*path); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		return
	}

	textColor := object.TextColorWhite
	if *color == "black" {
		textColor = object.TextColorBlack
	}

	font := getFont(*serif, *line)
	main := object.NewText(*mainText, font, object.MessageTypeMain, textColor)
	sub := object.NewText(*subText, font, object.MessageTypeSub, textColor)

	d := drawer.NewTextDrawer(main, sub)
	if err := d.Draw(*path); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func getFont(isSerif, isLine bool) object.Font {
	switch {
	case isSerif:
		return object.NotoSerifJP
	case isLine:
		return object.LINESeedJP
	default:
		return object.NotoSansJP
	}
}
