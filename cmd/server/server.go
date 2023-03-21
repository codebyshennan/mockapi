package main

import (
	"embed"

	server "bitbucket.org/libertywireless/circles-sandbox/internal/server"
)

//go:embed static
var static embed.FS

func main() {
	server.RunServer(static)
}
