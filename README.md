# LGoTM

LGoTM embeds `LGTM` string in image.

![lunch-lgtm](https://user-images.githubusercontent.com/31730505/194919185-341db7fd-5265-4daf-a842-8706b64604a5.jpg)

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
  -i string
        image path
```

### Example

Output `image-lgtm.jpeg` when run below command.

```sh
lgotm -i image.jpeg -c white
```
