package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB
var err error

//CreateDBConnection Connect App to POSTGRES Relational Database
func CreateDBConnection() {
	// Load .env file
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error while loading env file")
	}
	// Open the connection
	DB, err = sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}
	// defer DB.Close()

	// check the connection
	err = DB.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("📪 Connected to Database!")
}

func GetDB() *sql.DB {
	return DB
}
