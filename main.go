package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/server"
	"github.com/nguyenvanduocit/script-mcp/tools"
)

func main() {
	ssePort := flag.String("sse_port", "", "Port for SSE server. If not provided, will use stdio")
	flag.Parse()

	mcpServer := server.NewMCPServer(
		"Script Tool",
		"1.0.0",
		server.WithLogging(),
		server.WithPromptCapabilities(true),
		server.WithResourceCapabilities(true, true),
	)

	// Register Script tool
	tools.RegisterScriptTool(mcpServer)

	if *ssePort != "" {
		sseServer := server.NewSSEServer(mcpServer)
		if err := sseServer.Start(fmt.Sprintf(":%s", *ssePort)); err != nil {
			log.Fatalf("Server error: %v", err)
		}
	} else {
		if err := server.ServeStdio(mcpServer); err != nil {
			panic(fmt.Sprintf("Server error: %v", err))
		}
	}
}
