package main

import (
	"log"

	"github.com/isikhi/go-rate-limiter/config"
	"github.com/isikhi/go-rate-limiter/database"
	db "github.com/isikhi/go-rate-limiter/third_party/database"
)

// Version is injected using ldflags during build time
var Version string

func main() {
	log.Printf("Version: %s\n", Version)

	cfg := config.New()
	store := db.NewSqlx(cfg.Database)
	migrator := database.Migrator(store.DB)

	migrator.Up()
}
