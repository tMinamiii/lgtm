![Code Coverage](https://img.shields.io/badge/Code%20Coverage-33%25-critical?style=flat)

# lgtm

https://img.shields.io/badge/Code%20Coverage-83%25-success?style=flat

lgtm embeds `LGTM` string on an image.

![lunch-lgtm](https://user-images.githubusercontent.com/31730505/194919314-fc3b9fb9-fd47-46bf-a91a-2d148caf50b3.jpg)

## Installation

```sh
go install github.com/tMinamiii/lgtm@latest
```

## Usage

```sh
lgtm -h

Usage of lgtm:
  -c string
        color 'white' or 'black' (default "white")
  -f string
        sans, serif, line (default "sans")
  -gopher
        embed gopher
  -i string
        image file path
```

### Example

Output `image-lgtm.jpeg` when run below command.

```sh
lgtm -i image.jpeg -c white
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
