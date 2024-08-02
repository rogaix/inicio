package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() {
	var err error

	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	dbname := os.Getenv("MYSQL_DATABASE")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, host, dbname)

	for {
		db, err = sql.Open("mysql", dataSourceName)
		if err != nil {
			log.Printf("Failed to connect to database: %v, retrying in 3 seconds...\n", err)
			time.Sleep(3 * time.Second)
			continue
		}

		if err = db.Ping(); err != nil {
			log.Println("Failed to ping database, retrying in 3 seconds...")
			time.Sleep(3 * time.Second)
			continue
		}

		break
	}
}

func SetUpDatabase() {
	err := os.Setenv("MYSQL_HOST", "mysql")
	if err != nil {
		log.Fatal("Error setting environment variable:", err)
	}

	initDB()

	err = db.Ping()
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}

	log.Println("Database setup complete")
}

func GetDB() *sql.DB {
	if db == nil {
		SetUpDatabase()
	}
	return db
}
