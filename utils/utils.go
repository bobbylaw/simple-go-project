package utils

import (
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

func CreateTeams(rawData string, output *map[string]model.Team) {
	teams := strings.Split(rawData, "\n")
	*output = make(map[string]model.Team)
	for _, str := range teams {
		team := strings.Split(str, " ")
		(*output)[team[0]] = model.Team{
			Name:             team[0],
			RegistrationDate: team[1],
			GroupID:          ConvertToInt(strings.TrimSpace(team[2])),
		}
	}
}

func CreateGroupRecord(rawData string, output *map[int]map[string]model.GroupRecord) {
	teams := strings.Split(rawData, "\n")
	*output = make(map[int]map[string]model.GroupRecord)

	for _, str := range teams {
		team := strings.Split(str, " ")
		groupID := ConvertToInt(strings.TrimSpace(team[2]))
		if (*output)[groupID] == nil {
			(*output)[groupID] = make(map[string]model.GroupRecord)
		}

		(*output)[groupID][team[0]] = model.GroupRecord{
			GroupID: groupID,
			Team: model.Team{
				Name:             team[0],
				RegistrationDate: team[1],
			},
			NumberOfWin:  0,
			NumberOfLose: 0,
			TotalGoal:    0,
			TotalScore:   0,
		}
	}
}

func UpdateMatchResult(rawData string, teams *map[string]model.Team, records *map[int]map[string]model.GroupRecord) {
	matches := strings.Split(rawData, "\n")
	for _, str := range matches {
		match := strings.Split(str, " ")

		firstTeamName := match[0]
		secondTeamName := match[1]
		firstTeamGoals := ConvertToInt(strings.TrimSpace(match[2]))
		secondTeamGoals := ConvertToInt(strings.TrimSpace(match[3]))
		firstTeamGroupID := (*teams)[firstTeamName].GroupID
		secondTeamGroupID := (*teams)[secondTeamName].GroupID

		firstTeamRecord := (*records)[firstTeamGroupID][firstTeamName]
		secondTeamRecord := (*records)[secondTeamGroupID][secondTeamName]
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

		(*records)[firstTeamGroupID][firstTeamName] = firstTeamRecord
		(*records)[secondTeamGroupID][secondTeamName] = secondTeamRecord
	}
}

func SortResult(records *map[int]map[string]model.GroupRecord) [][]model.GroupRecord {
	totalGroup := make([][]model.GroupRecord, 0, len(*records))

	for groupID := range *records {
		currGroup := make([]model.GroupRecord, 0, len((*records)[groupID]))

		for team := range (*records)[groupID] {
			currGroup = append(currGroup, (*records)[groupID][team])
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
