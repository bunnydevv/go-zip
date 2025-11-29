# go-zip

A powerful CLI tool for file compression and decompression, similar to 7zip, built with Go.

## Features

- **Multiple Compression Formats**: ZIP, TAR, TAR.GZ, TAR.BZ2, GZIP
- **Compression Levels**: Adjustable compression levels (0-9)
- **Password Protection**: Encrypt ZIP archives with passwords
- **Directory Support**: Compress entire directories recursively
- **List Contents**: View archive contents without extracting
- **Fast & Efficient**: Built with Go for high performance
- **Cross-Platform**: Works on Windows, macOS, and Linux

## Installation

### From Source

```bash
git clone https://github.com/bunnydevv/go-zip.git
dcd go-zip
go build -o go-zip
```

### Using Go Install

```bash
go install github.com/bunnydevv/go-zip@latest
```

## Usage

### Compress Files

```bash
# Compress files to ZIP
go-zip compress file1.txt file2.txt -o archive.zip

# Compress directory to TAR.GZ
go-zip compress mydir/ -t tar.gz -o backup.tar.gz

# Compress with specific compression level

go-zip compress file.txt -l 9 -o file.zip

# Compress with password protection

go-zip compress secret.txt -p mypassword -o secret.zip
```

### Decompress/Extract Files

```bash
# Extract archive to current directory

go-zip decompress archive.zip

# Extract to specific directory

go-zip decompress archive.tar.gz -o ./extracted

# Extract password-protected archive

go-zip decompress secret.zip -p mypassword
```

### List Archive Contents

```bash
# List files in archive

go-zip list archive.zip

# List files in TAR archive

go-zip list backup.tar.gz
```

### Version

```bash
go-zip version
```

## Supported Formats

| Format | Extension | Compress | Decompress | List |
|--------|-----------|----------|------------|------|
| ZIP | `.zip` | ✅ | ✅ | ✅ |
| TAR | `.tar` | ✅ | ✅ | ✅ |
| TAR.GZ | `.tar.gz`, `.tgz` | ✅ | ✅ | ✅ |
| TAR.BZ2 | `.tar.bz2`, `.tbz2` | 🚧 | ✅ | ✅ |
| GZIP | `.gz` | ✅ | ✅ | ❌ |

## Command Reference

### compress (c, add)

Compress files or directories into an archive.

**Flags:**
- `-t, --type`: Compression type (zip, tar, tar.gz, tar.bz2, gzip) [default: zip]
- `-o, --output`: Output archive file name
- `-l, --level`: Compression level 0-9 [default: 6]
- `-p, --password`: Password for encryption (ZIP only)

### decompress (d, extract, x)

Extract files from a compressed archive.

**Flags:**
- `-o, --output`: Output directory for extracted files [default: .]
- `-p, --password`: Password for encrypted archives

### list (l, ls)

Display the contents of an archive.

## Examples

### Backup a directory

```bash
go-zip compress ~/Documents/project -t tar.gz -l 9 -o project-backup.tar.gz
```

### Create encrypted archive

```bash
go-zip compress sensitive-data/ -p StrongPassword123 -o data.zip
```

### Extract and view

```bash

go-zip list archive.zip
go-zip decompress archive.zip -o ./output
```

## Development

### Project Structure

```
go-zip/
├── cmd/
│   ├── root.go          # Root command
│   ├── compress.go      # Compress command
│   ├── decompress.go    # Decompress command
│   ├── list.go          # List command
│   └── version.go       # Version command
├── pkg/
│   └── compression/
│       ├── zip.go       # ZIP operations
│       ├── tar.go       # TAR operations
│       └── gzip.go      # GZIP operations
├── main.go
├── go.mod
└── README.md
```

### Building

```bash
go build -o go-zip
```

### Running Tests

```bash
go test ./...
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License

## Roadmap

- [ ] Add bzip2 compression support
- [ ] Add 7z format support
- [ ] Progress bars for large files
- [ ] Parallel compression
- [ ] File splitting for large archives
- [ ] Archive integrity checking
- [ ] Self-extracting archives

## Author

Created by [@bunnydevv](https://github.com/bunnydevv)
