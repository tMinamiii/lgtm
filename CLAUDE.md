# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go CLI tool called "lgtm" that embeds custom text on images with customizable colors. By default it embeds "LGTM" as main text and "Looks Good To Me" as sub-text, but users can customize both using the `--text` and `--sub-text` flags. The tool can also embed a gopher image and outputs the result as a JPEG file. Uses Cobra for modern CLI interface. When no output path is specified with `-o`, files are saved to the current directory.

## Development Commands

### Build and Run
```bash
# Build the CLI tool from root directory
go build -o lgtm ./cmd/lgtm

# Run directly with cobra
go run ./cmd/lgtm [flags]

# Install globally
go install github.com/tMinamiii/lgtm/cmd/lgtm@latest

# Show help (cobra provides rich help)
./lgtm --help
```

### CLI Usage Examples
```bash
# Basic usage (embeds "LGTM" and "Looks Good To Me")
./lgtm -i image.jpg

# With custom main text only
./lgtm -i image.jpg -t "Hello World"

# With custom sub-text only
./lgtm -i image.jpg -s "Custom subtitle"

# With both custom main and sub-text
./lgtm -i image.jpg -t "Hello" -s "World"

# With output path specification
./lgtm -i image.jpg -o output.jpg -t "Custom Text" -s "Custom Sub"

# With color option
./lgtm -i image.jpg -c black -t "My Message" -s "My Subtitle"

# Gopher mode
./lgtm -i image.jpg --gopher
```

### Testing
```bash
# Run all tests
go test ./...

# Run tests with coverage (matches CI pipeline)
go test -v ./... -covermode=count -coverprofile=coverage.out
go tool cover -func=coverage.out -o=coverage.out

# Run specific test file
go test -v ./image_test.go

# Build and test (CI pipeline)
go build -o lgtm ./cmd/lgtm
```

### Test Data
- **testdata/images/**: Contains comprehensive test images of various sizes and aspect ratios
- **testdata/results/**: Expected output images for regression testing
- Tests cover extreme aspect ratios, tiny/huge images, and various formats

### Code Quality
```bash
# Format code
go fmt ./...

# Tidy dependencies
go mod tidy

# Run staticcheck (if available)
staticcheck ./...
```

## Architecture

The codebase follows a library-with-CLI architecture:

### Core Components
- **cmd/lgtm/main.go**: CLI entry point using Cobra framework that imports the core library
- **Root directory**: Contains the core library implementation (`drawer.go`, `image_drawer.go`, `gopher_drawer.go`, `text.go`, `font.go`)
- **data/**: Static assets including NotoSansMono-Bold.otf font and gopher.png image

### Key Architecture Patterns
- **Library-first design**: Core functionality is implemented as a Go library in the root directory
- **CLI wrapper**: `cmd/lgtm/main.go` provides command-line interface using Cobra
- **Interface-based drawing**: `Drawer` interface with `TextDrawer` and `GopherDrawer` implementations
- **Text abstraction**: `Text` struct with embedded font sizing and positioning logic
- **Two drawing modes**: Text embedding (default "LGTM") vs. gopher image embedding

### Data Flow
1. Cobra CLI framework parses flags including required `-i` input path in `cmd/lgtm/main.go`
2. Based on `--gopher` flag, either `NewGopherDrawer` or `NewTextDrawer` is instantiated
3. For text mode: Text objects are created using `NewMainText` and `NewSubText` with custom text or defaults
4. Drawer processes the input image and outputs to specified path or auto-generated filename in current directory
5. Core processing (text sizing, image manipulation, font rendering) is handled by the library in the root directory

### Important Notes
- This repository contains both the CLI tool and the core library implementation
- The core library handles complex logic including font sizing algorithms, image processing, and text rendering
- Image processing supports multiple formats including GIF animation
- Font rendering uses embedded NotoSansMono-Bold.otf for consistent output across platforms

## Dependencies

### Core Dependencies
- **github.com/spf13/cobra**: Modern CLI framework for command structure and flag parsing
- **github.com/fogleman/gg**: 2D graphics library for drawing operations
- **github.com/disintegration/imaging**: Image processing and manipulation
- **golang.org/x/image**: Extended image format support and font rendering
- **github.com/pkg/errors**: Enhanced error handling
- **github.com/stretchr/testify**: Testing utilities and assertions

### Development Dependencies
- **github.com/tMinamiii/lgtm-core**: External reference (likely older version or tag of this repository)
- **staticcheck**: Static analysis tool (configured in staticcheck.conf)

## CI/CD Pipeline

The GitHub Actions workflow (`/.github/workflows/build.yml`) automatically:
1. Builds the CLI tool with `go build -o lgtm`
2. Runs tests with coverage reporting using `go test -v ./... -covermode=count -coverprofile=coverage.out`
3. Updates coverage badge in README.md using tj-actions/coverage-badge-go
4. Commits and pushes changes back to the repository (including test result images)

The pipeline runs on every push to the main branch and automatically updates testdata/results/ with new test outputs.

## ルール

- プロジェクトルートに画像を生成してはいけない
- ビルド成果物は削除する
- YAGNI（You Aren't Gonna Need It）：今必要じゃない機能は作らない
- DRY（Don't Repeat Yourself）：同じコードを繰り返さない
- KISS（Keep It Simple Stupid）：シンプルに保つ
