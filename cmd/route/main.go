package main

import (
	"fmt"

	"github.com/isikhi/go-rate-limiter/internal/server"
)

func main() {
	s := server.New()
	s.InitDomains()
	fmt.Print("Registered Routes:\n\n")
	s.PrintAllRegisteredRoutes()
}
