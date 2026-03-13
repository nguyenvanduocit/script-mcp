package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "execute", "exec", "run":
		runExecute(os.Args[2:])
	case "help", "--help", "-h":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `Usage: script-cli <command> [flags]

Commands:
  execute, exec, run  Execute a script

Flags for execute:
  --content       Script content to execute (required)
  --interpreter   Interpreter path (default: /bin/sh)
  --working-dir   Working directory
  --env           Path to .env file
  --output        Output format: text|json (default: text)
`)
}

type result struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
	Error  string `json:"error,omitempty"`
}

func runExecute(args []string) {
	fs := flag.NewFlagSet("execute", flag.ExitOnError)
	env := fs.String("env", "", "Path to .env file")
	content := fs.String("content", "", "Script content (required)")
	interpreter := fs.String("interpreter", "/bin/sh", "Interpreter path")
	workingDir := fs.String("working-dir", "", "Working directory")
	output := fs.String("output", "text", "Output format: text|json")
	fs.Parse(args)

	if *env != "" {
		if err := godotenv.Load(*env); err != nil {
			fmt.Fprintf(os.Stderr, "Error loading .env file: %v\n", err)
			os.Exit(1)
		}
	}

	if *content == "" {
		fmt.Fprintln(os.Stderr, "Error: --content is required")
		fs.Usage()
		os.Exit(1)
	}

	// Create temporary script file
	tmpFile, err := os.CreateTemp("", "script-*.sh")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to create temporary file: %v\n", err)
		os.Exit(1)
	}
	defer os.Remove(tmpFile.Name())

	// Write content to temporary file
	if _, err := tmpFile.WriteString(*content); err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to write to temporary file: %v\n", err)
		os.Exit(1)
	}
	if err := tmpFile.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to close temporary file: %v\n", err)
		os.Exit(1)
	}

	// Make the script executable
	if err := os.Chmod(tmpFile.Name(), 0700); err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to make script executable: %v\n", err)
		os.Exit(1)
	}

	// Create command with 30-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, *interpreter, tmpFile.Name())

	if *workingDir != "" {
		cmd.Dir = *workingDir
	}

	cmd.Env = os.Environ()

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	execErr := cmd.Run()

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Fprintln(os.Stderr, "Error: script execution timed out after 30 seconds")
		os.Exit(1)
	}

	switch *output {
	case "json":
		r := result{
			Stdout: stdout.String(),
			Stderr: stderr.String(),
		}
		if execErr != nil {
			r.Error = execErr.Error()
		}
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		enc.Encode(r)
		if execErr != nil {
			os.Exit(1)
		}
	default: // text
		if stdout.Len() > 0 {
			fmt.Print("Output:\n")
			fmt.Print(stdout.String())
			fmt.Print("\n")
		}
		if stderr.Len() > 0 {
			fmt.Print("Errors:\n")
			fmt.Print(stderr.String())
			fmt.Print("\n")
		}
		if execErr != nil {
			fmt.Fprintf(os.Stderr, "\nExecution error: %v\n", execErr)
			os.Exit(1)
		}
	}
}
