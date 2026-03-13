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


## CLI Usage

In addition to the MCP server, `script-mcp` ships a standalone CLI binary (`script-cli`) for direct terminal use — no MCP client needed.

### Installation

```bash
just install-cli
# or
go install github.com/nguyenvanduocit/script-mcp/cmd/script-cli@latest
```

### Quick Start

No credentials required. The CLI executes scripts locally.

```bash
script-cli execute --content "echo Hello, World!"
```

### Commands

| Command | Description |
|---------|-------------|
| `execute` (or `exec`, `run`) | Safely execute a script with timeout |

### Examples

```bash
# Run a shell command
script-cli execute --content "ls -la"

# Run with a specific interpreter
script-cli execute --content "print('Hello')" --interpreter /usr/bin/python3

# Run with a working directory
script-cli execute --content "pwd" --working-dir /tmp

# Run a multi-line script
script-cli execute --content "#!/bin/bash
echo 'Starting...'
date
echo 'Done'"

# JSON output (captures stdout/stderr separately)
script-cli exec --content "echo hello" --output json
```

### Flags

| Flag | Description |
|------|-------------|
| `--content` | Script content to execute (required) |
| `--interpreter` | Interpreter path (default: `/bin/sh`) |
| `--working-dir` | Working directory for execution |
| `--env` | Path to `.env` file |
| `--output` | Output format: `text` (default) or `json` |

> **Note:** Scripts time out after 30 seconds.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Installation

### Homebrew (macOS/Linux)

```bash
brew install nguyenvanduocit/tap/script-mcp
```
