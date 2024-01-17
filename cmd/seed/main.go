package main

import (
	"fmt"
	"github.com/isikhi/go-rate-limiter/config"
	"github.com/isikhi/go-rate-limiter/database"
	db "github.com/isikhi/go-rate-limiter/third_party/database"
)

func main() {
	cfg := config.New()
	store := db.NewSqlx(cfg.Database)

	seeder := database.Seeder(store.DB)
	seeder.SeedUsers()
	fmt.Println("seeding completed.")
}
