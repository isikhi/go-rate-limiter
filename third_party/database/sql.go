package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/isikhi/go-rate-limiter/config"
	_ "github.com/isikhi/go-rate-limiter/ent/gen/runtime"
	"github.com/isikhi/go-rate-limiter/internal/utility/database"
)

func New(cfg *config.Config) *sql.DB {
	var dsn string
	switch cfg.Database.Driver {
	case "postgres":
		dsn = fmt.Sprintf("%s://%s:%d/%s?sslmode=%s&user=%s&password=%s",
			cfg.Database.Driver,
			cfg.Database.Host,
			cfg.Database.Port,
			cfg.Database.Name,
			cfg.Database.SslMode,
			cfg.Database.User,
			cfg.Database.Pass)
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true",
			cfg.Database.User,
			cfg.Database.Pass,
			cfg.Database.Host,
			cfg.Database.Port,
			cfg.Database.Name,
		)
	default:
		log.Fatal("Must choose a database driver")
	}

	db, err := sql.Open(cfg.Database.Driver, dsn)
	if err != nil {
		log.Fatal(err)
	}

	database.Alive(db)

	db.SetMaxOpenConns(cfg.Database.MaxConnectionPool)

	return db
}
