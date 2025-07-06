# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go CLI tool called "lgtm" that embeds "LGTM" text on images with customizable colors and fonts. The tool can also embed a gopher image and outputs the result as a JPEG file.

## Development Commands

### Build and Run
```bash
# Build the project
go build -o lgtm

# Run directly
go run main.go [options] <image-path>

# Install globally
go install github.com/tMinamiii/lgtm@latest
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
- **main**: Entry point with CLI flag parsing and orchestration
- **drawer**: Contains the `Drawer` interface and implementations (`TextDrawer`, `GopherDrawer`, `ImageDrawer`)
- **object**: Contains data structures and utilities for text, images, and fonts

### Key Architecture Patterns
- **Interface-based design**: The `Drawer` interface allows different drawing strategies
- **Two drawing modes**: Text-based LGTM embedding vs. gopher image embedding
- **Font abstraction**: Supports multiple font types (NotoSansMono, LINESeedJP) through the `Font` type
- **Color theming**: Supports white and black text colors

### Data Flow
1. CLI flags parsed in `main.go`
2. Based on flags, either `GopherDrawer` or `TextDrawer` is instantiated
3. For text mode: `Text` objects are created with specified font, color, and message type
4. Drawer processes the input image and outputs `<filename>-lgtm.jpeg`

## Dependencies

- **github.com/fogleman/gg**: 2D graphics library for drawing operations
- **github.com/disintegration/imaging**: Image processing and manipulation
- **github.com/stretchr/testify**: Testing framework for assertions
- **github.com/pkg/errors**: Enhanced error handling

## Testing Strategy

The project uses table-driven tests with testify assertions. Test files are located alongside their corresponding source files (e.g., `font_test.go`, `image_test.go`, `text_test.go`).