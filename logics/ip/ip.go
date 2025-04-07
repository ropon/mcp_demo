package ip

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/ropon/requests/v2"
)

const (
	GetCurrentIpInfoToolName = "get_current_ip_info"
	GetIpInfoToolName        = "get_ip_info"
)

var GetCurrentIpInfoTool = mcp.NewTool(
	GetCurrentIpInfoToolName,
	mcp.WithDescription("This is a tool from the demo MCP server.\nGet information about the current ip info"),
)

var GetIpInfoTool = mcp.NewTool(
	GetIpInfoToolName,
	mcp.WithDescription("Get IP information"),
	mcp.WithString(
		"ip",
		mcp.Description("IP address"),
	),
)

func GetCurrentIpInfoHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	res, err := requests.Get("https://ip.rss.ink/json")
	if err != nil {
		return mcp.NewToolResultText("Network error: Unable to connect to Ip API"), err
	}

	return mcp.NewToolResultText(res.Text()), nil
}

func GetIpInfoHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return mcp.NewToolResultText(""), nil
}
