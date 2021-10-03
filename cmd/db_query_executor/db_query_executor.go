package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		return
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL_MODE"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		`
			INSERT INTO users (name, age)
			VALUES ($1, $2)
		`,
		"Masahiko Kobayashi",
		27,
	)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`
		SELECT id, name, age 
		FROM users 
		ORDER BY name
	`)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var id int64
		var name string
		var age int64
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name, age)
	}

	row := db.QueryRow(
		`
			SELECT name, age
			FROM users
			WHERE id = $1
		`,
		5,
	)
	var name string
	var age int64
	err = row.Scan(&name, &age)
}
