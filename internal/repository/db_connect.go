package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPool(dbUrl string) *sql.DB {
	pool, err := sql.Open("pgx", dbUrl)
	if err != nil {
		log.Fatalf("Unable to connect to the datbase 😭:%v\n", err)
	}

	if err := pool.Ping(); err != nil {
		log.Fatalf("Unable to ping database 🧐: %v\n", err)
	}
	fmt.Println("Database connected successfully 😎")
	return pool
}
