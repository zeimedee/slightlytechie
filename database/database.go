package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDb() *sql.DB {
	str := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", os.Getenv("HOST"), os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))

	conn, err := sql.Open("postgres", str)
	if err != nil {
		panic(err)
	}

	// defer conn.Close()

	return conn
}
