package database

import (
	"context"
	"crypto/rand"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alexedwards/argon2id"
)

type Seed struct {
	DB *sql.DB
}

func Seeder(db *sql.DB) *Seed {
	return &Seed{
		DB: db,
	}
}

type user struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func userTableIsExistsCheck(m *Seed) bool {
	for i := 0; i < 10; i++ {
		fmt.Printf("User table is exists control #%d: ", i+1)
		var result string
		err := m.DB.QueryRow("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'users')").Scan(&result)
		if err != nil {
			panic(err)
		}

		if result == "true" {
			fmt.Println("User table is OK")
			return true
		} else {
			fmt.Println("User table not exists. Retry:", time.Second)
		}

		// 1 saniye bekletme
		time.Sleep(time.Second)
	}
	panic("User table not exists.")
}

func (m *Seed) SeedUsers() {
	users := []user{
		{
			FirstName: "First Name",
			LastName:  "Last Name",
			Email:     "admin@test.com",
			Password:  randomAndWrite(16),
		},
	}
	userTableIsExistsCheck(m)
	for _, u := range users {
		password, err := argon2id.CreateHash(u.Password, argon2id.DefaultParams)
		if err != nil {
			log.Fatalln(err)
		}
		_, err = m.DB.ExecContext(
			context.Background(),
			`INSERT INTO users (first_name, last_name, email, password, verified_at) 
				VALUES ($1, $2, $3, $4, $5);`,
			u.FirstName,
			u.LastName,
			u.Email,
			password,
			time.Now(),
		)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func writeToEnv(password string) {
	f, err := os.OpenFile(".env",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString("\nADMIN_PASSWORD=" + password + "\n"); err != nil {
		log.Println(err)
	}
}

func randomAndWrite(n int) string {
	adminPassword, exists := os.LookupEnv("ADMIN_PASSWORD")
	if exists {
		writeToEnv(adminPassword)
		return adminPassword
	}
	var chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()_+"

	ll := len(chars)
	b := make([]byte, n)
	_, _ = rand.Read(b)
	for i := 0; i < n; i++ {
		b[i] = chars[int(b[i])%ll]
	}

	str := string(b)
	fmt.Printf("Password is: %s\n", str)

	writeToEnv(str)

	return str
}
