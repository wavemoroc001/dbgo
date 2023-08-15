package db

import (
	"database/sql"
	"log"
)

func ProvideDBCon() (*sql.DB, func()) {
	conn, err := sql.Open("postgres", "postgresql://admin:admin123@localhost:5432/testdb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")
	return conn, func() {
		conn.Close()
	}
}
