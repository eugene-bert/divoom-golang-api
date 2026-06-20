# Changelog

## v1.0.0

### Added
- **Device Discovery** — `DiscoverDevices()` finds PIXOO64 devices on your LAN via Divoom cloud API
- **Image Display** — `DisplayImageFile()` and `DisplayImageURL()` for one-liner image display
- **Canvas Resize** — `DrawImageResized()` and `DrawImageFill()` with stdlib nearest-neighbor scaling
- **Context Support** — all multi-step methods have `*Context()` variants for cancellation
- **Unit Tests** — 17 tests covering canvas drawing, concurrent access, client construction
- **CI** — GitHub Actions with `go vet` and `-race` detection on Go 1.25/1.26
- **Examples** — 6 runnable examples: basic, image, gif, scrolltext, clock, cpu

### Fixed
- Thread safety: all Canvas drawing methods now properly lock the mutex
- GIF delay array bounds check prevents panic on malformed GIFs
- HTTP status code validation on all device responses
- Centralized ErrorCode checking in `sendCommandWithResponse`

### Removed
- Auto-generated documentation noise (FIXES.md, REAL_FIX.md, etc.)
- Swagger UI package (belongs in server, not library)
