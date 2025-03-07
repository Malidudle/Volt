# Volt

[![Go Report Card](https://goreportcard.com/badge/github.com/username/volt)](https://goreportcard.com/report/github.com/username/volt)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/github/go-mod/go-version/username/volt)](https://github.com/username/volt)

A Go-based API framework with automatic route discovery and hot reloading.

## Features

- Automatic API route discovery and registration
- Built-in cross-platform hot reloading
- JSON response formatting
- Clean directory structure
- Simple route definition

## Installation

```bash
# Clone the repository
git clone https://github.com/username/volt.git
cd volt

# Install dependencies
go mod download

# Install Air for hot reloading (optional)
make install-air
```

## Directory Structure

API routes follow this organization:

```
app/
├── route.go        # Root API route (/api)
└── example/
    └── route.go    # Example API route (/example)
```

## Route Definition

Each route exports:

```go
// HTTP method (GET, POST, etc.)
var Method = http.MethodGet

// Handler function
func Handler(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
    return responseData, statusCode, err
}
```

## Usage

### Commands

| Command            | Description                   |
| ------------------ | ----------------------------- |
| `make run`         | Run the API server            |
| `make generate`    | Generate route definitions    |
| `make dev`         | Run with hot reload           |
| `make build`       | Build application binary      |
| `make install-air` | Install Air for hot reloading |

Custom port: `PORT=3000 make run`

### Cross-Platform Hot Reloading

The framework supports hot reloading on all platforms:

- **All Platforms**: `make dev`
- **macOS/Linux**: `./dev.sh`
- **Windows**: `dev.bat`

### Troubleshooting Hot Reloading

1. **Install Air**: `make install-air`
2. **Add Go bin to PATH**:
   - Bash/zsh: `export PATH=$PATH:$HOME/go/bin`
   - Windows: `set PATH=%PATH%;%USERPROFILE%\go\bin`
3. **Port conflicts**: Default is 8080; change in `.air.toml` if needed
4. **Infinite rebuilds**: Already prevented by default configuration

## Adding a New Route

1. Create a directory under `app/` (e.g., `app/users/`)
2. Add `route.go` with `Method` and `Handler`
3. Run `make generate`
4. The route will be available at the corresponding path (e.g., `/users`)

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
