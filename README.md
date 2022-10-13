# LGoTM

LGoTM embeds `LGTM` string in image.

![lunch-lgtm](https://user-images.githubusercontent.com/31730505/194919314-fc3b9fb9-fd47-46bf-a91a-2d148caf50b3.jpg)

## Installation

```sh
go install github.com/tMinamiii/lgotm@latest
```

## Usage

Enable to specify image path and color which is `white` or `black`.

```sh
lgotm -h

Usage of lgotm:
  -c string
        color 'white' or 'black' (default "white")
  -gopher
        embed gopher
  -i string
        image path
  -main string
        main text (default "L G T M")
  -serif
        serif font
  -sub string
        sub text (default "L o o k s   G o o d   T o   M e")
```

### Example

Output `image-lgtm.jpeg` when run below command.

```sh
lgotm -i image.jpeg -c white
```

## License

### About LGoTM

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
