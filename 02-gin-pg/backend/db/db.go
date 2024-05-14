package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func ConnectDatabase() {
	port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), port, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PASSWORD"))
	db, errSql := sql.Open("postgres", psqlSetup)
	if errSql != nil {
		fmt.Println("There is an error while connecting to the database ", errSql)
		panic(errSql)
	} else {
		Db = db
		fmt.Println("Successfully connected to database!")
	}
}

func CreateTableUser() {
	_, err := Db.Exec("CREATE TABLE IF NOT EXISTS users (username text, password text);")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Create/Init table! users")
	}
}
