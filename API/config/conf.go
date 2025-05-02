package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func Conf() {
	DbConnect()
	DbCreation()
}
func DbConnect() {
	var err error
	Db, err = sql.Open("postgres", "host=quest-db user=postgres password=admin1234 dbname=quest sslmode=disable")
	if err != nil {
		log.Fatal("Database connection error")
	}
}
func DbCreation() {
	_, err := Db.Exec("CREATE TABLE IF NOT EXISTS challange (id SERIAL PRIMARY KEY, name VARCHAR(50) Not NULL)")
	if err != nil {
		log.Fatal("Tablo Oluşturma Hatası")
	}
	_, err = Db.Exec("CREATE TABLE IF NOT EXISTS quest (id SERIAL PRIMARY KEY, question VARCHAR(150) ,answer VARCHAR(150),challange_id INT REFERENCES challange(id) ON DELETE CASCADE)")
	if err != nil {
		log.Fatal("Soru tablosu oluşturlamadı")
	}
	_, err = Db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, username VARCHAR(50) Not NULL,sifre VARCHAR(50),role VARCHAR(50))")
	if err != nil {
		log.Fatal("Kullanıcı Tablosu oluşturulamadı")
	}
	_, err = Db.Exec("CREATE TABLE IF NOT EXISTS USER_CHALLANGE (id SERIAL PRIMARY KEY, user_id INT REFERENCES users(id), challange_id INT REFERENCES challange(id) ON DELETE NO ACTION, score INT DEFAULT 0, UNIQUE(user_id, challange_id))")
	if err != nil {
		log.Fatal("Kullanıcı_Challange tablosu oluşturulamadı")
	}
}
