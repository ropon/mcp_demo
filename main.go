package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mark3labs/mcp-go/server"
	"github.com/ropon/mcp_demo/logics/ip"
)

const Version = "1.0.0"

func main() {
	testToken := flag.String("token", "", "Test access token")
	showVersion := flag.Bool("version", false, "Show version information")
	transport := flag.String("transport", "stdio", "Transport type (stdio or sse)")
	addr := flag.String("sse-address", "localhost:8000", "The host and port to start the sse server on")
	flag.Parse()

	if *showVersion {
		fmt.Printf("Demo MCP Server\n")
		fmt.Printf("Version: %s\n", Version)
		os.Exit(0)
	}

	if *testToken == "" {
		fmt.Println("Error: token is required")
		os.Exit(1)
	}

	fmt.Printf("Test token: %s\n", *testToken)

	if err := run(*transport, *addr); err != nil {
		panic(err)
	}
}

func run(transport, addr string) error {
	s := newMCPServer()
	addTools(s)

	switch transport {
	case "stdio":
		if err := server.ServeStdio(s); err != nil {
			if errors.Is(err, context.Canceled) {
				return nil
			}
			return err
		}
	case "sse":
		srv := server.NewSSEServer(s, server.WithBaseURL("http://"+addr))
		log.Printf("SSE server listening on %s", addr)
		if err := srv.Start(addr); err != nil {
			if errors.Is(err, context.Canceled) {
				return nil
			}
			return fmt.Errorf("server error: %v", err)
		}
	default:
		return fmt.Errorf(
			"invalid transport type: %s. Must be 'stdio' or 'sse'",
			transport,
		)
	}
	return nil
}

func newMCPServer() *server.MCPServer {
	return server.NewMCPServer(
		"mcp_demo",
		Version,
		server.WithToolCapabilities(true),
		server.WithLogging(),
	)
}

func addTools(s *server.MCPServer) {
	s.AddTool(ip.GetCurrentIpInfoTool, ip.GetCurrentIpInfoHandler)
	s.AddTool(ip.GetIpInfoTool, ip.GetIpInfoHandler)
}
