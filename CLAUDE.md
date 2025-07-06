# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go CLI tool called "lgtm" that embeds custom text on images with customizable colors. By default it embeds "LGTM" as main text and "Looks Good To Me" as sub-text, but users can customize both using the `--text` and `--sub-text` flags. The tool can also embed a gopher image and outputs the result as a JPEG file. Uses Cobra for modern CLI interface. When no output path is specified with `-o`, files are saved to the current directory.

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

# Run tests with coverage
go test ./... -coverprofile=coverage.out
go tool cover -func=coverage.out -o=coverage.out

# Build and test (CI pipeline)
go build -o lgtm
```

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

The codebase follows a simple wrapper architecture:

### Core Components
- **main.go**: CLI entry point using Cobra framework, acts as a thin wrapper around the core library
- **github.com/tMinamiii/lgtm-core**: External core library containing all image processing and text rendering functionality

### Key Architecture Patterns
- **Library-based design**: Core functionality is abstracted into a separate module (`lgtm-core`)
- **CLI wrapper**: Main application focuses solely on command-line interface and flag parsing
- **Two drawing modes**: Text embedding (default "LGTM") vs. gopher image embedding
- **Simplified API**: Uses factory functions from core library (`NewGopherDrawer`, `NewTextDrawer`, `NewMainText`, `NewSubText`)

### Data Flow
1. Cobra CLI framework parses flags including required `-i` input path in `main.go`
2. Based on `--gopher` flag, either `lgtm.NewGopherDrawer` or `lgtm.NewTextDrawer` is instantiated
3. For text mode: Text objects are created using `lgtm.NewMainText` and `lgtm.NewSubText` with custom text or defaults
4. Drawer processes the input image and outputs to specified path or auto-generated filename in current directory
5. All core processing (text sizing, image manipulation, font rendering) is handled by the external core library

### Important Notes
- The core library handles all complex logic including font sizing algorithms, image processing, and text rendering
- This repository only contains the CLI interface - all core functionality is in `github.com/tMinamiii/lgtm-core`
- When working on core functionality, you'll need to work with the separate `lgtm-core` repository

## Dependencies

### Core Dependencies
- **github.com/spf13/cobra**: Modern CLI framework for command structure and flag parsing
- **github.com/tMinamiii/lgtm-core**: External core library containing all LGTM functionality

### Transitive Dependencies (from core library)
- **github.com/fogleman/gg**: 2D graphics library for drawing operations
- **github.com/disintegration/imaging**: Image processing and manipulation
- **golang.org/x/image**: Extended image format support and font rendering
- **github.com/pkg/errors**: Enhanced error handling

## CI/CD Pipeline

The GitHub Actions workflow (`/.github/workflows/build.yml`) automatically:
1. Builds the CLI tool
2. Runs tests with coverage reporting
3. Updates coverage badge in README.md
4. Commits and pushes changes back to the repository

The pipeline runs on every push to the main branch.