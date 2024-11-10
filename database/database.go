package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Kennedy-lsd/ExpenseTracker/config"
	_ "github.com/lib/pq"
)

func Init() (*sql.DB, error) {
	conf := config.InitEnv()
	consStr := fmt.Sprintf("user=%s dbname=%s sslmode=%s", conf.User, conf.DbName, conf.SSL)
	db, err := sql.Open("postgres", consStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Successfully connected to the database!")

	return db, nil
}
