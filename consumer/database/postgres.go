package database

import (
	"database/sql"
	"log"
	"os"
	"fmt"
)

func Connect() *sql.DB{
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}

	fmt.Println("Database connected")

	createTableCommand := `
		CREATE TABLE IF NOT EXISTS skills (
		key TEXT PRIMARY KEY,
		name TEXT NOT NULL DEFAULT '',
		description TEXT NOT NULL DEFAULT '',
		logo TEXT NOT NULL DEFAULT '',
		tags TEXT [] NOT NULL DEFAULT '{}'
	);

	`
	_, err = db.Exec(createTableCommand)

	if err != nil {
		log.Fatal("Can't create table", err)
	}

	fmt.Println("Create table success")

	return db

}