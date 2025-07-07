# lgtm
![Coverage](https://img.shields.io/badge/Coverage-70.0%25-brightgreen)

A Go library and CLI tool for embedding custom text on images. By default, it embeds "LGTM" (Looks Good To Me) on images, but you can customize the text, colors, and even embed a gopher image instead.

![lunch-lgtm](https://user-images.githubusercontent.com/31730505/194919314-fc3b9fb9-fd47-46bf-a91a-2d148caf50b3.jpg)

## Installation

### CLI Tool
```sh
go install github.com/tMinamiii/lgtm/cmd/lgtm@latest
```

### Go Library
```sh
go get github.com/tMinamiii/lgtm
```

## Usage

### CLI Tool

```sh
lgtm --help

LGTM is a CLI tool that embeds custom text on images with customizable colors.
It can also embed a gopher image and outputs the result as a JPEG file.
By default, it embeds "LGTM" as main text and "Looks Good To Me" as sub-text.
You can customize both using the --text and --sub-text flags.

Usage:
  lgtm [flags]

Flags:
  -c, --color string      text color: 'white' or 'black' (optional) (default "white")
      --gopher            embed gopher image instead of text (optional)
  -h, --help              help for lgtm
  -i, --input string      input image path (required)
  -o, --output string     output file path (optional, default: current directory with auto-generated filename)
  -s, --sub-text string   custom sub-text to embed (optional, default: 'Looks Good To Me')
  -t, --text string       custom text to embed (optional, default: 'LGTM')
```

#### CLI Examples

Output `image-lgtm.jpeg` in the current directory when run below command.

```sh
# Basic usage (embeds "LGTM" and "Looks Good To Me")
lgtm -i image.jpeg

# With custom main text only
lgtm -i image.jpeg -t "Hello World"

# With custom sub-text only
lgtm -i image.jpeg -s "Custom subtitle"

# With both custom main and sub-text
lgtm -i image.jpeg -t "Hello" -s "World"

# With custom color
lgtm -i image.jpeg -c black -t "Custom Text" -s "Custom Sub"

# With custom output path
lgtm -i image.jpeg -o output.jpg -t "My Text" -s "My Subtitle"

# Gopher mode
lgtm -i image.jpeg --gopher
```

### Go Library

You can also use lgtm as a Go library in your projects:

```go
package main

import (
    "github.com/tMinamiii/lgtm"
)

func main() {
    // Text mode with custom text
    mainText := lgtm.NewMainText("Hello", lgtm.TextColorWhite)
    subText := lgtm.NewSubText("World", lgtm.TextColorWhite)
    drawer := lgtm.NewTextDrawer(mainText, subText, "input.jpg", "output.jpg")
    if err := drawer.Draw(); err != nil {
        panic(err)
    }

    // Gopher mode
    gopherDrawer := lgtm.NewGopherDrawer("input.jpg", "output-gopher.jpg")
    if err := gopherDrawer.Draw(); err != nil {
        panic(err)
    }
}
```

#### Available Types and Functions

- `NewMainText(text string, color TextColor) *Text` - Creates main text with specified color
- `NewSubText(text string, color TextColor) *Text` - Creates sub-text with specified color
- `NewTextDrawer(main, sub *Text, inputPath, outputPath string) Drawer` - Creates text drawer
- `NewGopherDrawer(inputPath, outputPath string) Drawer` - Creates gopher drawer
- `TextColorWhite` and `TextColorBlack` - Available text colors
- `DefaultMainText` and `DefaultSubText` - Default text constants

#### Supported Features

- **Image Formats**: JPEG, PNG, GIF (including animated GIFs)
- **Text Colors**: White and black with automatic contrast
- **Font**: Embedded NotoSansMono-Bold for consistent rendering
- **Auto-sizing**: Automatic font size calculation based on image dimensions
- **Flexible Output**: Custom output paths or auto-generated filenames

## License

### About lgtm cli

The MIT License (MIT)

Copyright (c) 2022 Takahiro Minami

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

### About Gopher

The Go gopher was designed by Renee French. (http://reneefrench.blogspot.com/)
The design is licensed under the Creative Commons 3.0 Attributions license.
Read this article for more details: https://blog.golang.org/gopher


### About Noto font

Usage and redistribution conditions are specified in the license. The most common license is the SIL Open Font License. Some fonts are under the Apache license or Ubuntu Font License. You can redistribute open source fonts according to those conditions.

https://developers.google.com/fonts/faq
