package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/tMinamiii/lgtm/drawer"
	"github.com/tMinamiii/lgtm/object"
)

var (
	color      string
	fontName   string
	gopher     bool
	outputPath string
)

var rootCmd = &cobra.Command{
	Use:   "lgtm [flags] <image-path>",
	Short: "Embed 'LGTM' text or gopher image on images",
	Long: `LGTM is a CLI tool that embeds "LGTM" text on images with customizable colors and fonts.
It can also embed a gopher image and outputs the result as a JPEG file.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		inputPath := args[0]

		if gopher {
			d := drawer.NewGopherDrawer()
			if err := d.Draw(inputPath, outputPath); err != nil {
				log.Fatal(err)
			}
			return
		}

		textColor := object.TextColorWhite
		if color == "black" {
			textColor = object.TextColorBlack
		}

		font := getFont(fontName)
		main := object.NewText(object.DefaultMainText, font, object.MessageTypeMain, textColor)
		sub := object.NewText(object.DefaultSubText, font, object.MessageTypeSub, textColor)

		d := drawer.NewTextDrawer(main, sub)
		if err := d.Draw(inputPath, outputPath); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.Flags().StringVarP(&color, "color", "c", "white", "text color: 'white' or 'black'")
	rootCmd.Flags().StringVarP(&fontName, "font", "f", "sans", "font type: 'sans' or 'line'")
	rootCmd.Flags().BoolVar(&gopher, "gopher", false, "embed gopher image instead of text")
	rootCmd.Flags().StringVarP(&outputPath, "output", "o", "", "output file path")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getFont(fontName string) object.Font {
	switch fontName {
	case "line":
		return object.LINESeedJP
	case "sans":
		return object.NotoSansMono
	default:
		return object.NotoSansMono
	}
}
