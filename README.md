# Script Tool

A tool for executing command line scripts through MCP.

## Features

- Execute command line scripts safely
- Support for different interpreters
- Timeout protection
- Output and error capture
- Cross-platform support (Linux, macOS, Windows)

## Installation

There are several ways to install the Script Tool:

### Option 1: Download from GitHub Releases

1. Visit the [GitHub Releases](https://github.com/nguyenvanduocit/script-mcp/releases) page
2. Download the binary for your platform:
   - script-mcp_linux_amd64` for Linux
   - `script-mcp_darwin_amd64` for macOS
   - `script-mcp_windows_amd64.exe` for Windows
3. Make the binary executable (Linux/macOS):
   ```bash
   chmod +x script-mcp_*
   ```
4. Move it to your PATH (Linux/macOS):
   ```bash
   sudo mv script-mcp_* /usr/local/bin/script-mcp@latest
   ```

### Option 2: Go install

```
go install github.com/nguyenvanduocit/script-mcp
```

## Config

### Claude

```
{
  "mcpServers": {
    "script": {
      "command": "/path-to/script-mcp"
    }
  }
}
```


## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
