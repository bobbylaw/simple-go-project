package model

import (
	"database/sql"
	"log"
)

type GroupRecord struct {
	GroupID      int
	Team         Team
	NumberOfWin  int
	NumberOfLose int
	NumberOfDraw int
	TotalGoal    int
	TotalScore   int
}

func AddGroupRecord(database *sql.DB, groupRecord GroupRecord) {
	statement, _ := database.Prepare("INSERT INTO group_record (id, team_id, number_of_win, number_of_lose, number_of_draw, total_goal, total_score) VALUES (?, ?, ?, ?, ?, ?, ?)")
	statement.Exec(groupRecord.GroupID, groupRecord.Team.ID, groupRecord.NumberOfWin, groupRecord.NumberOfLose, groupRecord.NumberOfDraw, groupRecord.TotalGoal, groupRecord.TotalScore)
}

func GetGroupRecord(database *sql.DB) []GroupRecord {
	rows, err := database.Query("SELECT g.id, t.id, t.team_name, t.registration_date, g.number_of_win, g.number_of_lose, g.number_of_draw, g.total_goal, g.total_score FROM group_record as g join team as t on g.team_id = t.id")

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	groups := make([]GroupRecord, 0)

	for rows.Next() {
		currGroup := GroupRecord{
			Team: Team{},
		}

		err = rows.Scan(&currGroup.GroupID, &currGroup.Team.ID, &currGroup.Team.Name, &currGroup.Team.RegistrationDate, &currGroup.NumberOfWin, &currGroup.NumberOfLose, &currGroup.NumberOfDraw, &currGroup.TotalGoal, &currGroup.TotalScore)
		if err != nil {
			log.Fatal(err)
		}

		groups = append(groups, currGroup)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return groups
}
