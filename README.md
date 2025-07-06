# lgtm
![Coverage](https://img.shields.io/badge/Coverage-7.8%25-red)

lgtm embeds `LGTM` string on an image.

![lunch-lgtm](https://user-images.githubusercontent.com/31730505/194919314-fc3b9fb9-fd47-46bf-a91a-2d148caf50b3.jpg)

## Installation

```sh
go install github.com/tMinamiii/lgtm@latest
```

## Usage

```sh
lgtm --help

LGTM is a CLI tool that embeds "LGTM" text on images with customizable colors and fonts.
It can also embed a gopher image and outputs the result as a JPEG file.

Usage:
  lgtm [flags]

Flags:
  -c, --color string    text color: 'white' or 'black' (default "white")
  -f, --font string     font type: 'sans' or 'line' (default "sans")
      --gopher          embed gopher image instead of text
  -h, --help            help for lgtm
  -i, --input string    input image path
  -o, --output string   output file path
```

### Examples

Output `image-lgtm.jpeg` in the current directory when run below command.

```sh
# Basic usage
lgtm -i image.jpeg

# With custom color and font
lgtm -i image.jpeg -c black -f line

# With custom output path
lgtm -i image.jpeg -o output.jpg

# Gopher mode
lgtm -i image.jpeg --gopher
```

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


### About LINE Seed Licensing

All content of LINE Seed is copyrighted material owned by LINE Corp.
All typefaces are released under free, open source license.
You can use them for any personal or commercial purpose.
However, the software font files themselves cannot be sold by the other parties other than LINE Corp.
For commercial use, we highly recommend to include attribution in product or service.

### About Noto font

Usage and redistribution conditions are specified in the license. The most common license is the SIL Open Font License. Some fonts are under the Apache license or Ubuntu Font License. You can redistribute open source fonts according to those conditions.

https://developers.google.com/fonts/faq
