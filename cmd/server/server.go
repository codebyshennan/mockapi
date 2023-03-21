package main

import (
	"embed"

	server "github.com/codebyshennan/mockapi/internal/server"
)

//go:embed static
var static embed.FS

func main() {
	server.RunServer(static)
}
