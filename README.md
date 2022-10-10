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
  -i string
        image path
```

### Example

Output `image-lgtm.jpeg` when run below command.

```sh
lgotm -i image.jpeg -c white
```
