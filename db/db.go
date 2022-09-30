package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBname   string
}

func Init() (*sql.DB, error) {
	cfg := getDBConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Printf("DB Connection Error: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("DB Connection Error: %v", err)
		return nil, err
	}
	log.Printf("Successfully connected to DB")
	return db, nil
}

func getDBConfig() DBConfig {
	cfg := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBname:   os.Getenv("DB_NAME"),
	}
	return cfg
}
