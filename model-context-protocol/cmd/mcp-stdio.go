package main

import (
	"context"
	"log"
	toolhandlers "mcp-sample/pkg/tool-handlers"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func main() {
	ctx := context.Background()

	// 1. Instantiate the server implementation.
	// fmt.Println("Creating new mcp server")
	server := mcp.NewServer(&mcp.Implementation{Name: "greeter", Version: "v0.0.1"}, nil)

	// 2. Register the greet tool.
	// fmt.Println("Adding tool to mcp server")
	mcp.AddTool(server, &mcp.Tool{
		Name:        "greet",
		Description: "say hi",
	}, toolhandlers.SayHi)

	// 3. Run the server on STDIO transport (blocks until client closes pipes).
	// fmt.Println("Run the mcp server over the stdio transport")
	if err := server.Run(ctx, mcp.NewStdioTransport()); err != nil {
		log.Fatal(err)
	}

}
