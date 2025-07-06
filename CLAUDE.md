# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go CLI tool called "lgtm" that embeds "LGTM" text on images with customizable colors and fonts. The tool can also embed a gopher image and outputs the result as a JPEG file. Uses Cobra for modern CLI interface. When no output path is specified with `-o`, files are saved to the current directory.

## Development Commands

### Build and Run
```bash
# Build the project
go build -o lgtm

# Run directly with cobra
go run main.go [flags]

# Install globally
go install github.com/tMinamiii/lgtm@latest

# Show help (cobra provides rich help)
./lgtm --help
```

### CLI Usage Examples
```bash
# Basic usage
./lgtm -i image.jpg

# With output path specification
./lgtm -i image.jpg -o output.jpg

# With color and font options
./lgtm -i image.jpg -c black -f line -o custom.jpg

# Gopher mode
./lgtm -i image.jpg --gopher
```

### Testing
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test ./... -coverprofile=coverage.txt

# Run a specific test file
go test ./object -v

# Run a specific test function
go test ./object -run TestSpecificFunction
```

### Code Quality
```bash
# Run staticcheck (all checks enabled)
staticcheck ./...

# Format code
go fmt ./...

# Tidy dependencies
go mod tidy
```

## Architecture

The codebase follows a clean architecture pattern with three main packages:

### Core Packages
- **main**: Entry point using Cobra CLI framework for command handling and orchestration
- **drawer**: Contains the `Drawer` interface and implementations (`TextDrawer`, `GopherDrawer`)
- **object**: Contains data structures and utilities for text, images, and fonts with embedded font files

### Key Architecture Patterns
- **Interface-based design**: The `Drawer` interface with `Draw(inputPath, outputPath string) error` method allows different drawing strategies
- **Two drawing modes**: Text-based LGTM embedding vs. gopher image embedding
- **Font abstraction**: Supports multiple font types (NotoSansMono, LINESeedJP) through embedded `Font` type using Go's `//go:embed` directive
- **Color theming**: Supports white and black text colors via `TextColor` type
- **Adaptive text sizing**: Dynamic font sizing based on image dimensions, text length, and message type (main/sub)
- **Output path control**: Configurable output destination with fallback to auto-generated filenames in current directory

### Data Flow
1. Cobra CLI framework parses flags including required `-i` input path in `main.go`
2. Based on `--gopher` flag, either `GopherDrawer` or `TextDrawer` is instantiated
3. For text mode: `Text` objects are created with specified font, color, and message type
4. Drawer processes the input image and outputs to specified path or auto-generated filename in current directory
5. Text sizing algorithm considers image aspect ratio, safe area calculations, and text length for optimal placement

### Text Sizing Logic
The `FontSize` method in `object/text.go` implements sophisticated sizing:
- **Safe area**: Uses 90% of image width and 70% of image height
- **Aspect ratio awareness**: Adjusts positioning for wide, tall, and normal images
- **Message type scaling**: Main text uses 60% of safe height, sub-text uses 36%
- **Short text optimization**: Applies 1.5x multiplier for text ≤4 characters
- **Size constraints**: Min 12px, max 400px with dynamic calculations

## Dependencies

### Core Dependencies
- **github.com/spf13/cobra**: Modern CLI framework for command structure and flag parsing
- **github.com/fogleman/gg**: 2D graphics library for drawing operations
- **github.com/disintegration/imaging**: Image processing and manipulation
- **golang.org/x/image**: Extended image format support and font rendering
- **github.com/pkg/errors**: Enhanced error handling

### Development Dependencies
- **github.com/stretchr/testify**: Testing framework for assertions

### Font System
The project embeds font files directly in the binary using `//go:embed`:
- **NotoSansMono-Bold.otf**: Default sans-serif font for English text
- **LINESeedJP_OTF_Bd.otf**: Japanese-optimized font for better Asian character support
- **gopher.png**: Embedded gopher image for `--gopher` mode

## Testing Strategy

The project uses table-driven tests with testify assertions. Test files are located alongside their corresponding source files (e.g., `font_test.go`, `image_test.go`, `text_test.go`). Tests focus on:
- Font loading and face creation
- Text sizing calculations and edge cases
- Image processing and format support
- Color and positioning logic