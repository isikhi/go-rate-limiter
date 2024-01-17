package main

import (
	//	"fmt"
	"github.com/isikhi/go-rate-limiter/internal/server"
	//	"os"
)

// Version is injected using ldflags during build time
var Version = "v0.1.0"

func main() {
	s := server.New(server.WithVersion(Version))
	s.Init()
	s.Run()
}
