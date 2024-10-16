package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func Connect(cfg *Config) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	DB, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
		return nil, err
	}

	// Set connection pool settings
	DB.SetMaxOpenConns(cfg.DBMaxOpenConns)
	DB.SetMaxIdleConns(cfg.DBMaxIdleConns)
	DB.SetConnMaxLifetime(time.Duration(cfg.DBConnMaxLifetime) * time.Second)

	// Verify connection
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Error connecting to Database: %v", err)
		return nil, err
	}

	log.Println("Successfully connected to the database")

	return DB, nil
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	return DB
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection is closed")
	}
}
