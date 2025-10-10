# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2025-01-10

### Added
- Initial release of Divoom PIXOO64 Golang API
- Complete HTTP client for PIXOO64 device communication
- **DisplayText()** helper method for easy text display
- **SendBlankScreen()** helper for blank backgrounds
- **SendColorScreen()** helper for colored backgrounds
- Channel management (SetChannelIndex, GetChannelIndex, SetCustomPageIndex)
- Device management (reboot, time, screen control)
- Display settings (brightness, rotation, mirror mode, white balance)
- Text display with customizable position, font, color, and alignment
- GIF animation support (upload custom, play from URL)
- Tools support (timer, stopwatch, scoreboard, buzzer)
- System settings (timezone, weather, location)
- **Embedded Swagger UI** for interactive API documentation
- OpenAPI 3.0 specification
- Comprehensive examples in `cmd/` directory
- Full documentation (README, API_REFERENCE, SOLUTION)

### Features
- ✅ One-line text display: `client.DisplayText("Hello!", "#00FF00")`
- ✅ Custom text options (position, font, alignment, scrolling)
- ✅ Colored backgrounds with text overlay
- ✅ GIF from URL with text overlay
- ✅ Custom 2-frame animations with text
- ✅ Automatic GIF ID reset and channel management
- ✅ Proper timing delays for reliable operation

### Technical Details
- Text display requires Custom channel (3) and CustomPageIndex 1
- Text overlays on GIF/animation (cannot display standalone)
- 2-frame animations with different frame data required
- ResetGifID() necessary before uploading new GIFs
- Timing delays (500ms-2s) required between operations

### Examples Included
- `cmd/working-example/` - All 4 display modes
- `cmd/final-test/` - Simple one-line test
- `cmd/swagger-server/` - Interactive documentation server
- Multiple test commands for debugging

### Documentation
- Complete README with quick start examples
- API_REFERENCE.md with all methods documented
- SOLUTION.md with detailed technical explanation
- TROUBLESHOOTING.md for common issues
- Inline code documentation
- Interactive Swagger UI

### Dependencies
- Go 1.21 or higher
- No external dependencies (uses standard library only)

## [Unreleased]

### Planned
- Additional animation helpers
- More font options
- Bulk operations support
- Performance optimizations

---

## Version History

- **1.0.0** (2025-01-10) - Initial release with full PIXOO64 API support

[1.0.0]: https://github.com/eugene-bert/divoom-golang-api/releases/tag/v1.0.0
