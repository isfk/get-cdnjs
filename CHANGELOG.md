# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [2.0.0] - 2026-02-10

### ⚠️ Breaking Changes

- **Migration from cdnjs to jsDelivr**
  - Changed API endpoint from `api.cdnjs.com` to `data.jsdelivr.com`
  - Changed CDN URL from `cdnjs.cloudflare.com` to `cdn.jsdelivr.net`
  - Updated package source from cdnjs libraries to npm packages
  - File path format changed: `cdnjs/{lib}/{ver}/{file}` → `npm/{lib}@{ver}/{file}`

- **Configuration Update**
  - Default `file_path` changed from `cdnjs` to `npm`
  - All API calls now use `@` separator for versions (e.g., `jquery@3.7.1`)

### ✨ Added

#### Core Features
- **Proxy Support** - All network requests now respect configured proxy settings
  - API queries go through proxy
  - File downloads go through proxy
  - Solves jsDelivr network issues in restricted regions

- **Delete Command** - New `delete` command for removing directories from qiniu
  ```bash
  ./get-cdnjs delete [directory]
  ```
  - Interactive directory browser with subdirectory display
  - Detailed file preview before deletion
  - Safe deletion with confirmation requirement

- **Retry Mechanism** - Automatic retry for failed downloads
  - 3 retry attempts per file
  - Configurable proxy support for retries
  - Detailed retry logging

- **Batch Processing** - Separated download and upload phases
  - Phase 1: Download all files (with retry)
  - Phase 2: Upload all successful downloads
  - Better error visibility and tracking

#### User Experience
- **Enhanced Error Tracking** - Separate tracking for download and upload failures
  - Detailed error messages with file URLs
  - Statistics summary (success/failure counts)
  - Color-coded terminal output

- **Interactive Directory Browser** - Show existing directories before deletion
  - Tree-structure display with subdirectories
  - Visual representation of qiniu storage
  - Helps identify versions to clean up

- **Progress Feedback** - Real-time progress indicators
  - Download phase with file-by-file status
  - Upload phase with success confirmation
  - Byte count for downloaded files
  - Colored status messages

### 🔄 Changed

- **Download Strategy** - Switched from qiniu Fetch to local download + upload
  - Previous: qiniu server directly fetched from CDN (no proxy support)
  - Current: Download locally (with proxy), then upload to qiniu
  - Benefit: Full proxy support for all network operations

- **File Handling** - In-memory storage instead of temporary files
  - Files stored in memory during download phase
  - Direct upload from memory
  - No disk I/O for better performance

- **API Response Structures** - Updated for jsDelivr API compatibility
  - `JSDelivrVersionsRet` - New structure for version queries
  - `JSDelivrFilesRet` - New nested structure for file listings
  - `FileEntry` - Recursive file entry structure
  - `flattenFiles()` - Helper to extract files from nested structure

### 🐛 Fixed

- Network timeout issues with proxy configuration
- Partial file uploads in error scenarios
- Missing visibility into failed operations

### 📚 Documentation

- Updated README with jsDelivr examples
- New command usage documentation
- Enhanced configuration examples
- Updated operation samples with new URLs

### 🛠️ Internal

- Added `internal/qiniu/delete.go` for delete command logic
- Added `cmd/delete.go` for CLI command
- Enhanced `pkg/qiniu.go` with batch delete operations
- Improved error handling throughout

---

## [1.0.0] - 2024-08-09

### ✨ Added

- Initial release
- cdnjs API integration
- Basic download and upload functionality
- Qiniu cloud storage support
- List command for viewing uploaded resources
- Configuration file support

### 📋 Features

- Download libraries from cdnjs
- Automatic upload to qiniu cloud
- Proxy configuration support
- Version selection interface
- File listing in tree structure

[2.0.0]: https://github.com/isfk/get-cdnjs/compare/v1.0.0...v2.0.0
[1.0.0]: https://github.com/isfk/get-cdnjs/releases/tag/v1.0.0
