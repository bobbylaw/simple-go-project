package config

import (
	"database/sql"
	"fmt"
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

var onCreate = func(database *sql.DB) {
	statement1, _ := database.Prepare("CREATE TABLE IF NOT EXISTS team (id INTEGER PRIMARY KEY, team_name TEXT, registration_date TEXT)")
	statement1.Exec()
	statement2, _ := database.Prepare("CREATE TABLE IF NOT EXISTS group_record (id INTEGER, team_id INTEGER, number_of_win INTEGER, number_of_lose INTEGER, number_of_draw INTEGER, total_goal INTEGER, total_score INTEGER, PRIMARY KEY (id, team_id), FOREIGN KEY (team_id) REFERENCES team(id))")
	statement2.Exec()
}

func TruncateDatabase() {
	statement1, _ := db.Prepare("DELETE FROM team")
	statement1.Exec()

	fmt.Println("Deleting Teams")

	statement2, _ := db.Prepare("DELETE FROM group_record")
	statement2.Exec()

	fmt.Println("Deleting Group Records")
}

func GetDB() *sql.DB {
	return db
}
