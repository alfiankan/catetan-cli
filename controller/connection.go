package controller

import "database/sql"

func Connect() *sql.DB {
	sqliteDatabase, _ := sql.Open("sqlite3", "./ctt_db.db")

	return sqliteDatabase
}
