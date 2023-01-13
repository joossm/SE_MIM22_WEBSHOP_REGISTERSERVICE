package model

import "database/sql"

func CloseDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		print(err)
	}
}

func OpenDB() *sql.DB {
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/books")
	if err != nil {
		print(err)
	}
	return db
}
