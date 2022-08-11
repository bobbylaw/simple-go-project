package utils

import (
	"database/sql"
	"fmt"
	"simple-go-project/model"
	"sort"
	"strconv"
	"strings"
)

func ConvertToInt(num string) int {
	res, err := strconv.Atoi(num)
	if err != nil {
		fmt.Println(err)
	}
	return res
}

func Register(rawData string, database *sql.DB) {
	teams := strings.Split(rawData, "\n")
	for index, str := range teams {
		team := strings.Split(str, " ")
		groupID := ConvertToInt(strings.TrimSpace(team[2]))
		newTeam := model.Team{
			ID:               index + 1,
			Name:             team[0],
			RegistrationDate: team[1],
		}

		newGroupRecord := model.GroupRecord{
			GroupID:      groupID,
			Team:         newTeam,
			NumberOfWin:  0,
			NumberOfLose: 0,
			TotalGoal:    0,
			TotalScore:   0,
		}

		model.AddTeam(database, newTeam)
		model.AddGroupRecord(database, newGroupRecord)
	}
}

func ConvertGroups(groups []model.GroupRecord) map[int]map[int]model.GroupRecord {
	output := make(map[int]map[int]model.GroupRecord)
	for _, val := range groups {
		if output[val.GroupID] == nil {
			output[val.GroupID] = make(map[int]model.GroupRecord)
		}

		output[val.GroupID][val.Team.ID] = val
	}

	return output
}

func UpdateMatchResult(rawData string, database *sql.DB) {
	matches := strings.Split(rawData, "\n")
	for _, str := range matches {
		match := strings.Split(str, " ")

		firstTeamName := match[0]
		secondTeamName := match[1]
		firstTeamGoals := ConvertToInt(strings.TrimSpace(match[2]))
		secondTeamGoals := ConvertToInt(strings.TrimSpace(match[3]))

		firstTeamRecord := model.GetGroupRecord(database, firstTeamName)
		secondTeamRecord := model.GetGroupRecord(database, secondTeamName)

		firstTeamRecord.TotalGoal += firstTeamGoals
		secondTeamRecord.TotalGoal += secondTeamGoals

		if secondTeamGoals > firstTeamGoals {
			secondTeamRecord.NumberOfWin++
			secondTeamRecord.TotalScore += 3
			firstTeamRecord.NumberOfLose++
		} else if firstTeamGoals > secondTeamGoals {
			firstTeamRecord.NumberOfWin++
			firstTeamRecord.TotalScore += 3
			secondTeamRecord.NumberOfLose++
		} else {
			firstTeamRecord.NumberOfDraw++
			secondTeamRecord.NumberOfDraw++
			firstTeamRecord.TotalScore += 1
			secondTeamRecord.TotalScore += 1
		}

		model.UpdateGroupRecord(database, firstTeamRecord)
		model.UpdateGroupRecord(database, secondTeamRecord)
	}
}

func SortResult(records map[int]map[int]model.GroupRecord) [][]model.GroupRecord {
	totalGroup := make([][]model.GroupRecord, 0, len(records))

	for groupID := range records {
		currGroup := make([]model.GroupRecord, 0, len((records)[groupID]))

		for team := range (records)[groupID] {
			currGroup = append(currGroup, (records)[groupID][team])
		}

		sort.SliceStable(currGroup, func(i, j int) bool {
			if currGroup[i].TotalScore == currGroup[j].TotalScore {
				if currGroup[i].TotalGoal == currGroup[j].TotalGoal {
					firstTeamNewScore := 0
					secondTeamNewScore := 0

					firstTeamNewScore += currGroup[i].NumberOfWin * 5
					firstTeamNewScore += currGroup[i].NumberOfDraw * 3
					firstTeamNewScore += currGroup[i].NumberOfLose * 1

					secondTeamNewScore += currGroup[j].NumberOfWin * 5
					secondTeamNewScore += currGroup[j].NumberOfDraw * 3
					secondTeamNewScore += currGroup[j].NumberOfLose * 1

					if firstTeamNewScore == secondTeamNewScore {
						firstTeamRegDate := strings.Split(currGroup[i].Team.RegistrationDate, "/")
						secondTeamRegDate := strings.Split(currGroup[j].Team.RegistrationDate, "/")
						if firstTeamRegDate[1] == secondTeamRegDate[1] {
							return firstTeamRegDate[0] < secondTeamRegDate[0]
						}
						return firstTeamRegDate[1] < secondTeamRegDate[1]
					}

					return firstTeamNewScore > secondTeamNewScore
				}
				return currGroup[i].TotalGoal > currGroup[j].TotalGoal
			}
			return currGroup[i].TotalScore > currGroup[j].TotalScore
		})
		totalGroup = append(totalGroup, currGroup)
	}

	return totalGroup

}

func CheckIfTeamRegistered(database *sql.DB) bool {
	teams := model.GetAllTeam(database)
	return len(teams) != 0
}

func CheckIfMatchInputted(database *sql.DB) bool {
	group := model.GetAllGroupRecord(database)

	if len(group) == 0 {
		return false
	}

	for _, val := range group {
		if val.TotalScore != 0 {
			return true
		}
	}

	return false
}
