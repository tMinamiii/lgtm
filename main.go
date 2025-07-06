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
	color         string
	gopher        bool
	inputPath     string
	outputPath    string
	customText    string
	customSubText string
)

var rootCmd = &cobra.Command{
	Use:   "lgtm [flags]",
	Short: "Embed custom text or gopher image on images",
	Long: `LGTM is a CLI tool that embeds custom text on images with customizable colors.
It can also embed a gopher image and outputs the result as a JPEG file.
By default, it embeds "LGTM" as main text and "Looks Good To Me" as sub-text.
You can customize both using the --text and --sub-text flags.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
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

		mainText := object.DefaultMainText
		subText := object.DefaultSubText

		if customText != "" {
			mainText = customText
		}

		if customSubText != "" {
			subText = customSubText
		}

		main := object.NewText(mainText, object.MessageTypeMain, textColor)
		sub := object.NewText(subText, object.MessageTypeSub, textColor)

		d := drawer.NewTextDrawer(main, sub)
		if err := d.Draw(inputPath, outputPath); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	// Required flags
	rootCmd.Flags().StringVarP(&inputPath, "input", "i", "", "input image path (required)")
	rootCmd.MarkFlagRequired("input")

	// Optional flags
	rootCmd.Flags().StringVarP(&outputPath, "output", "o", "", "output file path (optional, default: current directory with auto-generated filename)")
	rootCmd.Flags().StringVarP(&customText, "text", "t", "", "custom text to embed (optional, default: 'LGTM')")
	rootCmd.Flags().StringVarP(&customSubText, "sub-text", "s", "", "custom sub-text to embed (optional, default: 'Looks Good To Me')")
	rootCmd.Flags().StringVarP(&color, "color", "c", "white", "text color: 'white' or 'black' (optional)")
	rootCmd.Flags().BoolVar(&gopher, "gopher", false, "embed gopher image instead of text (optional)")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
