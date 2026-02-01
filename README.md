# pkbin

**pkbin** is a lightweight, `npm run`, but without `npm`

## Features

- 🚀 **Fast**: Written in Go, starts instantly
- 💬 **JSONC Support**: Use comments in your config file for better documentation
- 🔄 **Cross-platform**: Works on Windows, macOS, and Linux
- 📦 **Zero Runtime Dependencies**: Single binary, no external dependencies required
- 🎯 **Simple**: Just define scripts and run them

## Installation

```bash
go install github.com/mofax/pkbin@latest
```

Or download the latest release from the [releases page](https://github.com/mofax/pkbin/releases).

## Quick Start

1. Create a `pkbin.jsonc` file in your project root:

```jsonc
{
  // Your project scripts
  "scripts": {
    "build": "make build",
    "test": "pytest",
    "dev": "python -m http.server 8000",
    "lint": "ruff check ."
  }
}
```

2. Run your scripts:

```bash
pk build
pk test
pk dev
```

## Configuration

The `pkbin.jsonc` file uses JSONC format, which allows comments for better documentation:

```jsonc
{
  // Build scripts
  "scripts": {
    // Build the project
    "build": "make build",
    
    // Run tests with coverage
    "test": "pytest --cov=. --cov-report=html",
    
    // Development server
    "dev": "uvicorn main:app --reload",
    
    // Lint the codebase
    "lint": "ruff check . && mypy .",
    
    // Run multiple commands
    "ci": "pk lint && pk test && pk build"
  }
}
```

### Script Format

Each script is a shell command that will be executed in your project's working directory. The command runs in:
- `/bin/sh` on Unix-like systems (macOS, Linux)
- `cmd` on Windows

All environment variables from your current shell are automatically passed through to the script.

## Usage

```bash
# Run a script
pk <script-name>

If a script is not found, pkbin will exit with an error message.

## Examples

### Python Project

```jsonc
{
  "scripts": {
    "install": "pip install -r requirements.txt",
    "test": "pytest -v",
    "test:coverage": "pytest --cov=. --cov-report=html",
    "lint": "ruff check . && mypy .",
    "format": "black . && isort .",
    "dev": "uvicorn main:app --reload"
  }
}
```

### Node.js/TypeScript Project

```jsonc
{
  "scripts": {
    "install": "npm install",
    "build": "tsc && npm run bundle",
    "test": "jest",
    "test:watch": "jest --watch",
    "lint": "eslint . --ext .ts,.tsx",
    "dev": "tsx watch src/index.ts"
  }
}
```

### Go Project

```jsonc
{
  "scripts": {
    "build": "go build -o bin/app ./cmd/app",
    "test": "go test -v -race ./...",
    "test:coverage": "go test -coverprofile=coverage.out ./...",
    "lint": "golangci-lint run",
    "fmt": "gofmt -s -w .",
    "clean": "rm -rf bin/ coverage.out"
  }
}
```

### Rust Project

```jsonc
{
  "scripts": {
    "build": "cargo build --release",
    "test": "cargo test",
    "test:verbose": "cargo test -- --nocapture",
    "lint": "cargo clippy -- -D warnings",
    "fmt": "cargo fmt",
    "check": "cargo check"
  }
}
```

### Multi-language Project

```jsonc
{
  "scripts": {
    // Python backend
    "backend:install": "cd backend && pip install -r requirements.txt",
    "backend:test": "cd backend && pytest",
    
    // React frontend
    "frontend:install": "cd frontend && npm install",
    "frontend:build": "cd frontend && npm run build",
    
    // Combined workflows
    "install": "pk backend:install && pk frontend:install",
    "build": "pk backend:test && pk frontend:build",
    "test": "pk backend:test"
  }
}
```

## Development

```bash
# Clone the repository
git clone https://github.com/mofax/pkbin.git
cd pkbin

# Run tests
go test ./...

# Build
go build -o pkbin

# Install locally
go install
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT
