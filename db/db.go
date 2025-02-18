package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Global db variable to hold the connection
var DB *sql.DB

func InitDB() *sql.DB {
	//check .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the environment variables
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbn := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	// Build the connection string correctly
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, dbn, port))

	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to psql database.")
	DB = db
	return db
}

// Exec executes a query on the existing database connection
func Exec(query string, args ...interface{}) (sql.Result, error) {
	return DB.Exec(query, args...)
}
