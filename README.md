# LGoTM

Compatible image types are below extensions.

- jpeg, jpg
- png

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
