package model

import (
	"database/sql"
	"fmt"
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

	fmt.Printf("Added %v %v \n", groupRecord.GroupID, groupRecord.Team.Name)
}
