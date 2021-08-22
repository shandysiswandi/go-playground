package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(db)
}

func connectDB() (*pgx.Conn, error) {
	// "postgres://username:password@localhost:5432/postgres"
	urlExample := ""
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		return nil, err
	}

	defer conn.Close(context.Background())
	return conn, nil
}
