package model

import (
	"database/sql"
	"fmt"
)

type Team struct {
	ID               int
	Name             string
	RegistrationDate string
}

func AddTeam(database *sql.DB, team Team) {
	statement, _ := database.Prepare("INSERT INTO team (id, team_name, registration_date) VALUES (?, ?, ?)")
	statement.Exec(team.ID, team.Name, team.RegistrationDate)

	fmt.Printf("Added %v %v \n", team.Name, team.RegistrationDate)
}
