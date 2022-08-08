package model

type Team struct {
	Name             string
	RegistrationDate string
	GroupID          int
}

type GroupRecord struct {
	GroupID      int
	Team         Team
	NumberOfWin  int
	NumberOfLose int
	NumberOfDraw int
	TotalGoal    int
	TotalScore   int
}
