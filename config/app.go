package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

var RegisterDatabase = func() {
	database, err := sql.Open("sqlite3", "./database/database.db")
	if err != nil {
		log.Fatal(err)
	}

	onCreate(database)
	db = database
}

func onCreate(database *sql.DB) {
	statement1, _ := database.Prepare("CREATE TABLE IF NOT EXISTS team (id INTEGER PRIMARY KEY, team_name TEXT, registration_date TEXT)")
	_, err1 := statement1.Exec()

	if err1 != nil {
		log.Fatal(err1)
	}

	statement2, _ := database.Prepare("CREATE TABLE IF NOT EXISTS group_record (id INTEGER, team_id INTEGER, number_of_win INTEGER, number_of_lose INTEGER, number_of_draw INTEGER, total_goal INTEGER, total_score INTEGER, PRIMARY KEY (id, team_id), FOREIGN KEY (team_id) REFERENCES team(id))")
	_, err2 := statement2.Exec()

	if err2 != nil {
		log.Fatal(err2)
	}
}

func TruncateDatabase() {
	statement1, _ := db.Prepare("DELETE FROM team")
	_, err1 := statement1.Exec()

	if err1 != nil {
		log.Fatal(err1)
	}

	statement2, _ := db.Prepare("DELETE FROM group_record")
	_, err2 := statement2.Exec()

	if err2 != nil {
		log.Fatal(err2)
	}
}

func GetDB() *sql.DB {
	return db
}
