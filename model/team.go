package model

import (
	"database/sql"
	"fmt"
)

type Team struct {
	Name             string
	RegistrationDate string
	GroupID          int
}

func AddTeam(database *sql.DB, team Team) {
	statement, _ := database.Prepare("INSERT INTO team (id, team_name, registration_date) VALUES (?, ?, ?)")
	statement.Exec(nil, team.Name, team.RegistrationDate)

	fmt.Printf("Added %v %v \n", team.Name, team.RegistrationDate)
}
