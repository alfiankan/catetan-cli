package controller

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func Migrate() {
	log.Println("Creating ctt_db.db...")
	file, err := os.Create("ctt_db.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("ctt_db.db created")
	sqliteDatabase, _ := sql.Open("sqlite3", "./ctt_db.db")

	createTable(sqliteDatabase)
}

func createTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE ctt (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"title" TEXT,
		"createdAt" TEXT,
		"modifiedAt" TEXT,
		"content" TEXT		
	  );` // SQL Statement for Create Table

	log.Println("Create table...")
	statement, err := db.Prepare(createTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("table created")
}
