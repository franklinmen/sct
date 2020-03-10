package config

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	// sql server driver
	_ "github.com/denisenkom/go-mssqldb"
)

func GetConnection() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	}
	server := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	driver := os.Getenv("DB_DRIVER")
	database := os.Getenv("DB_NAME")

	//dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;", server, user, password, port)
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", user, password, server, port, database)
	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}