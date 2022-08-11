package model

import (
	"database/sql"
	"log"
)

type Team struct {
	ID               int
	Name             string
	RegistrationDate string
}

func AddTeam(database *sql.DB, team Team) {
	statement, _ := database.Prepare("INSERT INTO team (id, team_name, registration_date) VALUES (?, ?, ?)")
	_, err := statement.Exec(team.ID, team.Name, team.RegistrationDate)
	if err != nil {
		log.Fatal(err)
	}
}

func GetAllTeam(database *sql.DB) []Team {
	rows, err := database.Query("SELECT id, team_name, registration_date FROM team")

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	teams := make([]Team, 0)

	for rows.Next() {
		currTeam := Team{}

		err = rows.Scan(&currTeam.ID, &currTeam.Name, &currTeam.RegistrationDate)
		if err != nil {
			log.Fatal(err)
		}

		teams = append(teams, currTeam)
	}

	return teams
}
