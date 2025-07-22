package toolhandlers

import (
	"context"
	"mcp-sample/pkg/models"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func SayHi(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[models.SayHiParams]) (*mcp.CallToolResultFor[any], error) {
	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: "Hi " + params.Arguments.Name,
			},
		},
	}, nil
}
