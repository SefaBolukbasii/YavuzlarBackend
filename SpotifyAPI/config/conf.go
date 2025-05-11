package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func Conf() {
	DbConnect()

}
func DbConnect() {
	var err error
	Db, err = sql.Open("postgres", "host=spotify-db user=postgres password=admin1234 dbname=spotify sslmode=disable")
	if err != nil {
		log.Fatal("Database connection error")
	}
}
