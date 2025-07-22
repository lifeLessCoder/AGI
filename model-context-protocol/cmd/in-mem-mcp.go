package main

import (
	"context"
	"fmt"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type SayHiParams struct {
	Name string `json:"name"`
}

func SayHi(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[SayHiParams]) (*mcp.CallToolResultFor[any], error) {
	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: "Hi " + params.Arguments.Name,
			},
		},
	}, nil
}

func main() {
	ctx := context.Background()
	fmt.Println("Creating client and server in memory transports")
	clientTransport, serverTransport := mcp.NewInMemoryTransports()
	fmt.Println("Creating new mcp server")
	server := mcp.NewServer(&mcp.Implementation{Name: "greeter", Version: "v0.0.1"}, nil)
	fmt.Println("Adding tool to mcp server")
	mcp.AddTool(server, &mcp.Tool{
		Name:        "greet",
		Description: "say hi",
	}, SayHi)

	fmt.Println("Connecting mcp server over the in memory server transport")
	serverSession, err := server.Connect(ctx, serverTransport)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creating new mcp client")
	client := mcp.NewClient(&mcp.Implementation{Name: "client"}, nil)
	fmt.Println("Connecting the mcp client to the mcp server over the in memory client transport")
	clientSession, err := client.Connect(ctx, clientTransport)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Client calling the tool defined on the server")
	res, err := clientSession.CallTool(ctx, &mcp.CallToolParams{
		Name:      "greet",
		Arguments: map[string]any{"name": "Ankush"},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Printing response from mcp server")
	fmt.Println(res.Content[0].(*mcp.TextContent).Text)

	fmt.Println("Close the client session")
	clientSession.Close()
	fmt.Println("Server waiting for the client to close the session")
	serverSession.Wait()
}
