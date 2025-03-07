# Contributing to Volt

Thank you for considering contributing to Volt!

## How to Contribute

1. Fork the repository
2. Create a feature branch: `git checkout -b my-new-feature`
3. Make your changes
4. Run tests: `make test`
5. Commit your changes: `git commit -am 'Add some feature'`
6. Push to the branch: `git push origin my-new-feature`
7. Submit a pull request

## Development Environment

Set up your development environment:

```bash
# Clone your fork
git clone https://github.com/YOUR-USERNAME/volt.git
cd volt

# Install Air for hot reloading
make install-air

# Run the dev server
make dev
```

## Code Style

- Format your code using `make fmt`
- Validate code with `make vet`
- Follow Go standard practices

## Adding a New Route

1. Create a directory under `app/` (e.g., `app/users/`)
2. Add `route.go` with `Method` and `Handler`
3. Run `make generate` to update routes

## Pull Request Process

1. Update the README.md with details of changes if appropriate
2. The PR should work on Go 1.19+
3. PRs need to be approved by a maintainer before merging
