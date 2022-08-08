package utils

import (
	"fmt"
	"simple-go-project/model"
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
