#!/bin/bash
# Development server starter script for Unix-like systems

echo "Running Backend Framework Development Server"

echo "Generating routes..."
go run cmd/generate/main.go

echo "Looking for Air..."
if command -v air >/dev/null 2>&1; then
    echo "Found Air in PATH"
    AIR_PATH="air"
elif [ -f "$HOME/go/bin/air" ]; then
    echo "Found Air in $HOME/go/bin"
    AIR_PATH="$HOME/go/bin/air"
elif [ -n "$GOPATH" ] && [ -f "$GOPATH/bin/air" ]; then
    echo "Found Air in $GOPATH/bin"
    AIR_PATH="$GOPATH/bin/air"
else
    echo "Air not found. Running without hot reload..."
    PORT=8080 go run main.go
    exit 0
fi

if [ ! -f ".air.toml" ]; then
    echo "Creating temporary Air config..."
    cat > ".air.toml.tmp" << EOF
root = "."
tmp_dir = "tmp"

[build]
cmd = "go run cmd/generate/main.go && go build -o ./tmp/main ."
bin = "./tmp/main"
include_ext = ["go"]
exclude_dir = ["tmp", "vendor", ".git"]
exclude_file = ["routes.go", "pkg/router/routes.go"]
exclude_regex = [".*routes\\.go$"]
delay = 2000
stop_on_error = true
send_interrupt = true

[env]
PORT = "8080"

[log]
time = true
EOF
    AIR_CONFIG=".air.toml.tmp"
else
    AIR_CONFIG=".air.toml"
fi

echo "Starting development server with Air..."
PORT=8080 $AIR_PATH -c $AIR_CONFIG 