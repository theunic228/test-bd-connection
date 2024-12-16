package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // импортируем pq для работы с PostgreSQL
)

var DB *sql.DB

// Инициализация соединения с базой данных
func init() {
	var err error
	// Строка подключения к PostgreSQL
	connStr := "user=*** dbname=*** password=*** host=*** port=5442 sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Проверка соединения
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
