# Release Checklist for Open Source Publication

This document outlines the steps to publish this project as an open source Go package.

## Pre-Release Checklist

### 1. Update Module Path
- [ ] Replace `github.com/eugene-bert/divoom-golang-api` with your actual GitHub username in:
  - `go.mod`
  - `README.md`
  - `examples/main.go`
  - `cmd/swagger-server/main.go`
  - `swagger/README.md`
  - `GETTING_STARTED.md`
  - `API_REFERENCE.md`

### 2. Update License
- [ ] Add your name and year to `LICENSE` file

### 3. Review Documentation
- [ ] Verify all links in README.md work
- [ ] Ensure examples are accurate
- [ ] Check that API_REFERENCE.md is complete

### 4. Code Quality
- [ ] Run `go fmt ./...`
- [ ] Run `go vet ./...`
- [ ] Test all examples with real device
- [ ] Verify Swagger UI works: `go run cmd/swagger-server/main.go`

### 5. Repository Setup
- [ ] Initialize git repository
- [ ] Add all files: `git add .`
- [ ] Create initial commit: `git commit -m "Initial commit"`
- [ ] Create GitHub repository
- [ ] Push code: `git push -u origin main`

## Publishing Steps

### 1. Create GitHub Repository

```bash
# Initialize git (if not already done)
git init

# Add remote
git remote add origin https://github.com/YOUR_USERNAME/divoom-golang-api.git

# Push to GitHub
git branch -M main
git add .
git commit -m "Initial commit: Divoom PIXOO64 Go API client"
git push -u origin main
```

### 2. Tag a Release

```bash
# Create and push a tag for v1.0.0
git tag v1.0.0
git push origin v1.0.0
```

### 3. Create GitHub Release
- Go to your repository on GitHub
- Click "Releases" → "Create a new release"
- Select tag `v1.0.0`
- Title: "v1.0.0 - Initial Release"
- Description: Highlight key features

### 4. Submit to pkg.go.dev
The package will automatically appear on pkg.go.dev once:
- You push a valid Go module with version tag
- Someone runs `go get github.com/YOUR_USERNAME/divoom-golang-api`

### 5. Optional: Announce
- [ ] Share on Reddit (r/golang)
- [ ] Post on Twitter/X with #golang
- [ ] Share in Go community forums

## Post-Release

### Enable GitHub Features
- [ ] Enable GitHub Issues
- [ ] Enable GitHub Discussions (optional)
- [ ] Add repository topics: `golang`, `pixoo64`, `divoom`, `api-client`, `led-display`
- [ ] Add repository description
- [ ] Set repository homepage to documentation site (if any)

### Documentation
- [ ] Star your own repository (for visibility)
- [ ] Add shields/badges to README (build status, Go version, etc.)
- [ ] Consider creating a GitHub Pages site for documentation

## Sample Commands for Quick Setup

```bash
# Update module path
find . -type f -name "*.go" -o -name "*.md" | xargs sed -i '' 's|github.com/eugene-bert|github.com/YOUR_USERNAME|g'

# Format code
go fmt ./...

# Verify build
go build -v ./...

# Git setup
git init
git add .
git commit -m "Initial commit: Divoom PIXOO64 Go API client"
git branch -M main
git remote add origin https://github.com/YOUR_USERNAME/divoom-golang-api.git
git push -u origin main
git tag v1.0.0
git push origin v1.0.0
```

## Module Path Updates Checklist

Files that need `github.com/eugene-bert` replaced:
- [ ] `go.mod` (line 1)
- [ ] `README.md` (multiple locations)
- [ ] `examples/main.go` (import statement)
- [ ] `cmd/swagger-server/main.go` (import statement)
- [ ] `swagger/README.md` (code examples)
- [ ] `GETTING_STARTED.md` (code examples)

## Recommended GitHub Repository Settings

**Repository name:** `divoom-golang-api` or `go-pixoo64`
**Description:** "Go client library for Divoom PIXOO64 LED display device"
**Topics:** `golang`, `go`, `pixoo64`, `divoom`, `api-client`, `led-display`, `swagger`
**Homepage:** Your documentation URL or leave blank
**License:** MIT

## Ready to Publish?

Once you've completed all items above, your package is ready for open source publication!

Users will be able to install it with:
```bash
go get github.com/YOUR_USERNAME/divoom-golang-api
```
