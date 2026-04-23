# Copilot instructions for GO-AWS-course

## Build, test, and lint commands
- Build entire repo: `go build ./...`
- Build a single package: `go build ./path/to/package`
- Run the main app (if running a module): `go run ./learning-go` or `go run ./path/to/cmd` (this repo contains a small learning-go module)
- Run all tests: `go test ./...`
- Run a single test by name: `go test ./path/to/package -run ^TestName$` (use `-run` regex to match the test)
- Run a single package's tests: `go test ./path/to/package`
- Quick vetting/static checks (if installed):
  - `go vet ./...`
  - `staticcheck ./...` (if staticcheck is available)
  - `golangci-lint run` (if repository adds golangci-lint config)

Notes: There are no repository-level Makefile or CI lint scripts detected. If you add CI workflows or a Makefile, include shortcut commands there.

## High-level architecture
- Purpose: this repository is a course companion containing exercises and small projects focused on Go and AWS integrations. It is not a single production microservice; instead it contains multiple small modules and sections for learning.
- Top-level layout (high-level):
  - README.md: course overview and goals
  - learning-go/: a Go module (`module learning-go`) containing example exercises and a small main entrypoint. Several folders like `1section`.. appear to hold exercises.
  - Other directories (e.g., `Mini_Project`) contain standalone exercises/projects.
- How components relate:
  - Each exercise or project is self-contained inside its folder; most use Go modules or the repository-level modules. Build and test can be run per-package with the standard `go` commands.
  - Course material expects Docker and AWS integrations conceptually (mentioned in README) — if implementations are added, they will likely live in per-project Dockerfile and infra subfolders.

## Key conventions for this repo
- Exercises are organized into numbered section folders (e.g., `1section`, `2section`, ...). Treat each folder as an independent exercise package.
- Go module usage:
  - The `learning-go` subfolder is a Go module (go 1.25). Follow module boundaries when running `go build` and `go test` (use package paths like `./learning-go/...`).
- Tests:
  - Tests are expected to be runnable with `go test` per-package. Use `-run` with an anchored regex to run a single test function name.
- Linting/CI:
  - No CI workflows or linter configs were detected. If adding linters prefer `golangci-lint` and put its config at repo root; `golangci-lint run` is the de-facto command.
- Repository-level docs:
  - Keep README as the canonical course overview. If adding contributing rules or code-of-conduct, prefer top-level CONTRIBUTING.md or files under `.github/`.

## Assistant/AI config discovery
- No existing assistant-specific config files were found (CLAUDE.md, .cursorrules, AGENTS.md, etc.). If such files are added, mirror important rules into this file.

---
If this file needs to include CI job snippets, Makefile targets, or project-specific lint commands, say which area should be covered and Copilot will add a focused section.
